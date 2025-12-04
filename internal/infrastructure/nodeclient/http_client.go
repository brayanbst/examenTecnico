package nodeclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	appqr "github.com/brayanbst/matrix-service-go/internal/application/qr"
)

// HTTPStatsClient implementa StatsPort usando HTTP hacia la API Node.
type HTTPStatsClient struct {
	baseURL string
	client  *http.Client
}

// NewHTTPStatsClient crea un nuevo cliente de estadísticas.
func NewHTTPStatsClient(baseURL string) *HTTPStatsClient {
	return &HTTPStatsClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type statsRequest struct {
	Matrices [][][]float64 `json:"matrices"`
}

type statsResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    appqr.Stats `json:"data"`
}

// ComputeStats llama a POST {baseURL}/api/stats y devuelve las estadísticas.
func (c *HTTPStatsClient) ComputeStats(ctx context.Context, matrices [][][]float64, authHeader string) (*appqr.Stats, error) {
	bodyBytes, err := json.Marshal(statsRequest{Matrices: matrices})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/stats", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("stats service returned status %d", resp.StatusCode)
	}

	var sr statsResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return nil, err
	}

	if sr.Code != "000" {
		return nil, fmt.Errorf("stats service error: %s", sr.Message)
	}

	return &sr.Data, nil
}
