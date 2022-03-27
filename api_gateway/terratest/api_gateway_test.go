package terratest

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestApiGateway(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	stageUrl := terraform.Output(t, terraformOptions, "deployment_invoke_url")
	time.Sleep(30 * time.Second)
	statusCode := DoGetRequest(t, stageUrl)
	assert.Equal(t, 200, statusCode)
}

func DoGetRequest(t *testing.T, api string) int {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response status on the line below.
	return resp.StatusCode
}
