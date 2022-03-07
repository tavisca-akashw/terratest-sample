# terratest-sample


This Respository Contains terraform manifests and terratest files for EC2 instance launch and api Gateway.
Please install below binaries required for the execution :
* Terraform version >= v1.0.0
* Golang version 1.17.6
* gcc (requires for go tests)

After installing this binaries you can execute test cases by following below steps :

Change you directory to the required test case for e.g for ec2 test case go to the folder **ec2_instance**
Then go to the location where we have test file located that is **terratest**
once you are in the terratest folder run below commands 

First install modules you imported in the test using 


**go mod init instance_ssh_test.go**


**go mod tidy**


 And finally execute test 
**go test -v ec2_instance_test.go**

