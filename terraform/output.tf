output "instance_id" {
    value = "${aws_instance.terratest_ec2.id}"
}

output "instance_ssh_key" {
    value = "${aws_instance.terratest_ec2.key_name}"
}
