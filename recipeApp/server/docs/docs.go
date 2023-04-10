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
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/server/favorites/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add a favorite item to the database",
                "parameters": [
                    {
                        "description": "favorite data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.favPostRequest"
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
        "/server/favorites/delete/{uid}/{rid}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete a favorite item from the database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id of favorite to delete",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "recipe id of favorite to delete",
                        "name": "rid",
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
        "/server/meals/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "post a mealplan item to the database",
                "parameters": [
                    {
                        "description": "meal data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.mealPostRequest"
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
        "/server/meals/bydate": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of meals in predefined date range.",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2023-04-01",
                        "description": "specify start date",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "specify a user id",
                        "name": "uid",
                        "in": "query",
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
        "/server/meals/delete/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete a meal",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of meal to delete",
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
        "/server/meals/edit": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update a mealplan item in the database",
                "parameters": [
                    {
                        "description": "meal data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.mealEditRequest"
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
        "/server/recipecount": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get the correct count of recipes for any pagination scenario",
                "parameters": [
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
                "summary": "Get a list of recipes in predefined amounts, searchable by keyword and ingredient.",
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
        "handler.favPostRequest": {
            "type": "object",
            "properties": {
                "recipeid": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "handler.mealEditRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "mealtype": {
                    "type": "string"
                },
                "mid": {
                    "type": "integer"
                },
                "recipeid": {
                    "type": "integer"
                }
            }
        },
        "handler.mealPostRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "mealtype": {
                    "type": "string"
                },
                "recipeid": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
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
