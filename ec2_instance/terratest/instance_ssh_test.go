package terratest

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEc2SshKey(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	ec2SshKey := terraform.Output(t, terraformOptions, "instance_ssh_key")
	assert.Equal(t, "terratest", ec2SshKey)
}
