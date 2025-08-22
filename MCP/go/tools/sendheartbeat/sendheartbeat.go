package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-sagemaker-edge-manager/mcp-server/config"
	"github.com/amazon-sagemaker-edge-manager/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func SendheartbeatHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/SendHeartbeat", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateSendheartbeatTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_SendHeartbeat",
		mcp.WithDescription("Use to get the current status of devices registered on SageMaker Edge Manager."),
		mcp.WithObject("DeploymentResult", mcp.Description("Input parameter: Information about the result of a deployment on an edge device that is registered with SageMaker Edge Manager.")),
		mcp.WithString("DeviceFleetName", mcp.Required(), mcp.Description("Input parameter: The name of the fleet that the device belongs to.")),
		mcp.WithString("DeviceName", mcp.Required(), mcp.Description("Input parameter: The unique name of the device.")),
		mcp.WithArray("Models", mcp.Description("Input parameter: Returns a list of models deployed on the the device.")),
		mcp.WithArray("AgentMetrics", mcp.Description("Input parameter: For internal use. Returns a list of SageMaker Edge Manager agent operating metrics.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SendheartbeatHandler(cfg),
	}
}
