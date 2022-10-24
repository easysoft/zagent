package _const

type HttpMethod string

const (
	Get    HttpMethod = "GET"
	Post   HttpMethod = "POST"
	Put    HttpMethod = "PUT"
	Delete HttpMethod = "DELETE"
)

func (e HttpMethod) ToString() string {
	return string(e)
}

type ValidMethod string

const (
	ValidProjectPath ValidMethod = "validProjectPath"
	ValidDictName    ValidMethod = "validDictName"
)

func (e ValidMethod) ToString() string {
	return string(e)
}
