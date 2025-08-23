package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/twilio-voice/mcp-server/config"
	"github.com/twilio-voice/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListdialingpermissionshrsprefixesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["PageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("PageSize=%v", val))
		}
		if val, ok := args["Page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Page=%v", val))
		}
		if val, ok := args["PageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("PageToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/DialingPermissions/Countries/%s/HighRiskSpecialPrefixes%s", cfg.BaseURL, IsoCode, queryString)
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

func CreateListdialingpermissionshrsprefixesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_DialingPermissions_Countries_IsoCode_HighRiskSpecialPrefixes",
		mcp.WithDescription("Fetch the high-risk special services prefixes from the country resource corresponding to the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)"),
		mcp.WithString("IsoCode", mcp.Required(), mcp.Description("The [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) to identify the country permissions from which high-risk special service number prefixes are fetched")),
		mcp.WithNumber("PageSize", mcp.Description("How many resources to return in each list page. The default is 50, and the maximum is 1000.")),
		mcp.WithNumber("Page", mcp.Description("The page index. This value is simply for client state.")),
		mcp.WithString("PageToken", mcp.Description("The page token. This is provided by the API.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListdialingpermissionshrsprefixesHandler(cfg),
	}
}
