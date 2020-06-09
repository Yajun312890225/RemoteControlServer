// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-08 15:13:08.092107 +0800 CST m=+0.058801191

package docs

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/binddevice": {
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "绑定设备",
                "parameters": [
                    {
                        "description": "UserDevice",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDevice"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/control": {
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "远程控制",
                "parameters": [
                    {
                        "description": "message",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getdevice": {
            "get": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "获取局域网设备",
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getstatus": {
            "get": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "获取状态",
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "小程序登录",
                "parameters": [
                    {
                        "description": "res",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Res"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Message": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "during": {
                    "description": "持续时间",
                    "type": "integer"
                },
                "type": {
                    "description": "1打开 2 关闭",
                    "type": "integer"
                }
            }
        },
        "model.Res": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.UserDevice": {
            "type": "object",
            "properties": {
                "deviceId": {
                    "type": "string"
                },
                "userId": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
