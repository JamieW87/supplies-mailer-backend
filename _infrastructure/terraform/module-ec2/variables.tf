variable "environment" {}

variable "pem_key_name" {}

variable "ecs_iam_instance_profile_id" {}

variable "vpc_id" {}

variable "ssh_security_group_cidr_blocks" {}

variable "postgres_security_group_cidr_blocks" {}

variable "count_of_services_servers" {}

variable "services_server_instance_type" {}

variable "services_server_ami" {}

variable "services_server_subnet_ids" {}

variable "services_server_volume_size" {}

variable "deploy_api_lb" {}