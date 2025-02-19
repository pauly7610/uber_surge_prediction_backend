provider "aws" {
  region = "us-west-2"
}

resource "aws_ecs_cluster" "default" {
  name = "surge-pricing-cluster"
}

resource "aws_ecs_service" "service" {
  name            = "surge-pricing-service"
  cluster         = aws_ecs_cluster.default.id
  desired_count   = 3
  launch_type     = "FARGATE"
  task_definition = aws_ecs_task_definition.surge_pricing.arn
}

resource "aws_ecs_task_definition" "surge_pricing" {
  family                   = "surge-pricing"
  container_definitions    = jsonencode([{
    name  = "surge-pricing-container"
    image = "your-docker-image"
    cpu   = 256
    memory = 512
    essential = true
  }])
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = "arn:aws:iam::your-account-id:role/ecsTaskExecutionRole"
} 