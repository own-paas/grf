// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/user/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页页面",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页大小",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "搜索",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"data\": {\"count\": 0, \"results\": []}}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "传入参数是struct",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"message\": \"操作成功！\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "查看用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"data\":{}}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "传入参数是struct",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"message\": \"操作成功！\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"message\": \"操作成功！\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "电话号码",
                    "type": "string"
                },
                "updated_at": {
                    "description": "创建时间",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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