// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/index": {
            "get": {
                "tags": [
                    "Index"
                ],
                "responses": {
                    "200": {
                        "description": "Welcome!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/createUser": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "username, password, repassword",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.createUserInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Create user success!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Passwords not same",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "delete": {
                "tags": [
                    "User"
                ],
                "summary": "DeleteUser",
                "parameters": [
                    {
                        "description": "username, password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.loginUserInfo"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Delete user success!"
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid username or password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUser": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "GetUser by username and password",
                "parameters": [
                    {
                        "description": "username, password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.loginUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid username or password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "GetUserList",
                "responses": {
                    "200": {
                        "description": "Get user list",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "put": {
                "tags": [
                    "User"
                ],
                "summary": "UpdateUser",
                "parameters": [
                    {
                        "description": "username, password, password/telephone/email, data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.updateUserInfo"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Update user success!"
                    },
                    "400": {
                        "description": "Invalid parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid username or password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.createUserInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "rePassword": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.loginUserInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "service.updateUserInfo": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "parameter": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
