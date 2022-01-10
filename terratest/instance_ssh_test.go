package test
import (
   "testing"
   "fmt"
   "time"
   "github.com/stretchr/testify/assert"
   "github.com/gruntwork-io/terratest/modules/terraform"
   "github.com/gruntwork-io/terratest/modules/aws"
   "github.com/gruntwork-io/terratest/modules/ssh"
   "github.com/gruntwork-io/terratest/modules/retry"
)

func TestInstanceSshKey(t *testing.T) {
    terraformOptions := configureTerraformOptions(t)
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    instanceSshKey := terraform.Output(t, terraformOptions, "instance_key")
    assert.Equal(t, "terra_test_key", instanceSshKey)
}
