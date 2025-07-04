package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// IPLocationResponse represents the response from ip-api.com
type IPLocationResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
	Message     string  `json:"message,omitempty"`
}

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"IP Geo Location Server 🌍",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// Add IP geolocation tool
	tool := mcp.NewTool("get_ip_location",
		mcp.WithDescription("Get geographical location information for an IP address"),
		mcp.WithString("ip",
			mcp.Required(),
			mcp.Description("IP address to lookup (IPv4 or IPv6). Use 'auto' to get location of current IP"),
		),
	)

	// Add tool handler
	s.AddTool(tool, ipLocationHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func ipLocationHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ip, err := request.RequireString("ip")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Build API URL
	var apiURL string
	if ip == "auto" || ip == "" {
		// Get current IP location
		apiURL = "http://ip-api.com/json/"
	} else {
		// Get specific IP location
		apiURL = fmt.Sprintf("http://ip-api.com/json/%s", ip)
	}

	// Make GET request
	resp, err := client.Get(apiURL)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to make request: %v", err)), nil
	}
	defer resp.Body.Close()

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return mcp.NewToolResultError(fmt.Sprintf("API request failed with status: %d", resp.StatusCode)), nil
	}

	// Parse JSON response
	var locationData IPLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationData); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to parse response: %v", err)), nil
	}

	// Check if the API returned an error
	if locationData.Status == "fail" {
		return mcp.NewToolResultError(fmt.Sprintf("API error: %s", locationData.Message)), nil
	}

	// Format the response
	result := formatLocationResult(locationData)
	return mcp.NewToolResultText(result), nil
}

func formatLocationResult(data IPLocationResponse) string {
	return fmt.Sprintf("country: %s \n regionName: %s \n city: %s", data.Country, data.RegionName, data.City)
}
