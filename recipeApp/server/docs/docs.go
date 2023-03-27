// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Wes Ahearn, Seth Paul"
        },
        "license": {
            "name": "unknown"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/server/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "allows a user to login",
                "parameters": [
                    {
                        "description": "user data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.userBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/server/recipes/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "post a recipe to the database",
                "parameters": [
                    {
                        "description": "recipe data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.recipePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/server/recipes/bypage": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get a json object (recipes) of default 10 recipes at a time",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "specify page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "results per page",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "specify one or more keywords",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "specify one or more ingredients",
                        "name": "ingredient",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "specify a user id",
                        "name": "uid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/server/recipes/delete/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete a recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of recipe to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/server/recipes/edit/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "edit a recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of recipe to edit",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "modified recipe data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.recipeEditRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/server/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "register a user to the site",
                "parameters": [
                    {
                        "description": "user data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.userBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.recipeEditRequest": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "image_name": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "string"
                },
                "instructions": {
                    "type": "string"
                },
                "rid": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "handler.recipePostRequest": {
            "type": "object",
            "properties": {
                "image_name": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "string"
                },
                "instructions": {
                    "type": "string"
                },
                "rid": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "handler.userBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Mallow Recipe Server API",
	Description:      "This is the server for the Mallow Recipe Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
