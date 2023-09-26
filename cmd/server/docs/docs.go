// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/odontologos": {
            "post": {
                "description": "Create a new Odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Odontologo example",
                "parameters": [
                    {
                        "name": "Nombre",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Apellido",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Matricula",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/odontologos/:id": {
            "get": {
                "description": "Get odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Odontologo example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id del odontologo",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Odontologo example",
                "parameters": [
                    {
                        "name": "Nombre",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Apellido",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Matricula",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id del odontologo",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "odontologo example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id del odontologo",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "UpdateName odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "odontologo example",
                "parameters": [
                    {
                        "name": "key",
                        "in": "query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "value",
                        "in": "query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id del odontologo",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/pacientes": {
            "post": {
                "description": "Create a new paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "paciente example",
                "parameters": [
                    {
                        "name": "Nombre",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Apellido",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Domicilio",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "DNI",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "name": "FechaAlta",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/pacientes/:id": {
            "get": {
                "description": "Get paciente by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "paciente example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update paciente by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "paciente example",
                "parameters": [
                    {
                        "name": "Nombre",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Apellido",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Domicilio",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "DNI",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "name": "FechaAlta",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete paciente by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "paciente example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "UpdateSubjet paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "paciente"
                ],
                "summary": "paciente example",
                "parameters": [
                    {
                        "name": "key",
                        "in": "query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "value",
                        "in": "query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/turnoByPaciente": {
            "post": {
                "description": "Create a turno with DNI of Paciente and Matricula of Odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "name": "Paciente",
                        "decription": "Paciente's DNI",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "name": "Odontologo",
                        "description": "Odontologo's matricula",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/turnos": {
            "post": {
                "description": "create a turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "name": "Paciente",
                        "decription": "Paciente's DNI",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "name": "Odontologo",
                        "description": "Odontologo's matricula",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "FechaHora",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Descripcion",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/turnos/:id": {
            "get": {
                "description": "Get a turno by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "name": "Paciente",
                        "decription": "Paciente's DNI",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "name": "Odontologo",
                        "description": "Odontologo's matricula",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "FechaHora",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "Descripcion",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a turno by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Any Subject on turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "name": "key",
                        "in": "query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "name": "value",
                        "in": "query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/turnos/dni/:dni": {
            "get": {
                "description": "Get a turno by DNI of paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turno"
                ],
                "summary": "turno example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paciente's DNI",
                        "name": "dni",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "web.errorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "web.response": {
            "type": "object",
            "properties": {
                "data": {}
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "in": "header",
            "name": "tokenPostman"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Odontologia grupo 4",
	Description:      "This is a sample od odontologia",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}