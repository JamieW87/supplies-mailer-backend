resource "aws_ecr_repository" "repository" {
  name = "onestop-backend"
}

output "repo" {
  value = aws_ecr_repository.repository.repository_url
}
