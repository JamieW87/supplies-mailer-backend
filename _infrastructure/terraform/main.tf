provider "aws" {
  region = var.aws_region
}

terraform {
  backend "s3" {}
}

# Configure the remote state data source
data "terraform_remote_state" "networks" {
  backend = "s3"

  config = {
    bucket = var.state_bucket
    key    = format("%s.tfstate", var.environment)
    region = var.aws_region
  }
}

module "cloudwatch" {
  source = "./module-cloudwatch"

  environment = var.environment

  log_retention_in_days = var.log_retention_in_days

  deploy_api = var.deploy_api
}

module "ec2" {
  source = "./module-ec2"

  environment = var.environment

  pem_key_name = var.pem_key_name

  ecs_iam_instance_profile_id = module.iam.ecs_iam_instance_profile_id

  vpc_id = var.vpc_id

  ssh_security_group_cidr_blocks = var.ssh_security_group_cidr_blocks

  postgres_security_group_cidr_blocks = var.postgres_security_group_cidr_blocks

  count_of_services_servers = var.count_of_services_servers
  services_server_ami = var.services_server_ami
  services_server_instance_type = var.services_server_instance_type
  services_server_subnet_ids = var.services_server_subnet_ids
  services_server_volume_size = var.services_server_volume_size

  deploy_api_lb = var.deploy_api_lb
}

module "ecr" {
  source = "./module-ecr"
}

module "ecs" {
  source = "./module-ecs"

  environment = var.environment

  aws_region = var.aws_region

  ecs_role_arn = module.iam.ecs_role_arn

  deploy_api = var.deploy_api
  api_memory = var.api_memory
  api_service_count = var.api_service_count
  api_service_version = var.api_service_version
  api_ecr_repo = module.ecr.repo
  api_target_group_arn = module.ec2.api_target_group_arn
}

module "iam" {
  source = "./module-iam"

  environment = var.environment
}

module "rds" {
  source = "./module-rds"

  environment = var.environment

  vpc_id = var.vpc_id

  postgres_security_group = module.ec2.postgres_security_group

  provision_db = var.provision_db
  count_of_db_instances = var.count_of_db_instances
  db_engine_version = var.db_engine_version
  db_instance_class = var.db_instance_class
  db_instance_is_provisioned = var.db_instance_is_provisioned
  db_password = var.db_password
  db_publicly_accessible = var.db_publicly_accessible
  db_skip_final_snapshot = var.db_skip_final_snapshot
  db_user = var.db_user
}
