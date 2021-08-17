package vendors

import (
	"github.com/alibabacloud-go/tea/tea"
	testconst "github.com/easysoft/zagent/cmd/test/const"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iamModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
)

type HuaweiCloudCommService struct {
}

func NewHuaweiCloudCommService() *HuaweiCloudCommService {
	return &HuaweiCloudCommService{}
}

func (s HuaweiCloudCommService) GetIamToken(client *iam.IamClient) (
	token string, err error) {

	request := &iamModel.KeystoneCreateUserTokenByPasswordRequest{
		Body: &iamModel.KeystoneCreateUserTokenByPasswordRequestBody{
			Auth: &iamModel.PwdAuth{
				Identity: &iamModel.PwdIdentity{
					Methods: []iamModel.PwdIdentityMethods{
						iamModel.GetPwdIdentityMethodsEnum().PASSWORD},
					Password: &iamModel.PwdPassword{
						User: &iamModel.PwdPasswordUser{
							Domain: &iamModel.PwdPasswordUserDomain{
								Name: testconst.HUAWEI_CLOUD_USER,
							},
							Name:     testconst.HUAWEI_CLOUD_IAM_USER,
							Password: testconst.HUAWEI_CLOUD_IAM_PASSWORD,
						},
					},
				},
				Scope: &iamModel.AuthScope{
					Domain: &iamModel.AuthScopeDomain{
						Name: tea.String(testconst.HUAWEI_CLOUD_USER),
					},
				},
			},
		},
	}
	response, err := client.KeystoneCreateUserTokenByPassword(request)

	token = *response.XSubjectToken

	return
}

func (s HuaweiCloudCommService) CreateIamClient(ak, sk, regionId string) (
	client *iam.IamClient, err error) {

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client = iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(region.ValueOf(regionId)).
			WithCredential(auth).
			Build())

	return
}
