// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/market/charts": {
            "get": {
                "description": "Get the charts data from an market and coin/token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charts"
                ],
                "summary": "Get charts data for a specific coin",
                "operationId": "get_charts_data",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 60,
                        "description": "Coin id",
                        "name": "coin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token id",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1574483028,
                        "description": "Start timestamp",
                        "name": "time_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 64,
                        "description": "Max number of items in result prices array",
                        "name": "max_items",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show charts",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.Chart"
                        }
                    }
                }
            }
        },
        "/v1/market/info": {
            "get": {
                "description": "Get the charts coin assets data from an market and coin/contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "Get charts coin assets data for a specific coin",
                "operationId": "get_charts_coin_info",
                "parameters": [
                    {
                        "type": "string",
                        "default": "60",
                        "description": "Coin id",
                        "name": "coin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token id",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show coin assets in",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.CoinDetails"
                        }
                    }
                }
            }
        },
        "/v1/market/rate": {
            "get": {
                "description": "Get rate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rates"
                ],
                "summary": "Get rate",
                "operationId": "get_rate",
                "parameters": [
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "From",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "RUB",
                        "description": "To",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "100",
                        "description": "Amount",
                        "name": "amount",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.RateResponse"
                        }
                    }
                }
            }
        },
        "/v1/market/ticker": {
            "post": {
                "description": "Get the ticker values from many market and coin/token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickers"
                ],
                "summary": "Get ticker values for a specific market",
                "operationId": "get_tickers",
                "parameters": [
                    {
                        "description": "Ticker",
                        "name": "tickers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.TickerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TickerResponse"
                        }
                    }
                }
            }
        },
        "/v2/market/charts/{id}": {
            "get": {
                "description": "Get the charts data from an market and coin/token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charts"
                ],
                "summary": "Get charts data for a specific id",
                "operationId": "get_charts_data_v2",
                "parameters": [
                    {
                        "type": "string",
                        "default": "c60",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1574483028,
                        "description": "Start timestamp",
                        "name": "time_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 64,
                        "description": "Max number of items in result prices array",
                        "name": "max_items",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show charts",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.Chart"
                        }
                    }
                }
            }
        },
        "/v2/market/info/{id}": {
            "get": {
                "description": "Get the charts coin assets data from an market and coin/contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "Get charts coin assets data for a specific coin",
                "operationId": "get_charts_coin_info_v2",
                "parameters": [
                    {
                        "type": "string",
                        "default": "c714",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show coin assets in",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.CoinDetails"
                        }
                    }
                }
            }
        },
        "/v2/market/ticker/{id}": {
            "get": {
                "description": "Get the ticker for specific id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickers"
                ],
                "summary": "Get ticker for a specific market",
                "operationId": "get_ticker",
                "parameters": [
                    {
                        "type": "string",
                        "default": "c714_tXRP-BF2",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show coin assets in",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TickerResponseV2"
                        }
                    }
                }
            }
        },
        "/v2/market/tickers": {
            "post": {
                "description": "Get the tickers for list of ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickers"
                ],
                "summary": "Get tickers for list of ids",
                "operationId": "post_tickers_v2",
                "parameters": [
                    {
                        "description": "Ticker",
                        "name": "tickers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.TickerRequestV2"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TickerResponseV2"
                        }
                    }
                }
            }
        },
        "/v2/market/tickers/{assets}": {
            "get": {
                "description": "Get the tickers for list of ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickers"
                ],
                "summary": "Get tickers for list of ids",
                "operationId": "get_tickers_v2",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List of asset ids",
                        "name": "assets",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency symbol",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TickerResponseV2"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Coin": {
            "type": "object",
            "properties": {
                "Coin": {
                    "type": "integer"
                },
                "token_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controllers.RateResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "controllers.TickerPrice": {
            "type": "object",
            "properties": {
                "change_24h": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "provider": {
                    "type": "string"
                }
            }
        },
        "controllers.TickerRequest": {
            "type": "object",
            "properties": {
                "Currency": {
                    "type": "string"
                },
                "assets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.Coin"
                    }
                }
            }
        },
        "controllers.TickerRequestV2": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "currency": {
                    "type": "string"
                }
            }
        },
        "controllers.TickerResponse": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "docs": {
                    "type": "object",
                    "$ref": "#/definitions/watchmarket.Tickers"
                }
            }
        },
        "controllers.TickerResponseV2": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "tickers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.TickerPrice"
                    }
                }
            }
        },
        "watchmarket.Chart": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "prices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/watchmarket.ChartPrice"
                    }
                },
                "provider": {
                    "type": "string"
                }
            }
        },
        "watchmarket.ChartPrice": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "watchmarket.CoinDetails": {
            "type": "object",
            "properties": {
                "circulating_supply": {
                    "type": "number"
                },
                "info": {
                    "type": "object",
                    "$ref": "#/definitions/watchmarket.Info"
                },
                "market_cap": {
                    "type": "number"
                },
                "provider": {
                    "type": "string"
                },
                "provider_url": {
                    "type": "string"
                },
                "total_supply": {
                    "type": "number"
                },
                "volume_24": {
                    "type": "number"
                }
            }
        },
        "watchmarket.Info": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "explorer": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "research": {
                    "type": "string"
                },
                "short_description": {
                    "type": "string"
                },
                "socials": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/watchmarket.SocialLink"
                    }
                },
                "source_code": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                },
                "white_paper": {
                    "type": "string"
                }
            }
        },
        "watchmarket.Price": {
            "type": "object",
            "properties": {
                "change_24h": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "watchmarket.SocialLink": {
            "type": "object",
            "properties": {
                "handle": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "watchmarket.Ticker": {
            "type": "object",
            "properties": {
                "coin": {
                    "type": "integer"
                },
                "coin_name": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "last_update": {
                    "type": "string"
                },
                "market_cap": {
                    "type": "number"
                },
                "price": {
                    "type": "object",
                    "$ref": "#/definitions/watchmarket.Price"
                },
                "token_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "watchmarket.Tickers": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/watchmarket.Ticker"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
	swag.Register(swag.Name, &s{})
}
