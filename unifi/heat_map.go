// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
)

func (c *Client) ListHeatMap(ctx context.Context, site string) ([]HeatMap, error) {
	return c.listHeatMap(ctx, site)
}

func (c *Client) GetHeatMap(ctx context.Context, site, id string) (*HeatMap, error) {
	return c.getHeatMap(ctx, site, id)
}

func (c *Client) DeleteHeatMap(ctx context.Context, site, id string) error {
	return c.deleteHeatMap(ctx, site, id)
}

func (c *Client) CreateHeatMap(ctx context.Context, site string, d *HeatMap) (*HeatMap, error) {
	return c.createHeatMap(ctx, site, d)
}

func (c *Client) UpdateHeatMap(ctx context.Context, site string, d *HeatMap) (*HeatMap, error) {
	return c.updateHeatMap(ctx, site, d)
}
