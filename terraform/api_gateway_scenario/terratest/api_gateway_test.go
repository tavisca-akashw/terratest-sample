package terratest
import (
   "log"
   "net/http"
   "time"
   terra_test "github.com/gruntwork-io/terratest/modules/testing"
   testing "testing"
   "github.com/stretchr/testify/assert"
   "github.com/aws/aws-sdk-go/aws"
   terra_aws "github.com/gruntwork-io/terratest/modules/aws"
   "github.com/gruntwork-io/terratest/modules/terraform"
   "github.com/aws/aws-sdk-go/service/apigateway"
   "github.com/stretchr/testify/require"
)

func TestApiGateway(t *testing.T) {
    //awsRegion := "eu-west-2"
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    //stageName := terraform.Output(t, terraformOptions, "stage_name")
    //apiId :=  terraform.Output(t, terraformOptions, "api_id")
    //actualStage := GetAPIGwStage(t, awsRegion, apiId)
    //assert.Equal(t, stageName , actualStage)
    stageUrl := terraform.Output(t, terraformOptions,"deployment_invoke_url")
    time.Sleep(30 * time.Second)
    statusCode := CurlRequest(t, stageUrl)
    assert.Equal(t, 200 , statusCode)
}


func CurlRequest(t terra_test.TestingT, api string) int{
   resp, err := http.Get(api)
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response status on the line below.
   return resp.StatusCode
}


func GetAPIGwStage(t terra_test.TestingT, awsRegion string, apiId string) string {
        StageName, err := GetAPIGwStageE(t, awsRegion, apiId)
	require.NoError(t, err)

	return StageName
}



func GetAPIGwStageE(t terra_test.TestingT, awsRegion string, apiId string) (string, error) {
	apiGWClient, err := NewAPIgwClientE(t, awsRegion)
	if err != nil {
		return "", err
	}

	res, err := apiGWClient.GetStage(&apigateway.GetStageInput{
		RestApiId: &apiId,
	})
	if err != nil {
		return "", err
	}

	return aws.StringValue(res.StageName), nil
}


func NewAPIgwClient(t terra_test.TestingT, region string) *apigateway.APIGateway {
	client, err := NewAPIgwClientE(t, region)
	require.NoError(t, err)

	return client
}

// NewAPIgwClientE creates an APIGw client.
func NewAPIgwClientE(t terra_test.TestingT, region string) (*apigateway.APIGateway, error) {
	sess, err := terra_aws.NewAuthenticatedSession(region)
	if err != nil {
		return nil, err
	}

	return apigateway.New(sess), nil
}
