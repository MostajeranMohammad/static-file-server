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
        "/file-meta-data/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "no comment",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "no comment",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "no comment",
                        "name": "uploader_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "no comment",
                        "name": "bucket_name",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/file-meta-data/delete-file/{file_name}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "no comment",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/file-meta-data/update-file-access/{file_name}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "no comment",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "no comment",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StaticFileMetaDataUpdateAccessDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/file-meta-data/{file_name}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "no comment",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/static-file/download/{file_name}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "for download private objects put you jwt here.",
                        "name": "authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "no comment",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/static-file/upload/{bucket_name}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "no comment",
                        "name": "bucket_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "no comment",
                        "name": "user_ids_who_access_this_file",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "no comment",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.StaticFileMetaDataUpdateAccessDto": {
            "type": "object",
            "required": [
                "userIdsWhoAccessThisFile"
            ],
            "properties": {
                "userIdsWhoAccessThisFile": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
