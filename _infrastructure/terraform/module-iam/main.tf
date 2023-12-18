resource "aws_iam_role" "ecs_service_role" {
  name               = format("%s%s", var.environment, "-ecs-service-role")
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.ecs_service_policy.json
}

resource "aws_iam_role_policy_attachment" "ecs_service_role_attachment" {
  role       = aws_iam_role.ecs_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

data "aws_iam_policy_document" "ecs_service_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecs_instance_role" {
  name               = format("%s%s", var.environment, "-ecs-instance-role")
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.ecs_instance_policy.json
}

data "aws_iam_policy_document" "ecs_instance_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "ecs_instance_role_attachment" {
  role       = aws_iam_role.ecs_instance_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_role_policy_attachment" "secretsmanager_instance_policy_attachment" {
  policy_arn = aws_iam_policy.secretsmanager_policy.arn
  role       = aws_iam_role.ecs_instance_role.name
}

resource "aws_iam_instance_profile" "ecs_instance_profile" {
  name = format("%s%s", var.environment, "-ecs-instance-profile")
  path = "/"
  role = aws_iam_role.ecs_instance_role.id
  provisioner "local-exec" {
    command = "sleep 10"
  }
}

data "aws_iam_policy_document" "ecs_task_policy" {
  statement {
    sid = "EcsTaskPolicy"
    actions = [
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage"
    ]
    resources = [
      "*" # you could limit this to only the ECR repo you want
    ]
  }
  statement {
    actions = [
      "ecr:GetAuthorizationToken"
    ]
    resources = [
      "*"
    ]
  }
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "*"
    ]
  }
}

data "aws_iam_policy_document" "ecs_assume_role_policy" {
  statement {
    sid    = ""
    effect = "Allow"
    actions = [
      "sts:AssumeRole",
    ]
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecs_role" {
  name               = format("%s%s", var.environment, "-ecs-role")
  assume_role_policy = data.aws_iam_policy_document.ecs_assume_role_policy.json

  inline_policy {
    name   = "EcsTaskExecutionPolicy"
    policy = data.aws_iam_policy_document.ecs_task_policy.json
  }
}

resource "aws_iam_role_policy_attachment" "secretsmanager_policy_attachment" {
  policy_arn = aws_iam_policy.secretsmanager_policy.arn
  role       = aws_iam_role.ecs_role.name
}

resource "aws_iam_role_policy_attachment" "ses_policy_attachment" {
  policy_arn = aws_iam_policy.ses_policy.arn
  role       = aws_iam_role.ecs_role.name
}

resource "aws_iam_policy" "secretsmanager_policy" {
  name   = format("%s%s", var.environment, "-sm-policy")
  policy = data.aws_iam_policy_document.secretsmanager_policy.json
}

data "aws_iam_policy_document" "secretsmanager_policy" {
  statement {
    sid    = ""
    effect = "Allow"
    actions = [
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:ListSecretVersionIds",
      "secretsmanager:GetRandomPassword",
      "secretsmanager:ListSecrets"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "ses_policy" {
  name   = format("%s%s", var.environment, "-ses-policy")
  policy = data.aws_iam_policy_document.ses_policy.json
}

data "aws_iam_policy_document" "ses_policy" {
  statement {
    sid    = ""
    effect = "Allow"
    actions = [
      "ses:SendEmail"
    ]
    resources = ["*"]
  }
}

output "ecs_iam_instance_profile_id" {
  value = aws_iam_instance_profile.ecs_instance_profile.id
}

output "ecs_role_arn" {
  value = aws_iam_role.ecs_role.arn
}
