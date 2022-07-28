// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "name": "GM-Center",
            "url": "https://gedge-platform.github.io/gm-center/"
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
        "/auth": {
            "post": {
                "description": "get JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User Info Body",
                        "name": "authBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    }
                }
            }
        },
        "/cluster/{name}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get cluster Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show detail cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the Cluster",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/cronjob/{name}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get cronjob Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show detail cronjob",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the Cronjob",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cluster Name of the Cronjob",
                        "name": "cluster",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CRONJOB"
                        }
                    }
                }
            }
        },
        "/cronjobs": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get cronjob List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show List cronjob",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CRONJOB"
                        }
                    }
                }
            }
        },
        "/job/:name": {
            "get": {
                "description": "get job Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show detail job",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.JOB"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    }
                }
            }
        },
        "/jobs": {
            "get": {
                "description": "get job List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show List job",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.JOB"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    }
                }
            }
        },
        "/pods": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get pods List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show List pods",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cluster Name of the pods",
                        "name": "cluster",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "workspace Name of the pods",
                        "name": "workspace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.POD"
                        }
                    }
                }
            }
        },
        "/pods/{name}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get pods Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show detail pods",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the pods",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cluster Name of the pods",
                        "name": "cluster",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "workspace Name of the pods",
                        "name": "workspace",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.POD"
                        }
                    }
                }
            }
        },
        "/pvcs": {
            "get": {
                "description": "get pvc List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show app PVCs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PVC"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    }
                }
            }
        },
        "/pvs": {
            "get": {
                "description": "get pv List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show app PVs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PV"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    }
                }
            }
        },
        "/pvs/:name": {
            "get": {
                "responses": {}
            }
        },
        "/spider/cloudos": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get CloudOS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cloudos",
                "responses": {}
            }
        },
        "/spider/credentials": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get ALLCredential",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Credential",
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get Credential",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Credential",
                "parameters": [
                    {
                        "description": "Credential Info Body",
                        "name": "CredentialBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/spider/credentials/{credentialName}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get Credential",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Credential",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the credentials",
                        "name": "credentialName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get Credential",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Credential",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the credentials",
                        "name": "credentialName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Active": {
            "type": "object",
            "properties": {
                "kind": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        },
        "model.CRONJOB": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Active"
                    }
                },
                "annotations": {},
                "cluster": {
                    "type": "string"
                },
                "concurrencyPolicy": {
                    "type": "string"
                },
                "containers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Containers"
                    }
                },
                "creationTimestamp": {
                    "type": "string"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EVENT"
                    }
                },
                "failedJobsHistoryLimit": {
                    "type": "integer"
                },
                "label": {},
                "lastScheduleTime": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "schedule": {
                    "type": "string"
                },
                "successfulJobsHistoryLimit": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                },
                "workspace": {
                    "type": "string"
                }
            }
        },
        "model.ConfigMapKeyRef": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ContainerStatuses": {
            "type": "object",
            "properties": {
                "containerID": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ready": {
                    "type": "boolean"
                },
                "restartCount": {
                    "type": "integer"
                },
                "started": {
                    "type": "boolean"
                }
            }
        },
        "model.Containers": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ENV": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                },
                "valueFrom": {
                    "$ref": "#/definitions/model.ValueFrom"
                }
            }
        },
        "model.EVENT": {
            "type": "object",
            "properties": {
                "cluster": {
                    "type": "string"
                },
                "eventTime": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.FieldRef": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string"
                },
                "fieldPath": {
                    "type": "string"
                }
            }
        },
        "model.JOB": {
            "type": "object",
            "properties": {
                "cluster": {
                    "type": "string"
                },
                "completions": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "duration": {},
                "name": {
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "workspace": {
                    "type": "string"
                }
            }
        },
        "model.OwnerReference": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.POD": {
            "type": "object",
            "properties": {
                "Podcontainers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PODCONTAINERS"
                    }
                },
                "annotations": {},
                "cluster": {
                    "type": "string"
                },
                "containerStatuses": {
                    "description": "VolumeMounts      []VolumeMounts      ` + "`" + `json:\"volumemounts\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerStatuses"
                    }
                },
                "creationTimestamp": {
                    "type": "string"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EVENT"
                    }
                },
                "hostIP": {
                    "type": "string"
                },
                "label": {},
                "name": {
                    "type": "string"
                },
                "node_name": {
                    "type": "string"
                },
                "ownerReferences": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OwnerReference"
                    }
                },
                "podIP": {
                    "type": "string"
                },
                "podIPs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PodIPs"
                    }
                },
                "project": {
                    "type": "string"
                },
                "qosClass": {
                    "type": "string"
                },
                "restart": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "statusConditions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StatusConditions"
                    }
                },
                "user": {
                    "type": "string"
                },
                "workspace": {
                    "type": "string"
                }
            }
        },
        "model.PODCONTAINERS": {
            "type": "object",
            "properties": {
                "env": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ENV"
                    }
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ports": {
                    "description": "ReadinessProbe ReadinessProbe ` + "`" + `json:\"readinessProbe\",omitempty` + "`" + `\nLivenessProbe  LivenessProbe  ` + "`" + `json:\"livenessProbe\",omitempty` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Ports"
                    }
                },
                "volumemounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VolumeMounts"
                    }
                }
            }
        },
        "model.PV": {
            "type": "object",
            "properties": {
                "accessMode": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "annotations": {},
                "capacity": {
                    "type": "string"
                },
                "claim": {},
                "cluster": {
                    "type": "string"
                },
                "createAt": {
                    "description": "Workspace string ` + "`" + `json:\"workspace\"` + "`" + `",
                    "type": "string"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EVENT"
                    }
                },
                "label": {},
                "name": {
                    "type": "string"
                },
                "reclaimPolicy": {
                    "type": "string"
                },
                "status": {},
                "storageClass": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "volumeMode": {
                    "description": "Reason        []EVENT          ` + "`" + `json:\"events\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "model.PVC": {
            "type": "object",
            "properties": {
                "accessMode": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "annotations": {},
                "capacity": {
                    "type": "string"
                },
                "clusterName": {
                    "type": "string"
                },
                "createAt": {
                    "type": "string"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EVENT"
                    }
                },
                "finalizers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "label": {
                    "description": "Reason        []EVENT          ` + "`" + `json:\"events\"` + "`" + `"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "status": {},
                "storageClass": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "volume": {},
                "workspaceName": {
                    "type": "string"
                }
            }
        },
        "model.PodIPs": {
            "type": "object",
            "properties": {
                "ip": {
                    "type": "string"
                }
            }
        },
        "model.Ports": {
            "type": "object",
            "properties": {
                "containerPort": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "string"
                }
            }
        },
        "model.StatusConditions": {
            "type": "object",
            "properties": {
                "lastTransitionTime": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "string"
                },
                "Password": {
                    "description": "Email    string ` + "`" + `json:\"email\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "model.ValueFrom": {
            "type": "object",
            "properties": {
                "configMapKeyRef": {
                    "$ref": "#/definitions/model.ConfigMapKeyRef"
                },
                "fieldRef": {
                    "$ref": "#/definitions/model.FieldRef"
                }
            }
        },
        "model.VolumeMounts": {
            "type": "object",
            "properties": {
                "mountpath": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "readonly": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "\"Type \\\"Bearer \\\" and then your API Token\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "192.168.160.216:8010",
	BasePath:         "/gmcapi/v2",
	Schemes:          []string{"http"},
	Title:            "Gedge GM-Center Swagger API",
	Description:      "This is a Gedge GM-Center Swagger API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
