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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles": {
            "get": {
                "description": "Get Article List based on query params",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Get Article List",
                "operationId": "get-article-list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "input search text",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success Response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.DefaultResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Article"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error Response",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an article based on given body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Create an article",
                "operationId": "create-article",
                "parameters": [
                    {
                        "description": "article body",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateArticleModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success Response",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request Response",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error Response",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Article": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/models.Person"
                },
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "default": "Lorem"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateArticleModel": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/models.Person"
                },
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "default": "Lorem"
                }
            }
        },
        "models.DefaultResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Person": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string",
                    "example": "John"
                },
                "lastname": {
                    "type": "string",
                    "example": "Doe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.1",
	Host:             "localhost:7070",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}