provider "aws" {
  region = "us-east-1"
}


resource "aws_security_group" "allow_ssh" {
    name = "test-sg"
    ingress {
        from_port = "22"
        to_port   = "22"
        protocol  = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
}


resource "aws_instance" "terratest_ec2" {
    ami           = "ami-08e4e35cccc6189f4"
    instance_type = "t2.micro"
    key_name = "terratest"
    security_groups = ["${aws_security_group.allow_ssh.name}"]
    user_data = <<-EOF
    #!/bin/bash
    echo 'Hello, World!' > /tmp/salut
    EOF
}
