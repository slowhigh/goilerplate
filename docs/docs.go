// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/auth/token": {
            "post": {
                "description": "Authenticates a user and provides a JWT to Authorize API calls",
                "produces": [
                    "application/json"
                ],
                "summary": "Provides a JSON Web Token",
                "operationId": "Authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User credentials",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User credentials",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JWT"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/videos": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get all the existing videos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos",
                    "list"
                ],
                "summary": "List existing videos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Video"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Create a new video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos",
                    "create"
                ],
                "summary": "Create new videos",
                "parameters": [
                    {
                        "description": "Create video",
                        "name": "video",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Video"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/videos/{id}": {
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Update a single video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Update videos",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Video ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update video",
                        "name": "video",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Video"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Delete a single video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "videos"
                ],
                "summary": "Remove videos",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Video ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.JWT": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.Person": {
            "type": "object",
            "required": [
                "age",
                "email",
                "firstname",
                "lastname"
            ],
            "properties": {
                "age": {
                    "description": "\"gte\"는 최소값, \"lte\"는 최대값, \"gte=1,lte=130\"는 1이상 100이하의 값",
                    "type": "integer",
                    "maximum": 120,
                    "minimum": 1
                },
                "email": {
                    "description": "\"email\"은 이메일 형식 유효성",
                    "type": "string"
                },
                "firstname": {
                    "description": "\"required\" =\u003e 필수 항목을 의미함",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "entity.Video": {
            "type": "object",
            "required": [
                "author",
                "url"
            ],
            "properties": {
                "author": {
                    "$ref": "#/definitions/entity.Person"
                },
                "description": {
                    "description": "Title       string ` + "`" + `gorm:\"type:varchar(10)\" json:\"title\" binding:\"min=2,max=10\" validate:\"is-cool\"` + "`" + ` // \"is-cool\"이라는 이름의 사용자 유효성 검사 사용 video-controller.go 파일에 New메소드 참조",
                    "type": "string",
                    "maxLength": 20
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 2
                },
                "url": {
                    "description": "\"url\"은 url 형식 유효성",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}