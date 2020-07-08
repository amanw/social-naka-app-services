// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server in Go with Swagger endpoints definition",
    "title": "social-naka-app",
    "version": "0.0.1"
  },
  "paths": {
    "/registerUser": {
      "post": {
        "description": "It registers the User\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Registers the User",
        "operationId": "registerUser",
        "parameters": [
          {
            "name": "User Body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "409": {
            "$ref": "#/responses/UserExists"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "It gets all the Users\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Get all the Users",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "description": "It gets the user information details by ID\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Gets a user by ID.",
        "operationId": "getUserbyID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      },
      "delete": {
        "description": "It delets the user information details by ID\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Deletes a user by ID.",
        "operationId": "deleteUserbyID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/Status"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      },
      "patch": {
        "description": "It updates the user information details by ID\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Updates a user by ID.",
        "operationId": "updateUserbyID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "User Body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "409": {
            "$ref": "#/responses/UserExists"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "status": {
          "type": "integer"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "id",
        "first_name",
        "last_name",
        "email_address",
        "mobile_no",
        "dob",
        "sex",
        "username",
        "password",
        "location"
      ],
      "properties": {
        "created_at": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "dob": {
          "type": "string",
          "format": "date"
        },
        "email_address": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "is_active": {
          "type": "boolean"
        },
        "is_staying_here": {
          "type": "boolean"
        },
        "last_name": {
          "type": "string"
        },
        "location": {
          "type": "string"
        },
        "mobile_no": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "role_name": {
          "type": "string"
        },
        "sex": {
          "type": "string"
        },
        "updated_at": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "The api is Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "InternalServerError": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "NotFound": {
      "description": "The api is not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "Status": {
      "description": "API Custom Status",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "UnauthorizedError": {
      "description": "The api is Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "UserExists": {
      "description": "The user already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server in Go with Swagger endpoints definition",
    "title": "social-naka-app",
    "version": "0.0.1"
  },
  "paths": {
    "/registerUser": {
      "post": {
        "description": "It registers the User\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Registers the User",
        "operationId": "registerUser",
        "parameters": [
          {
            "name": "User Body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "The api is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "The user already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "It gets all the Users\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Get all the Users",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "400": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "The api is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "description": "It gets the user information details by ID\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Gets a user by ID.",
        "operationId": "getUserbyID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "The api is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "It delets the user information details by ID\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Deletes a user by ID.",
        "operationId": "deleteUserbyID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "API Custom Status",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "400": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "The api is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "description": "It updates the user information details by ID\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Updates a user by ID.",
        "operationId": "updateUserbyID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "User Body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "400": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "The api is Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "The api is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "The user already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "status": {
          "type": "integer"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "id",
        "first_name",
        "last_name",
        "email_address",
        "mobile_no",
        "dob",
        "sex",
        "username",
        "password",
        "location"
      ],
      "properties": {
        "created_at": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "dob": {
          "type": "string",
          "format": "date"
        },
        "email_address": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "is_active": {
          "type": "boolean"
        },
        "is_staying_here": {
          "type": "boolean"
        },
        "last_name": {
          "type": "string"
        },
        "location": {
          "type": "string"
        },
        "mobile_no": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "role_name": {
          "type": "string"
        },
        "sex": {
          "type": "string"
        },
        "updated_at": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "The api is Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "InternalServerError": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "NotFound": {
      "description": "The api is not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "Status": {
      "description": "API Custom Status",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "UnauthorizedError": {
      "description": "The api is Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "UserExists": {
      "description": "The user already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  }
}`))
}
