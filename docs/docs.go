// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/opening": {
            "put": {
                "description": "Update Opening job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Update Opening",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Opening indetification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Opening data to update",
                        "name": "opening",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateOpeningRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateOpeningReponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create  a new Opening job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Create Opening",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateOpeningRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateOpeningReponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Opening job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Openings"
                ],
                "summary": "Delete Opening",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Opening indetification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteOpeningReponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorReponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateOpeningReponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.CreateOpeningRequest": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "handler.DeleteOpeningReponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ErrorReponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateOpeningReponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.OpeningResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateOpeningRequest": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "schemas.OpeningResponse": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
