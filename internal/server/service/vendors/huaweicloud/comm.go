package huaweicloud

import (
	"encoding/json"
	"github.com/alibabacloud-go/tea/tea"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iamModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	"io/ioutil"
	"net/http"
	"strings"
)

type HuaweiCloudCommService struct {
}

func NewHuaweiCloudCommService() *HuaweiCloudCommService {
	return &HuaweiCloudCommService{}
}

func (s HuaweiCloudCommService) GetIamToken(huaweiCloudUser, huaweiCloudIamUser, huaweiCloudIamPassword string,
	client *iam.IamClient) (
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
								Name: huaweiCloudUser,
							},
							Name:     huaweiCloudIamUser,
							Password: huaweiCloudIamPassword,
						},
					},
				},
				Scope: &iamModel.AuthScope{
					Domain: &iamModel.AuthScopeDomain{
						Name: tea.String(huaweiCloudUser),
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

func (s HuaweiCloudCommService) post(url string, reqBody interface{}, params, headers map[string]string) (
	ret []byte, success bool) {
	client := &http.Client{}

	reqBodyStr, err := json.Marshal(reqBody)
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, false
	}

	if params != nil {
		url += "?"
		for key, val := range params {
			url += key + "=" + val + "&"
		}
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(reqBodyStr)))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	ret, _ = ioutil.ReadAll(resp.Body)
	success = true
	return
}

func (s HuaweiCloudCommService) delete(url string, reqBody interface{}, params, headers map[string]string) (
	ret []byte, success bool) {

	client := &http.Client{}

	reqBodyStr, err := json.Marshal(reqBody)
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, false
	}

	if params != nil {
		url += "?"
		for key, val := range params {
			url += key + "=" + val + "&"
		}
	}

	req, reqErr := http.NewRequest("DELETE", url, strings.NewReader(string(reqBodyStr)))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	ret, _ = ioutil.ReadAll(resp.Body)
	success = true
	return

	return
}
