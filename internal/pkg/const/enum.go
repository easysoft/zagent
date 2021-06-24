package _const

type ResultCode int

const (
	ResultSuccess ResultCode = 1
	ResultFail    ResultCode = 0
)

func (c ResultCode) Int() int {
	return int(c)
}

type HttpMethod string

const (
	Get    HttpMethod = "GET"
	Post   HttpMethod = "POST"
	Put    HttpMethod = "PUT"
	Delete HttpMethod = "DELETE"
)

type ValidMethod string

const (
	ValidProjectPath ValidMethod = "validProjectPath"
	ValidDictName    ValidMethod = "validDictName"
)
