variable "aws_region" {}

variable "state_bucket" {
  default = "onestop-terraform-state"
}

variable "environment" {}

variable "pem_key_name" {}

variable "ssh_security_group_cidr_blocks" {}

variable "postgres_security_group_cidr_blocks" {}

variable "vpc_id" {}

variable "count_of_services_servers" {}

variable "services_server_instance_type" {}

variable "services_server_ami" {}

variable "services_server_subnet_ids" {}

variable "services_server_volume_size" {}

variable "deploy_api" {}

variable "deploy_api_lb" {}

variable "api_service_count" {}

variable "api_service_version" {}

variable "api_memory" {}

variable "log_retention_in_days" {}

variable "provision_db" {}

variable "db_instance_is_provisioned" {}

variable "db_user" {}

variable "db_password" {}

variable "db_engine_version" {}

variable "db_instance_class" {}

variable "db_publicly_accessible" {}

variable "db_skip_final_snapshot" {}

variable "count_of_db_instances" {}