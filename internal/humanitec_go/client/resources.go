package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sendgrid/rest"

	humanitec "github.com/score-spec/score-humanitec/internal/humanitec_go/types"
)

// ListResourceTypes lists all resource types available to the organization.
func (api *apiClient) ListResourceTypes(ctx context.Context, orgID string) ([]humanitec.ResourceType, error) {
	apiPath := fmt.Sprintf("/orgs/%s/resources/types", orgID)
	req := rest.Request{
		Method:  http.MethodGet,
		BaseURL: api.baseUrl + apiPath,
		Headers: map[string]string{
			"Authorization": "Bearer " + api.token,
			"Accept":        "application/json",
		},
	}

	resp, err := api.client.SendWithContext(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("humanitec api: %s %s: %w", req.Method, req.BaseURL, err)
	}

	switch resp.StatusCode {

	case http.StatusOK:
		{
			var res []humanitec.ResourceType
			if err = json.Unmarshal([]byte(resp.Body), &res); err != nil {
				return nil, fmt.Errorf("humanitec api: %s %s: parsing response: %w", req.Method, req.BaseURL, err)
			}
			return res, nil
		}

	default:
		return nil, fmt.Errorf("humanitec api: %s %s: HTTP %d - %s", req.Method, req.BaseURL, resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}
