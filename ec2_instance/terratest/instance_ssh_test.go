package terratest
import (
   "testing"
   "github.com/stretchr/testify/assert"
   "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestEc2SshKey(t *testing.T) {
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                TerraformDir: "../terraform",
        })
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    ec2SshKey  := terraform.Output(t, terraformOptions, "instance_ssh_key")
    assert.Equal(t, "terratest", ec2SshKey)
}
