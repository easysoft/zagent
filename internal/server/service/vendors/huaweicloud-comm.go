package vendors

import (
	"encoding/json"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	testconst "github.com/easysoft/zagent/cmd/test/const"
	"github.com/easysoft/zagent/internal/comm/domain"
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

func (s HuaweiCloudCommService) Create(image string, jobName string, cmd []string,
	token string, region string, nameapace string) (ret domain.CciRepsCreate, success bool) {

	reqCreate := domain.CciReqCreate{
		APIVersion: "batch/v1",
		Kind:       "Job",
		Metadata: domain.CciMetadata{
			Name: jobName,
		},
		SpecTempl: domain.CciSpecTempl{
			Template: domain.CciTemplate{
				Metadata: domain.CciMetadata{
					Name: jobName,
				},
				Spec: domain.CciSpec{
					Containers: []domain.CciContainers{
						{
							Name:  jobName,
							Image: image,
							Resources: domain.CciResources{
								Limits: domain.CciLimits{
									CPU:    "2000m",
									Memory: "4096Mi",
								},
								Requests: domain.CciRequests{
									CPU:    "2000m",
									Memory: "4096Mi",
								},
							},
							Command: cmd,
						}},
					ImagePullSecrets: []domain.ImagePullSecrets{
						{
							Name: "imagepull-secret",
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	createUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobCreate, region, nameapace)
	var resp []byte
	resp, success = s.post(createUrl, reqCreate, nil, map[string]string{"X-Auth-Token": token})

	json.Unmarshal(resp, &ret)
	name := ret.Metadata.Name
	_logUtils.Infof("%s#v", name)

	return
}

func (s HuaweiCloudCommService) Destroy(jobName string, token string, region string, nameapace string) (
	ret domain.CciRepsDestroy, success bool) {

	reqDestroy := domain.CciReqDestroy{
		Kind:              "DeleteOptions",
		APIVersion:        "v1",
		PropagationPolicy: "Orphan",
	}

	destroyUrl := fmt.Sprintf(testconst.HuaweiCloudUrlJobDestroy,
		testconst.HUAWEI_CLOUD_REGION, testconst.HUAWEI_CLOUD_NAMEAPACE, jobName)

	var resp []byte
	resp, success = s.delete(destroyUrl, reqDestroy, nil, map[string]string{"X-Auth-Token": token})

	json.Unmarshal(resp, ret)
	_logUtils.Infof("%#v", ret)

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
