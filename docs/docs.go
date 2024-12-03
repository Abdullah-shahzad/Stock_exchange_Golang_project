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
        "/stocks": {
            "get": {
                "description": "Retrieves all stock entries from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stock"
                ],
                "summary": "Retrieve all stocks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.Stock"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Saves new stock data into the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stock"
                ],
                "summary": "Create a new stock entry",
                "parameters": [
                    {
                        "description": "Stock data",
                        "name": "stock",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateStockRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/stocks/{ticker}": {
            "get": {
                "description": "Retrieves stock details based on the provided ticker symbol.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stock"
                ],
                "summary": "Retrieve stock by ticker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stock Ticker",
                        "name": "ticker",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Stock"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "Creates a new transaction record in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "description": "Transaction data",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{username}": {
            "get": {
                "description": "Retrieves a list of all transactions for a given user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get all Transactions for a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.Transaction"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{username}/{start_time}/{end_time}": {
            "get": {
                "description": "Retrieves transactions for a user within a specified timestamp.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get Transactions for a user by timestamp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "Start Date",
                        "name": "start_time",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "End Date",
                        "name": "end_time",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.Transaction"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Saves new user data into the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "description": "Retrieves user details based on the provided username.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateStockRequest": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number",
                    "example": 150.25
                },
                "ticker": {
                    "type": "string",
                    "example": "AAPL"
                }
            }
        },
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controllers.Stock": {
            "type": "object",
            "required": [
                "price",
                "ticker"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "ticker": {
                    "type": "string"
                }
            }
        },
        "controllers.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.Transaction": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "ticker": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "transaction_price": {
                    "type": "number"
                },
                "transaction_type": {
                    "type": "string"
                },
                "transaction_volume": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.TransactionRequest": {
            "type": "object",
            "properties": {
                "ticker": {
                    "type": "string",
                    "example": "Tia"
                },
                "transaction_type": {
                    "type": "string",
                    "example": "BUY"
                },
                "transaction_volume": {
                    "type": "integer",
                    "example": 10
                },
                "username": {
                    "type": "string",
                    "example": "abdullah"
                }
            }
        },
        "controllers.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.UserRequest": {
            "type": "object",
            "properties": {
                "initial_balance": {
                    "type": "number",
                    "example": 1000
                },
                "username": {
                    "type": "string",
                    "example": "abdullah"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Stock exchange API",
	Description:      "A stock exchange API project using the Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
