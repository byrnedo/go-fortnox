package fortnox

import (
	"context"
	"fmt"
)

// Label data type
type Label struct {
	ID          int    `json:"Id"`
	Description string `json:"Description,omitempty"`
}

// ListLabelsResp Response for multiple labels
type ListLabelsResp struct {
	Labels []*Label `json:"Labels"`
}

// ListLabels lists labels
func (c *Client) ListLabels(ctx context.Context) ([]*Label, error) {
	resp := &ListLabelsResp{}

	err := c.request(ctx, "GET", "labels", nil, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.Labels, nil
}

// CreateLabelReq is the request used for creating labels
type CreateLabelReq struct {
	Label struct {
		Description string `json:"Description"`
	} `json:"Label"`
}

// LabelResp Response for single label
type LabelResp struct {
	Label Label `json:"Label"`
}

// CreateLabel creates a label
func (c *Client) CreateLabel(ctx context.Context, name string) (*Label, error) {

	resp := &LabelResp{}

	req := CreateLabelReq{}
	req.Label.Description = name
	err := c.request(ctx, "POST", "labels", &req, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Label, nil
}

// An UpdateLabelsReq is used in updating labels
type UpdateLabelsReq CreateLabelReq

// UpdateLabel updates a label
func (c *Client) UpdateLabel(ctx context.Context, id int, name string) (*Label, error) {

	resp := &LabelResp{}

	req := CreateLabelReq{}
	req.Label.Description = name
	err := c.request(ctx, "PUT", fmt.Sprintf("labels/%d", id), &req, nil, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Label, nil
}

// DeleteLabel deletes a label
func (c *Client) DeleteLabel(ctx context.Context, id int) error {
	return c.deleteResource(ctx, fmt.Sprintf("labels/%d", id))
}
