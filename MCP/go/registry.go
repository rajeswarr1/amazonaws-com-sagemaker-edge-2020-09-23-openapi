package main

import (
	"github.com/amazon-sagemaker-edge-manager/mcp-server/config"
	"github.com/amazon-sagemaker-edge-manager/mcp-server/models"
	tools_getdeviceregistration "github.com/amazon-sagemaker-edge-manager/mcp-server/tools/getdeviceregistration"
	tools_sendheartbeat "github.com/amazon-sagemaker-edge-manager/mcp-server/tools/sendheartbeat"
	tools_getdeployments "github.com/amazon-sagemaker-edge-manager/mcp-server/tools/getdeployments"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_getdeviceregistration.CreateGetdeviceregistrationTool(cfg),
		tools_sendheartbeat.CreateSendheartbeatTool(cfg),
		tools_getdeployments.CreateGetdeploymentsTool(cfg),
	}
}
