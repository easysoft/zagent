// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package vm

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/easysoft/zv/issues",
            "email": "462626@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/service/check": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "检测虚拟机服务状态",
                "parameters": [
                    {
                        "description": "Service Check Request Object",
                        "name": "VmServiceCheckReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.VmServiceCheckReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code = success | fail",
                        "schema": {
                            "$ref": "#/definitions/v1.VmServiceCheckResp"
                        }
                    }
                }
            }
        },
        "/api/v1/service/setup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "安装虚拟机服务",
                "parameters": [
                    {
                        "description": "Service Install Request Object",
                        "name": "VmServiceCheckReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.VmServiceInstallReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code = success | fail",
                        "schema": {
                            "$ref": "#/definitions/v1.VmServiceInstallResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.VmServiceCheckReq": {
            "type": "object",
            "properties": {
                "services": {
                    "type": "string"
                }
            }
        },
        "v1.VmServiceCheckResp": {
            "type": "object",
            "properties": {
                "zdStatus": {
                    "description": "Enums consts.HostServiceStatus",
                    "type": "string"
                },
                "zdVersion": {
                    "type": "string"
                },
                "ztfStatus": {
                    "description": "Enums consts.HostServiceStatus",
                    "type": "string"
                },
                "ztfVersion": {
                    "type": "string"
                }
            }
        },
        "v1.VmServiceInstallReq": {
            "type": "object",
            "properties": {
                "ip": {
                    "description": "testing node ip, port ztf:56202, zd:56203",
                    "type": "string"
                },
                "name": {
                    "description": "tool name, ztf or zd",
                    "type": "string"
                },
                "secret": {
                    "description": "secret to access zentao",
                    "type": "string"
                },
                "server": {
                    "description": "zentao server url",
                    "type": "string"
                },
                "version": {
                    "description": "tool version",
                    "type": "string"
                }
            }
        },
        "v1.VmServiceInstallResp": {
            "type": "object",
            "properties": {
                "version": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "ZAgent虚拟机API文档",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
