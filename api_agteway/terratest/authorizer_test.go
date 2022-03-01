package terratest
import (
   "strings"
   "fmt"
   "strconv"
   "github.com/stretchr/testify/assert"
   terra_test "github.com/gruntwork-io/terratest/modules/testing"
   testing "testing"
   //"github.com/aws/aws-sdk-go/aws"
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
    apiId :=  terraform.Output(t, terraformOptions, "api_id")
    apiAuthorizer := GetAPIGwAuthorizers(t, awsRegion, apiId)
    apiAuthorizert := fmt.Sprint(apiAuthorizer)
    apiAuthorizerUri := fmt.Sprintf(strconv.FormatBool(strings.Contains(apiAuthorizert, YOUR_AUTH_LAMBDA_FUNCTION)))
    assert.Equal(t, "true" , apiAuthorizerUri)
}




func GetAPIGwAuthorizers(t terra_test.TestingT, awsRegion string, apiId string) []string {
        ApiAuthorizersUri, err := GetAPIGwAuthorizersE(t, awsRegion, apiId)
        require.NoError(t, err)

        return ApiAuthorizersUri
}



func GetAPIGwAuthorizersE(t terra_test.TestingT, awsRegion string, apiId string) ([]string, error) {
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

        authorizersUri := []string{}
        for _, authorizer := range out.Items {

              authorizersUri = append(authorizersUri, *authorizer.AuthorizerUri)
        }
        return authorizersUri, err
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
