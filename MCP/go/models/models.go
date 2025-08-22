package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetDeploymentsRequest represents the GetDeploymentsRequest schema from the OpenAPI specification
type GetDeploymentsRequest struct {
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Devicename interface{} `json:"DeviceName"`
}

// Definition represents the Definition schema from the OpenAPI specification
type Definition struct {
	Modelhandle interface{} `json:"ModelHandle,omitempty"`
	S3url interface{} `json:"S3Url,omitempty"`
	State interface{} `json:"State,omitempty"`
	Checksum interface{} `json:"Checksum,omitempty"`
}

// DeploymentModel represents the DeploymentModel schema from the OpenAPI specification
type DeploymentModel struct {
	Modelhandle interface{} `json:"ModelHandle,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Modelversion interface{} `json:"ModelVersion,omitempty"`
	Rollbackfailurereason interface{} `json:"RollbackFailureReason,omitempty"`
	State interface{} `json:"State,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusreason interface{} `json:"StatusReason,omitempty"`
	Desiredstate interface{} `json:"DesiredState,omitempty"`
}

// Model represents the Model schema from the OpenAPI specification
type Model struct {
	Latestinference interface{} `json:"LatestInference,omitempty"`
	Latestsampletime interface{} `json:"LatestSampleTime,omitempty"`
	Modelmetrics interface{} `json:"ModelMetrics,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Modelversion interface{} `json:"ModelVersion,omitempty"`
}

// EdgeDeployment represents the EdgeDeployment schema from the OpenAPI specification
type EdgeDeployment struct {
	TypeField interface{} `json:"Type,omitempty"`
	Definitions interface{} `json:"Definitions,omitempty"`
	Deploymentname interface{} `json:"DeploymentName,omitempty"`
	Failurehandlingpolicy interface{} `json:"FailureHandlingPolicy,omitempty"`
}

// GetDeviceRegistrationResult represents the GetDeviceRegistrationResult schema from the OpenAPI specification
type GetDeviceRegistrationResult struct {
	Cachettl interface{} `json:"CacheTTL,omitempty"`
	Deviceregistration interface{} `json:"DeviceRegistration,omitempty"`
}

// Checksum represents the Checksum schema from the OpenAPI specification
type Checksum struct {
	TypeField interface{} `json:"Type,omitempty"`
	Sum interface{} `json:"Sum,omitempty"`
}

// SendHeartbeatRequest represents the SendHeartbeatRequest schema from the OpenAPI specification
type SendHeartbeatRequest struct {
	Deploymentresult interface{} `json:"DeploymentResult,omitempty"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
	Devicename interface{} `json:"DeviceName"`
	Models interface{} `json:"Models,omitempty"`
	Agentmetrics interface{} `json:"AgentMetrics,omitempty"`
	Agentversion interface{} `json:"AgentVersion"`
}

// DeploymentResult represents the DeploymentResult schema from the OpenAPI specification
type DeploymentResult struct {
	Deploymentstatusmessage interface{} `json:"DeploymentStatusMessage,omitempty"`
	Deploymentendtime interface{} `json:"DeploymentEndTime,omitempty"`
	Deploymentmodels interface{} `json:"DeploymentModels,omitempty"`
	Deploymentname interface{} `json:"DeploymentName,omitempty"`
	Deploymentstarttime interface{} `json:"DeploymentStartTime,omitempty"`
	Deploymentstatus interface{} `json:"DeploymentStatus,omitempty"`
}

// GetDeviceRegistrationRequest represents the GetDeviceRegistrationRequest schema from the OpenAPI specification
type GetDeviceRegistrationRequest struct {
	Devicename interface{} `json:"DeviceName"`
	Devicefleetname interface{} `json:"DeviceFleetName"`
}

// GetDeploymentsResult represents the GetDeploymentsResult schema from the OpenAPI specification
type GetDeploymentsResult struct {
	Deployments interface{} `json:"Deployments,omitempty"`
}

// EdgeMetric represents the EdgeMetric schema from the OpenAPI specification
type EdgeMetric struct {
	Value interface{} `json:"Value,omitempty"`
	Dimension interface{} `json:"Dimension,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Timestamp interface{} `json:"Timestamp,omitempty"`
}
