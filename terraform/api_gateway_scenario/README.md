# API Gateway integrated with Load balancer and Lambda as the target (Deploy with Terraform)


In this repository, using Terraform, you can deploy a sample application with ALB and Lambda as the target.

### Components to be created with Terraform

1. New VPC
2. Two Subnets
3. Internet Gateway
4. Route Table with IGW attached (Here default route table is used)
5. Application Load Balancer
6. Security Group for ALB
7. target Group
8. Lambda function
9. Add Lambda as a target in ALB target group.
10. API Gateway with ALB as http_proxy for get requests and lambda as proxy for any http method other than get

### How to Use

1. Clone the repository.

2. run terrform initizalize and apply commands

terrform init
terraform apply

5. You can access the API using the output value `deployment_invoke_url`

### To delete the stack

By default ALB's termination policy is set to `true` which will prevent delete it. So, first needs to set it as false.

In the terraform.tfvars file set the `ALB_DELETION_PROTECTION` value to `false`

To apply the change, run
```
terraform plan -out tfout.plan
terraform apply tfout.plan
```

And to destroy the stack, run
```
terraform destroy
```


