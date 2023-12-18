resource "aws_rds_cluster" "postgres_db_cluster" {
  count                  = var.provision_db ? 1 : 0
  cluster_identifier     = format("%s-db-cluster", var.environment)
  engine                 = "aurora-postgresql"
  engine_mode            = "provisioned"
  engine_version         = var.db_engine_version
  database_name          = format("onestop_%s", var.environment)
  master_username        = var.db_user
  master_password        = var.db_password
  vpc_security_group_ids = [var.postgres_security_group]
  skip_final_snapshot    = var.db_skip_final_snapshot
  db_subnet_group_name   = format("default-%s", var.vpc_id)
}

resource "aws_rds_cluster_instance" "postgres_db_cluster_instances" {
  count               = (var.provision_db && var.db_instance_is_provisioned) ? var.count_of_db_instances : 0
  apply_immediately   = true
  identifier          = format("%s-db-%d", var.environment, count.index + 1)
  cluster_identifier  = aws_rds_cluster.postgres_db_cluster[0].id
  instance_class      = var.db_instance_class
  engine              = aws_rds_cluster.postgres_db_cluster[0].engine
  engine_version      = aws_rds_cluster.postgres_db_cluster[0].engine_version
  publicly_accessible = var.db_publicly_accessible
}