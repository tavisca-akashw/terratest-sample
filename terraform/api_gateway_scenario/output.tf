output "lb_address" {
  value = aws_lb.load-balancer.dns_name
  description = "DNS of load balancer"
}


output "deployment_invoke_url" {
  description = "Deployment invoke url"
  value       = "${aws_api_gateway_stage.test.invoke_url}/resource"
}
