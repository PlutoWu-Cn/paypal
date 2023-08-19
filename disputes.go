package paypal

import (
	"context"
	"fmt"
	"net/http"
)

// ListDisputes Lists disputes with a summary set of details
// Endpoint: GET /v1/customer/disputes
func (c *Client) ListDisputes(ctx context.Context, params *ListDisputesParameters) (*ListDisputesResponse, error) {
	url := fmt.Sprintf("%s%s", c.APIBase, "/v1/customer/disputes")
	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		q := req.URL.Query()
		q.Add("start_time", params.StartTime)
		q.Add("disputed_transaction_id", params.DisputedTransactionId)
		q.Add("page_size", params.PageSize)
		q.Add("next_page_token", params.NextPageToken)
		q.Add("dispute_state", params.DisputeState)
		q.Add("update_time_before", params.UpdateTimeBefore)
		q.Add("update_time_after", params.UpdateTimeAfter)
		req.URL.RawQuery = q.Encode()
	}

	response := &ListDisputesResponse{}
	if err = c.SendWithAuth(req, response); err != nil {
		return nil, err
	}

	return response, nil
}
