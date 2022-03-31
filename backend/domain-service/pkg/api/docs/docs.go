// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Source Code",
            "url": "https://github.com/hm-edu/portal-backend"
        },
        "license": {
            "name": "Apache License",
            "url": "https://github.com/hm-edu/portal-backend/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/domains/": {
            "get": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Lists all domains that are either owned or delegated, or a child of a owned or delegated domain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "List domains.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Domain"
                            }
                        }
                    },
                    "400": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Creates a new domain if the FQDN is not already taken. Approvement is automatically done, in case the user owns a upper zone or a upper zone was already delegated to him.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "Create a domain.",
                "parameters": [
                    {
                        "description": "The Domain to create",
                        "name": "domain",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DomainRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "400": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/domains/{id}": {
            "delete": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Deletes a domain. Existing certificates are not are not longer accessible.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "Delete a domain",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Unauthorized or Request Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/domains/{id}/approve": {
            "post": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Approves an outstanding domain request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "Approve domain request",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "400": {
                        "description": "Unauthorized or Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Access to domain denied",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Domain in zone does not exist",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/domains/{id}/delegation": {
            "post": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Adds a new delegation to an existing domain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "Add delegation.",
                "parameters": [
                    {
                        "description": "The Delegation to add",
                        "name": "delegation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelegationRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "400": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/domains/{id}/delegation/{delegation}": {
            "delete": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Deletes an existing delegation.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "Delete delegation.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Delegation ID",
                        "name": "delegation",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "400": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/domains/{id}/transfer": {
            "post": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Transfers a domain to a new owner. Transfering is only possible if you are either the owner of the domain itself or responsible for one of the parent zones.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Domains"
                ],
                "summary": "Transfer domain",
                "parameters": [
                    {
                        "description": "The Domain to transfer",
                        "name": "domain",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TransferRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Domain ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Domain"
                        }
                    },
                    "400": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "Used by Kubernetes Liveness Probe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Liveness check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.healthResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/api.healthResponse"
                        }
                    }
                }
            }
        },
        "/readyz": {
            "get": {
                "description": "Used by Kubernetes Readiness Probe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Readiness check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.healthResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/api.healthResponse"
                        }
                    }
                }
            }
        },
        "/whoami": {
            "get": {
                "security": [
                    {
                        "API": []
                    }
                ],
                "description": "Returns the username of the logged in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "whoami Endpoint",
                "responses": {
                    "200": {
                        "description": "Username",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.healthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "model.Delegation": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "model.DelegationRequest": {
            "type": "object",
            "required": [
                "user"
            ],
            "properties": {
                "user": {
                    "type": "string"
                }
            }
        },
        "model.Domain": {
            "type": "object",
            "properties": {
                "approved": {
                    "type": "boolean"
                },
                "delegations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Delegation"
                    }
                },
                "fqdn": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "model.DomainRequest": {
            "type": "object",
            "required": [
                "fqdn"
            ],
            "properties": {
                "fqdn": {
                    "type": "string"
                }
            }
        },
        "model.TransferRequest": {
            "type": "object",
            "required": [
                "owner"
            ],
            "properties": {
                "owner": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "API": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Domain Service",
	Description:      "Go microservice for Domain management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
