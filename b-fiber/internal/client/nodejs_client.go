package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type NodeJSClient struct {
	BaseURL string
	Client  *http.Client
}

type StatsRequest struct {
	Rotated [][]float64 `json:"rotated"`
	Q       [][]float64 `json:"Q"`
	R       [][]float64 `json:"R"`
}

type StatsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Average    float64 `json:"average"`
	Sum        float64 `json:"sum"`
	IsDiagonal struct {
		Rotated bool `json:"rotated"`
		Q       bool `json:"Q"`
		R       bool `json:"R"`
	} `json:"isDiagonal"`
}

func NewNodeJSClient(baseURL string) *NodeJSClient {
	return &NodeJSClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetStatistics ahora acepta el token JWT
func (c *NodeJSClient) GetStatistics(req StatsRequest, token string) (*StatsResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.BaseURL+"/api/stats", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error calling Node.js API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Node.js API error (status %d): %s", resp.StatusCode, string(body))
	}

	var result struct {
		Success bool          `json:"success"`
		Data    StatsResponse `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if !result.Success {
		return nil, errors.New("Node.js API returned success=false")
	}

	return &result.Data, nil
}
