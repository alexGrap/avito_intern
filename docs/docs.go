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
        "/createSegment": {
            "post": {
                "description": "Create a segment. You need to post name of segment and if this need percent of user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segment"
                ],
                "summary": "Create a new segment",
                "parameters": [
                    {
                        "description": "Body of creation segment",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SegmentBody"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "OK"
                    },
                    "503": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/deleteSegment": {
            "delete": {
                "description": "Delete a segment. You need to post name of segment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segment"
                ],
                "summary": "Delete a segment",
                "parameters": [
                    {
                        "description": "Name of segment",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SegmentBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "503": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/getById": {
            "get": {
                "description": "Get list of user` + "`" + `s subscriptions.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get subscriptions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Bad request"
                    },
                    "503": {
                        "description": "Not found"
                    }
                }
            }
        },
        "/getSegment": {
            "get": {
                "description": "Get a segment.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segment"
                ],
                "summary": "Get all existed segments",
                "responses": {
                    "204": {
                        "description": "OK"
                    },
                    "503": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/subscription": {
            "put": {
                "description": "Create new and delete existed subscription for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Make a subscription",
                "parameters": [
                    {
                        "description": "Body of subscription",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Subscriber"
                        }
                    }
                ],
                "responses": {
                    "203": {
                        "description": "OK"
                    },
                    "503": {
                        "description": "Bad request"
                    }
                }
            }
        },
        "/timeoutSubscribe": {
            "put": {
                "description": "Get user subscriptions history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "History"
                ],
                "summary": "Get history",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Time start. ex: 2020-03-20",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Time end. ex: 2024-03-20",
                        "name": "to",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "203": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad request"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SegmentBody": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "percents": {
                    "type": "integer"
                },
                "segmentName": {
                    "type": "string"
                }
            }
        },
        "models.SubscribeWithTimeout": {
            "type": "object",
            "properties": {
                "segmentName": {
                    "type": "string"
                },
                "timeToDie": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Subscriber": {
            "type": "object",
            "properties": {
                "add": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "delete": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "userId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "avito_intern",
	Description:      "Intern REST-service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
