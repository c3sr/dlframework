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
    "/v1/predict/close": {
      "post": {
        "tags": [
          "Predict"
        ],
        "summary": "Close a predictor clear it's memory.",
        "operationId": "Close",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictor"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictorCloseResponse"
            }
          }
        }
      }
    },
    "/v1/predict/dataset": {
      "post": {
        "description": "The result is a prediction feature list.",
        "tags": [
          "Predict"
        ],
        "summary": "Dataset method receives a single dataset and runs\nthe predictor on all elements of the dataset.",
        "operationId": "Dataset",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkDatasetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkFeaturesResponse"
            }
          }
        }
      }
    },
    "/v1/predict/images": {
      "post": {
        "description": "The result is a prediction feature list for each image.",
        "tags": [
          "Predict"
        ],
        "summary": "Image method receives a list base64 encoded images and runs\nthe predictor on all the images.",
        "operationId": "Images",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkImagesRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkFeaturesResponse"
            }
          }
        }
      }
    },
    "/v1/predict/open": {
      "post": {
        "tags": [
          "Predict"
        ],
        "summary": "Opens a predictor and returns an id where the predictor\nis accessible. The id can be used to perform inference\nrequests.",
        "operationId": "Open",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictorOpenRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkPredictor"
            }
          }
        }
      }
    },
    "/v1/predict/reset": {
      "post": {
        "tags": [
          "Predict"
        ],
        "summary": "Clear method clears the internal cache of the predictors",
        "operationId": "Reset",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkResetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkResetResponse"
            }
          }
        }
      }
    },
    "/v1/predict/stream/dataset": {
      "post": {
        "description": "The result is a prediction feature stream.",
        "tags": [
          "Predict"
        ],
        "summary": "Dataset method receives a single dataset and runs\nthe predictor on all elements of the dataset.",
        "operationId": "DatasetStream",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkDatasetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkFeatureResponse"
            }
          }
        }
      }
    },
    "/v1/predict/stream/images": {
      "post": {
        "description": "The result is a prediction feature stream for each image.",
        "tags": [
          "Predict"
        ],
        "summary": "Image method receives a list base64 encoded images and runs\nthe predictor on all the images.",
        "operationId": "ImagesStream",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkImagesRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkFeatureResponse"
            }
          }
        }
      }
    },
    "/v1/predict/stream/urls": {
      "post": {
        "description": "The result is a prediction feature stream for each url.",
        "tags": [
          "Predict"
        ],
        "summary": "Image method receives a stream of urls and runs\nthe predictor on all the urls. The",
        "operationId": "URLsStream",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkURLsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/dlframeworkFeatureResponse"
            }
          }
        }
      }
    },
    "/v1/predict/urls": {
      "post": {
        "description": "The result is a prediction feature stream for each url.",
        "tags": [
          "Predict"
        ],
        "summary": "Image method receives a stream of urls and runs\nthe predictor on all the urls. The",
        "operationId": "URLs",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dlframeworkURLsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/dlframeworkFeaturesResponse"
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
    "DatasetRequestDataset": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ImagesRequestImage": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "format": "byte",
          "title": "The image is base64 encoded"
        },
        "id": {
          "type": "string",
          "title": "An id used to identify the output feature: maps to input_id for output"
        },
        "preprocessed": {
          "type": "boolean",
          "format": "boolean",
          "title": "Preprocessed is set to true to disable preprocessing.\nIf enabled then the image is assumed to be rescaled and\nencoded as an array of float32 values"
        }
      }
    },
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
    "URLsRequestURL": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "title": "An id used to identify the output feature: maps to input_id for output"
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
        },
        "specification": {
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
    "dlframeworkDatasetRequest": {
      "type": "object",
      "properties": {
        "dataset": {
          "$ref": "#/definitions/DatasetRequestDataset"
        },
        "options": {
          "$ref": "#/definitions/dlframeworkPredictionOptions"
        },
        "predictor": {
          "$ref": "#/definitions/dlframeworkPredictor"
        }
      }
    },
    "dlframeworkFeature": {
      "type": "object",
      "properties": {
        "index": {
          "type": "string",
          "format": "int64"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
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
    "dlframeworkFeatureResponse": {
      "type": "object",
      "properties": {
        "features": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkFeature"
          }
        },
        "id": {
          "type": "string"
        },
        "input_id": {
          "type": "string"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "request_id": {
          "type": "string"
        }
      }
    },
    "dlframeworkFeaturesResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "responses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dlframeworkFeatureResponse"
          }
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
    "dlframeworkImagesRequest": {
      "type": "object",
      "properties": {
        "images": {
          "type": "array",
          "title": "A list of Base64 encoded images",
          "items": {
            "$ref": "#/definitions/ImagesRequestImage"
          }
        },
        "options": {
          "$ref": "#/definitions/dlframeworkPredictionOptions"
        },
        "predictor": {
          "$ref": "#/definitions/dlframeworkPredictor"
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
    "dlframeworkPredictionOptions": {
      "type": "object",
      "properties": {
        "feature_limit": {
          "type": "integer",
          "format": "int32"
        },
        "request_id": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictor": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "dlframeworkPredictorCloseResponse": {
      "type": "object"
    },
    "dlframeworkPredictorOpenRequest": {
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
    "dlframeworkResetRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "predictor": {
          "$ref": "#/definitions/dlframeworkPredictor"
        }
      }
    },
    "dlframeworkResetResponse": {
      "type": "object",
      "properties": {
        "predictor": {
          "$ref": "#/definitions/dlframeworkPredictor"
        }
      }
    },
    "dlframeworkURLsRequest": {
      "type": "object",
      "properties": {
        "options": {
          "$ref": "#/definitions/dlframeworkPredictionOptions"
        },
        "predictor": {
          "$ref": "#/definitions/dlframeworkPredictor"
        },
        "urls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/URLsRequestURL"
          }
        }
      }
    }
  },
  "externalDocs": {
    "url": "https://rai-project.github.io/carml"
  }
}`))
}
