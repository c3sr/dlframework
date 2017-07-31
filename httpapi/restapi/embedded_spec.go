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
    "description": "CarML (Cognitive ARtifacts for Machine Learning) is a framework allowing people to develop and deploy machine learning models. It allows machine learning (ML) developers to publish and evaluate their models, users to experiment with different models and frameworks through a web user interface or a REST api, and system architects to capture system resource usage to inform future system and hardware configuration.",
    "title": "CarML DLFramework",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/predict": {
      "post": {
        "tags": [
          "Predictor"
        ],
        "operationId": "Predict",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictResponse"
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
    "dlframeworkErrorStatus": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "ok": {
          "type": "boolean",
          "format": "boolean"
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
    "dlframeworkPredictRequest": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "format": "byte"
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
        "url": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictResponse": {
      "type": "object",
      "properties": {
        "error": {
          "$ref": "#/definitions/dlframeworkErrorStatus"
        },
        "features": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkPredictionFeature"
          }
        },
        "id": {
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
    }
  }
}`))
}
