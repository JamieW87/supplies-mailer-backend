resource "aws_ecs_cluster" "services_cluster" {
  name = format("onestop-%s-server-cluster", var.environment)
}

resource "aws_ecs_task_definition" "api" {
  count              = var.deploy_api ? 1 : 0
  family             = "api-services"
  execution_role_arn = var.ecs_role_arn
  container_definitions = jsonencode([
    {
      name      = "api"
      image     = format("%s:%s", var.api_ecr_repo, var.api_service_version)
      essential = true
      memory    = var.api_memory
      secrets = [
        {
          valueFrom = format("arn:aws:secretsmanager:%s:294786226104:secret:%s/environment", var.aws_region, var.environment)
          name      = "SECRETS_MANAGER_ENVIRONMENT_VARIABLES"
        }
      ]
      portMappings = [
        {
          containerPort = 8091
          hostPort      = 0
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region        = var.aws_region
          awslogs-stream-prefix = format("api-%s-logs", var.environment)
          awslogs-group         = format("api-%s-logs", var.environment)
        }
      }
    }
  ])
}

resource "aws_ecs_service" "api_service" {
  count                             = var.deploy_api ? 1 : 0
  cluster                           = aws_ecs_cluster.services_cluster.id
  desired_count                     = var.api_service_count
  launch_type                       = "EC2"
  name                              = "api"
  task_definition                   = aws_ecs_task_definition.api[0].arn
  health_check_grace_period_seconds = "300"
  load_balancer {
    container_name   = "api"
    container_port   = "8091"
    target_group_arn = var.api_target_group_arn
  }
  force_new_deployment = true
}