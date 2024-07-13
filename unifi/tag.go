// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
)

func (c *Client) ListTag(ctx context.Context, site string) ([]Tag, error) {
	return c.listTag(ctx, site)
}

func (c *Client) GetTag(ctx context.Context, site, id string) (*Tag, error) {
	return c.getTag(ctx, site, id)
}

func (c *Client) DeleteTag(ctx context.Context, site, id string) error {
	return c.deleteTag(ctx, site, id)
}

func (c *Client) CreateTag(ctx context.Context, site string, d *Tag) (*Tag, error) {
	return c.createTag(ctx, site, d)
}

func (c *Client) UpdateTag(ctx context.Context, site string, d *Tag) (*Tag, error) {
	return c.updateTag(ctx, site, d)
}
