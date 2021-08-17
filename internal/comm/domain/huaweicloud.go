package domain

type CciReqCreate struct {
	APIVersion string       `json:"apiVersion"`
	Kind       string       `json:"kind"`
	Metadata   CciMetadata  `json:"metadata"`
	SpecTempl  CciSpecTempl `json:"spec"`
}

type CciReqDestroy struct {
	Kind              string `json:"Kind"`
	APIVersion        string `json:"apiVersion"`
	PropagationPolicy string `json:"propagationPolicy"`
}

type CciRepsCreate struct {
	APIVersion string       `json:"apiVersion"`
	Kind       string       `json:"kind"`
	Metadata   CciMetadata  `json:"metadata"`
	Spec       CciSpecTempl `json:"spec"`
}

type CciRepsDestroy struct {
	APIVersion string      `json:"apiVersion"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Metadata   CciMetadata `json:"metadata"`
}

type CciMetadata struct {
	Name string `json:"name"`
}
type CciLimits struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type CciRequests struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type CciResources struct {
	Limits   CciLimits   `json:"limits"`
	Requests CciRequests `json:"requests"`
}
type CciContainers struct {
	Command   []string     `json:"command"`
	Image     string       `json:"image"`
	Name      string       `json:"name"`
	Resources CciResources `json:"resources"`
}
type ImagePullSecrets struct {
	Name string `json:"name"`
}
type CciSpec struct {
	Containers       []CciContainers    `json:"containers"`
	ImagePullSecrets []ImagePullSecrets `json:"imagePullSecrets"`
	RestartPolicy    string             `json:"restartPolicy"`
}
type CciTemplate struct {
	Metadata CciMetadata `json:"metadata"`
	Spec     CciSpec     `json:"spec"`
}
type CciSpecTempl struct {
	Template CciTemplate `json:"template"`
}
