// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/blocks/latest": {
            "get": {
                "description": "Finds the latest block and returns it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get the latest block",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.BlockInfoResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/blocks/{id}": {
            "get": {
                "description": "Finds a block by given hash or number and returns it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get block info by hash or by number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block Hash or Number",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.BlockInfoResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups": {
            "get": {
                "description": "Returns list of groups",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.ListGroupsResponse"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/groups/{id}": {
            "get": {
                "description": "Finds a group by ID and returns it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get group info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.GroupInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/indexes/{id}": {
            "get": {
                "description": "Finds an index by ID and returns it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get index info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Index ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.IndexInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "rest.BlockInfoResponse": {
            "type": "object",
            "properties": {
                "gas_used": {
                    "type": "integer"
                },
                "hash": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "string"
                },
                "transactions_count": {
                    "type": "integer"
                }
            }
        },
        "rest.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "rest.GroupInfoResponse": {
            "type": "object",
            "properties": {
                "indexes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "rest.IndexInfoResponse": {
            "type": "object",
            "properties": {
                "eth_price_in_wei": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "percentage_change": {
                    "type": "integer"
                },
                "usd_capitalization": {
                    "type": "integer"
                },
                "usd_price_in_cents": {
                    "type": "integer"
                }
            }
        },
        "rest.ListGroupsResponse": {
            "type": "object",
            "properties": {
                "groups": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "v1.0.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Tech assignment API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}