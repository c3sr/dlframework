// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

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
    "title": "dlframework.proto",
    "contact": {
      "name": "Abdul Dakkak, Cheng Li",
      "url": "https://github.com/rai-project/carml"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "paths": {
    "/v1/predict/dataset": {
      "post": {
        "description": "The result is a prediction feature stream.",
        "tags": [
          "Predictor"
        ],
        "summary": "Dataset method receives a single dataset and runs\nthe predictor on all elements of the dataset.",
        "operationId": "Dataset",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictDatasetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictionFeatureResponse"
            }
          }
        }
      }
    },
    "/v1/predict/images": {
      "post": {
        "description": "The result is a prediction feature stream for each image.",
        "tags": [
          "Predictor"
        ],
        "summary": "Image method receives a stream of images and runs\nthe predictor on all the images.",
        "operationId": "Images",
        "parameters": [
          {
            "description": "(streaming inputs)",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictImageRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictionFeatureResponse"
            }
          }
        }
      }
    },
    "/v1/predict/urls": {
      "post": {
        "description": "The result is a prediction feature stream for each url.",
        "tags": [
          "Predictor"
        ],
        "summary": "Image method receives a stream of urls and runs\nthe predictor on all the urls. The",
        "operationId": "URLs",
        "parameters": [
          {
            "description": "(streaming inputs)",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictURLRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictionFeatureResponse"
            }
          }
        }
      }
    },
    "/v1/registry/frameworks/agent": {
      "get": {
        "tags": [
          "Registry"
        ],
        "operationId": "FrameworkAgents",
        "parameters": [
          {
            "type": "string",
            "name": "framework_name",
            "in": "query"
          },
          {
            "type": "string",
            "name": "framework_version",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkAgents"
            }
          }
        }
      }
    },
    "/v1/registry/frameworks/manifest": {
      "get": {
        "tags": [
          "Registry"
        ],
        "operationId": "FrameworkManifests",
        "parameters": [
          {
            "type": "string",
            "name": "framework_name",
            "in": "query"
          },
          {
            "type": "string",
            "name": "framework_version",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkFrameworkManifestsResponse"
            }
          }
        }
      }
    },
    "/v1/registry/models/agent": {
      "get": {
        "tags": [
          "Registry"
        ],
        "operationId": "ModelAgents",
        "parameters": [
          {
            "type": "string",
            "name": "framework_name",
            "in": "query"
          },
          {
            "type": "string",
            "name": "framework_version",
            "in": "query"
          },
          {
            "type": "string",
            "name": "model_name",
            "in": "query"
          },
          {
            "type": "string",
            "name": "model_version",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkAgents"
            }
          }
        }
      }
    },
    "/v1/registry/models/manifest": {
      "get": {
        "tags": [
          "Registry"
        ],
        "operationId": "ModelManifests",
        "parameters": [
          {
            "type": "string",
            "name": "framework_name",
            "in": "query"
          },
          {
            "type": "string",
            "name": "framework_version",
            "in": "query"
          },
          {
            "type": "string",
            "name": "model_name",
            "in": "query"
          },
          {
            "type": "string",
            "name": "model_version",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkModelManifestsResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ModelManifestModel": {
      "type": "object",
      "properties": {
        "base_url": {
          "type": "string"
        },
        "graph_path": {
          "type": "string"
        },
        "is_archive": {
          "type": "boolean",
          "format": "boolean"
        },
        "weights_path": {
          "type": "string"
        }
      }
    },
    "TypeParameter": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    },
    "dlframeworkAgent": {
      "type": "object",
      "properties": {
        "host": {
          "type": "string"
        },
        "port": {
          "type": "string"
        }
      }
    },
    "dlframeworkAgents": {
      "type": "object",
      "properties": {
        "agents": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkAgent"
          }
        }
      }
    },
    "dlframeworkContainerHardware": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "string"
        },
        "gpu": {
          "type": "string"
        }
      }
    },
    "dlframeworkFrameworkManifest": {
      "type": "object",
      "properties": {
        "container": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/dlframeworkContainerHardware"
          }
        },
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "dlframeworkFrameworkManifestsResponse": {
      "type": "object",
      "properties": {
        "manifests": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkFrameworkManifest"
          }
        }
      }
    },
    "dlframeworkFrameworkRequest": {
      "type": "object",
      "properties": {
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        }
      }
    },
    "dlframeworkModelManifest": {
      "type": "object",
      "properties": {
        "after_postprocess": {
          "type": "string"
        },
        "after_preprocess": {
          "type": "string"
        },
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "before_postprocess": {
          "type": "string"
        },
        "before_preprocess": {
          "type": "string"
        },
        "container": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/dlframeworkContainerHardware"
          }
        },
        "description": {
          "type": "string"
        },
        "framework": {
          "$ref": "#/definitions/dlframeworkFrameworkManifest"
        },
        "hidden": {
          "type": "boolean",
          "format": "boolean"
        },
        "inputs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkModelManifestType"
          }
        },
        "license": {
          "type": "string"
        },
        "model": {
          "$ref": "#/definitions/ModelManifestModel"
        },
        "name": {
          "type": "string"
        },
        "output": {
          "$ref": "#/definitions/dlframeworkModelManifestType"
        },
        "postprocess": {
          "type": "string"
        },
        "preprocess": {
          "type": "string"
        },
        "reference": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "version": {
          "type": "string"
        }
      }
    },
    "dlframeworkModelManifestType": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "parameters": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/TypeParameter"
          }
        },
        "type": {
          "type": "string"
        }
      }
    },
    "dlframeworkModelManifestsResponse": {
      "type": "object",
      "properties": {
        "manifests": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkModelManifest"
          }
        }
      }
    },
    "dlframeworkModelRequest": {
      "type": "object",
      "properties": {
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictDatasetRequest": {
      "type": "object",
      "properties": {
        "dataset_category": {
          "type": "string"
        },
        "dataset_name": {
          "type": "string"
        },
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "limit": {
          "type": "integer",
          "format": "int32"
        },
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictImageRequest": {
      "type": "object",
      "properties": {
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "image": {
          "type": "string",
          "format": "byte",
          "title": "Base64 encoded image"
        },
        "input_id": {
          "type": "string"
        },
        "limit": {
          "type": "integer",
          "format": "int32"
        },
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictURLRequest": {
      "type": "object",
      "properties": {
        "framework_name": {
          "type": "string"
        },
        "framework_version": {
          "type": "string"
        },
        "input_id": {
          "type": "string"
        },
        "limit": {
          "type": "integer",
          "format": "int32"
        },
        "model_name": {
          "type": "string"
        },
        "model_version": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictionFeature": {
      "type": "object",
      "properties": {
        "index": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "probability": {
          "type": "number",
          "format": "float"
        }
      }
    },
    "dlframeworkPredictionFeatureResponse": {
      "type": "object",
      "properties": {
        "feature": {
          "$ref": "#/definitions/dlframeworkPredictionFeature"
        },
        "id": {
          "type": "string"
        },
        "input_id": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      }
    }
  },
  "externalDocs": {
    "url": "https://rai-project.github.io/carml"
  }
}`))
}
