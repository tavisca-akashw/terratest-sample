resource "aws_security_group" "allow_ssh" {
    ingress {
        from_port = "22"
        to_port   = "22"
        protocol  = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
}


resource "aws_instance" "terratest_ec2" {
    ami           = "ami-5026902d"
    instance_type = "t2.micro"
    key_name = "terra_test_key"
    vpc_security_group_ids = ["${aws_security_group.allow_ssh.id}]"
    user_data = <<-EOF
    #!/bin/bash
    echo 'Hello, World!' > /tmp/salut
    EOF
}


