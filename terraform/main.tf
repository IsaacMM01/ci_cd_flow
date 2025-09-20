provider "aws" {
  region = var.aws_region
}

resource "aws_security_group" "docker_app_sg" {
  name        = "docker_app_sg"
  description = "Permite SSH y HTTP"
  vpc_id      = var.vpc_id

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "docker_app_sg"
  }
}

resource "aws_instance" "docker_app_ec2" {
  ami                         = var.ami_id
  instance_type               = "t2.micro"
  subnet_id                   = var.subnet_id
  vpc_security_group_ids      = [aws_security_group.docker_app_sg.id]
  key_name                    = var.key_pair_name

  tags = {
    Name = "docker_app_ami"
  }
}