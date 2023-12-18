#! /bin/bash
mkdir -p /etc/ecs
echo ECS_CLUSTER=onestop-${environment}-server-cluster > /etc/ecs/ecs.config
echo ECS_ENABLE_CONTAINER_METADATA=true >> /etc/ecs/ecs.config

sudo amazon-linux-extras disable docker

sudo amazon-linux-extras install -y ecs

sudo systemctl enable --now --no-block ecs.service
