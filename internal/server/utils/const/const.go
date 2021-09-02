package serverConst

const (
	WsNamespace   = "default"
	WsEvent       = "OnChat"
	WsDefaultRoom = "square"

	PageSize = 15

	ALIYUN_ECS_URL     = "ecs-%s.aliyuncs.com"
	ALIYUN_ECS_URL_VNC = "https://g.alicdn.com/aliyun/ecs-console-vnc2/0.0.8/index.html" +
		"?vncUrl=%s&instanceId=%s&isWindows=%t&password=%s"
	ALIYUN_ECI_URL = "eci.aliyuncs.com"
	ALIYUN_VPC_URL = "vpc.aliyuncs.com"

	HuaweiCloudUrlJobCreate  = "https://cci.%s.myhuaweicloud.cn/apis/batch/v1/namespaces/%s/jobs"
	HuaweiCloudUrlJobDestroy = "https://cci.%s.myhuaweicloud.cn/apis/batch/v1/namespaces/%s/jobs/%s"
)

type WsEventAction string

const (
	TaskUpdate WsEventAction = "task_update"
)

func (e WsEventAction) ToString() string {
	return string(e)
}

var (
	SlotTypeAbbrMap = map[string]string{"synonym": "syn", "lookup": "lkp", "regex": "rgx"}
)
