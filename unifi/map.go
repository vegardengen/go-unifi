// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
)

func (c *Client) ListMap(ctx context.Context, site string) ([]Map, error) {
	return c.listMap(ctx, site)
}

func (c *Client) GetMap(ctx context.Context, site, id string) (*Map, error) {
	return c.getMap(ctx, site, id)
}

func (c *Client) DeleteMap(ctx context.Context, site, id string) error {
	return c.deleteMap(ctx, site, id)
}

func (c *Client) CreateMap(ctx context.Context, site string, d *Map) (*Map, error) {
	return c.createMap(ctx, site, d)
}

func (c *Client) UpdateMap(ctx context.Context, site string, d *Map) (*Map, error) {
	return c.updateMap(ctx, site, d)
}
