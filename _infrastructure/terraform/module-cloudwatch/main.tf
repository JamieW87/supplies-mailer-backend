resource "aws_cloudwatch_log_group" "api" {
  count             = var.deploy_api ? 1 : 0
  name              = format("api-%s-logs", var.environment)
  retention_in_days = var.log_retention_in_days
}
