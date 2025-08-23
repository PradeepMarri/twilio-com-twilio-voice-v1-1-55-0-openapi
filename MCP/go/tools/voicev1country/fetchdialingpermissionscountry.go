package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/twilio-voice/mcp-server/config"
	"github.com/twilio-voice/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func FetchdialingpermissionscountryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		IsoCodeVal, ok := args["IsoCode"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: IsoCode"), nil
		}
		IsoCode, ok := IsoCodeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: IsoCode"), nil
		}
		url := fmt.Sprintf("%s/v1/DialingPermissions/Countries/%s", cfg.BaseURL, IsoCode)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
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
		var result models.GeneratedType
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

func CreateFetchdialingpermissionscountryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_DialingPermissions_Countries_IsoCode",
		mcp.WithDescription("Retrieve voice dialing country permissions identified by the given ISO country code"),
		mcp.WithString("IsoCode", mcp.Required(), mcp.Description("The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the DialingPermissions Country resource to fetch")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    FetchdialingpermissionscountryHandler(cfg),
	}
}
