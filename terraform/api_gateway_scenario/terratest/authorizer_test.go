package terratest
import (
   "log"
   "fmt"
   //"bytes"
   terra_test "github.com/gruntwork-io/terratest/modules/testing"
   testing "testing"
   "github.com/aws/aws-sdk-go/aws"
   terra_aws "github.com/gruntwork-io/terratest/modules/aws"
   "github.com/gruntwork-io/terratest/modules/terraform"
   "github.com/aws/aws-sdk-go/service/apigateway"
   "github.com/stretchr/testify/require"
)

func TestApiGateway(t *testing.T) {
    awsRegion := "eu-west-1"
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                TerraformDir: "../",
        })
        //defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    //stageName := terraform.Output(t, terraformOptions, "stage_name")
    apiId :=  terraform.Output(t, terraformOptions, "api_id")
    apiAuthorizer := GetAPIGwAuthorizers(t, awsRegion, apiId)
    log.Println("Map: " + fmt.Sprint(apiAuthorizer))
    //assert.Equal(t, stageName , actualStage)
}




func GetAPIGwAuthorizers(t terra_test.TestingT, awsRegion string, apiId string) map[string]string {
        ApiAuthorizers, err := GetAPIGwAuthorizersE(t, awsRegion, apiId)
        require.NoError(t, err)

        return ApiAuthorizers
}



func GetAPIGwAuthorizersE(t terra_test.TestingT, awsRegion string, apiId string) (map[string]string, error) {
        apiGWClient, err := NewAPIgwClientE(t, awsRegion)
        if err != nil {
                return nil, err
        }

        out, err := apiGWClient.GetAuthorizers(&apigateway.GetAuthorizersInput{
                RestApiId: &apiId,
        })
        if err != nil {
                return nil, err
        }

        authorizers := map[string]string{}
        for _, authorizer := range out.Items {

                authorizers[aws.StringValue(authorizer.Name)] = aws.StringValue(authorizer.Type)
                //authorizers[aws.StringValue(authorizer.Name)] = aws.StringValue(authorizer.Name)
        }
        return authorizers, nil
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
