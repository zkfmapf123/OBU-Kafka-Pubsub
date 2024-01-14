provider "aws" {
  profile = "root"
  region  = "ap-northeast-2"
}

resource "aws_security_group" "ins-sg" {
  vpc_id      = "vpc-07b29bf03817840e5"
  name        = "inst-sg"
  description = "description"

  ingress {
    from_port   = 0
    to_port     = 0
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
    Name = "default-ins-sg"
  }
}

## Freetier는 cpu 옵션을 사용할수 없습니다.
module "default-public-ins" {
  source = "zkfmapf123/simpleEC2/lee"

  instance_name      = "pubsub-test"
  instance_region    = "ap-northeast-2a"
  instance_subnet_id = "subnet-06240758a52eb3483"
  instance_sg_ids    = [aws_security_group.ins-sg.id]

  instance_ami  = "ami-0a3c0384147c78fac"
  instance_type = "t4g.medium"

  instance_ip_attr = {
    is_public_ip  = true
    is_eip        = true
    is_private_ip = false
    private_ip    = ""
  }

  instance_root_device = {
    size = 50
    type = "gp3"
  }

  instance_key_attr = {
    is_alloc_key_pair = false
    is_use_key_path   = true
    key_name          = ""
    key_path          = "~/.ssh/id_rsa.pub"
  }

  instance_tags = {
    "Monitoring" : true,
    "MadeBy" : "terraform",
    "Name" : "pubsub-test"
  }
}
