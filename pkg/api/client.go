package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/CryptoRodeo/issues-cli/pkg/config"
	"github.com/CryptoRodeo/issues-cli/pkg/models"
)

// Client is the API client for the Konflux issues API
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// New creates a new API client
func New() *Client {
	cfg := config.GetConfig()
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: cfg.APIUrl,
	}
}

// GetIssues retrieves issues with optional filters
func (c *Client) GetIssues(namespace string, filters map[string]string) ([]models.Issue, error) {
	// Build query parameters
	params := url.Values{}
	params.Add("namespace", namespace)
	for key, value := range filters {
		if value != "" {
			params.Add(key, value)
		}
	}

	// Make request
	url := fmt.Sprintf("%s/issues?%s", c.baseURL, params.Encode())
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get issues: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var response models.IssuesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse issues: %w", err)
	}

	return response.Data, nil
}

// GetIssueDetails retrieves details for a specific issue
func (c *Client) GetIssueDetails(id, namespace string) (*models.Issue, error) {
	// Build query parameters
	params := url.Values{}
	params.Add("namespace", namespace)

	// Make request
	url := fmt.Sprintf("%s/issues/%s?%s", c.baseURL, id, params.Encode())
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get issue details: %w", err)
	}
	defer resp.Body.Close()

	// Handle not found and access denied responses
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("issue with ID %s not found", id)
	}
	if resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("access denied to namespace %s", namespace)
	}

	// Check other response statuses
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var issue models.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, fmt.Errorf("failed to parse issue details: %w", err)
	}

	return &issue, nil
}
