resource "aws_instance" "services_server" {
  count                       = var.count_of_services_servers
  ami                         = var.services_server_ami # Amazon Linux AMI
  iam_instance_profile        = var.ecs_iam_instance_profile_id
  instance_type               = var.services_server_instance_type
  subnet_id                   = var.services_server_subnet_ids[count.index % length(var.services_server_subnet_ids)]
  vpc_security_group_ids      = [aws_security_group.ssh_security_group.id, aws_security_group.instance_services_security_group.id]
  key_name                    = var.pem_key_name
  user_data                   = templatefile("${path.module}/install-ecs.tpl", { environment = var.environment })
  associate_public_ip_address = true

  root_block_device {
    volume_size = var.services_server_volume_size
  }

  tags = {
    Name        = format("onestop-%s-server-%d", var.environment, count.index + 1)
    Server_Type = "services-server"
    Server_Env  = var.environment
  }

  lifecycle {
    ignore_changes = [user_data]
  }
}

resource "aws_security_group" "ssh_security_group" {
  name   = "ssh-security-group"
  vpc_id = var.vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = var.ssh_security_group_cidr_blocks
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "services_security_group" {
  name   = format("%s-services-sg", var.environment)
  vpc_id = var.vpc_id

  ingress {
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  ingress {
    from_port        = 80
    to_port          = 80
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "instance_services_security_group" {
  name   = format("%s-inst-services-sg", var.environment)
  vpc_id = var.vpc_id

  ingress {
    protocol    = "tcp"
    from_port   = 32768
    to_port     = 65535
    description = "Access from ALB"

    security_groups = [
      aws_security_group.services_security_group.id,
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "postgres_security_group" {
  name   = "postgres-security-group"
  vpc_id = var.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = concat(var.postgres_security_group_cidr_blocks, formatlist("%s/32", aws_instance.services_server.*.private_ip))
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_lb" "api_load_balancer" {
  count           = var.deploy_api_lb ? 1 : 0
  internal        = false
  name            = format("%s-api-alb", var.environment)
  subnets         = var.services_server_subnet_ids
  security_groups = [aws_security_group.services_security_group.id]
}

resource "aws_lb_listener" "api_listener" {
  count             = var.deploy_api_lb ? 1 : 0
  load_balancer_arn = aws_lb.api_load_balancer[0].arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.api_target_group[0].id
    type             = "forward"
  }
}

resource "aws_lb_target_group" "api_target_group" {
  count                = var.deploy_api_lb ? 1 : 0
  name                 = format("%s-api-tg", var.environment)
  port                 = "80"
  protocol             = "HTTP"
  vpc_id               = var.vpc_id
  deregistration_delay = 120

  health_check {
    path    = "/health-check"
    matcher = "204"
  }
}

output "api_target_group_arn" {
  value = var.deploy_api_lb ? aws_lb_target_group.api_target_group[0].arn : null
}

output "postgres_security_group" {
  value = aws_security_group.postgres_security_group.id
}