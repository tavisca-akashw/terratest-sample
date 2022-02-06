resource "aws_api_gateway_rest_api" "api" {
  name = "terratest_api"
}

resource "aws_api_gateway_resource" "resource" {
  path_part   = "resource"
  parent_id   = aws_api_gateway_rest_api.api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_method" "get_method" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_resource.resource.id
  http_method   = "GET"
  authorization = "NONE"
}


resource "aws_api_gateway_method" "any_method" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_resource.resource.id
  http_method   = "ANY"
  authorization = "NONE"
}


resource "aws_api_gateway_integration" "lambda_integration" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_resource.resource.id
  http_method             = aws_api_gateway_method.any_method.http_method
  integration_http_method = "ANY"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda.invoke_arn
}


resource "aws_api_gateway_integration" "alb_integration" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_resource.resource.id
  http_method             = aws_api_gateway_method.get_method.http_method
  integration_http_method = "GET"
  type                    = "HTTP_PROXY"
  uri                     = "http://${aws_lb.load-balancer.dns_name}"
}



resource "aws_security_group" "allow-http-traffic" {
  name        = "allow_http_traffic"
  description = "Allow http traffic"
  vpc_id      = aws_vpc.mydemovpc.id

  ingress {
    description = "Http traffic"
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
}



data "template_file" "api_gateway_assume_role_policy" {
  template = file("./policies/api_gateway_assume_role.json")
}

resource "aws_iam_role" "api-gateway-assume-role-policy" {
  assume_role_policy = data.template_file.api_gateway_assume_role_policy.rendered
}

resource "aws_iam_role_policy_attachment" "api-gateway-policy-role-attachment" {
  policy_arn =  "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role = aws_iam_role.api-gateway-assume-role-policy.name
}


# Provides lambda execution policy to api gateway

resource "aws_lambda_permission" "with_api_gateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = aws_lb_target_group.lambda-target-group.arn
}


resource "aws_api_gateway_deployment" "test_deployment" {
  depends_on = [
                 aws_api_gateway_rest_api.api,
                 aws_api_gateway_method.any_method,
                 aws_api_gateway_integration.lambda_integration,
                 aws_api_gateway_method.get_method,
                 aws_api_gateway_integration.alb_integration,
               ]
  rest_api_id = aws_api_gateway_rest_api.api.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.api.body))
  }

  lifecycle {
    create_before_destroy = true
  }

}

resource "aws_api_gateway_stage" "test" {
  deployment_id = aws_api_gateway_deployment.test_deployment.id
  rest_api_id   = aws_api_gateway_rest_api.api.id
  stage_name    = "test"
}
