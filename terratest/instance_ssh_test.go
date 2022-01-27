package terratest
import (
   "testing"
   "github.com/stretchr/testify/assert"
   "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestInstanceSshKey(t *testing.T) {
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                TerraformDir: "../terraform",
        })
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    instanceSshKey := terraform.Output(t, terraformOptions, "instance_key")
    assert.Equal(t, "terratest", instanceSshKey)
}
