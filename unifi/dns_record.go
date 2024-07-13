// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
)

func (c *Client) ListDNSRecord(ctx context.Context, site string) ([]DNSRecord, error) {
	return c.listDNSRecord(ctx, site)
}

func (c *Client) GetDNSRecord(ctx context.Context, site, id string) (*DNSRecord, error) {
	return c.getDNSRecord(ctx, site, id)
}

func (c *Client) DeleteDNSRecord(ctx context.Context, site, id string) error {
	return c.deleteDNSRecord(ctx, site, id)
}

func (c *Client) CreateDNSRecord(ctx context.Context, site string, d *DNSRecord) (*DNSRecord, error) {
	return c.createDNSRecord(ctx, site, d)
}

func (c *Client) UpdateDNSRecord(ctx context.Context, site string, d *DNSRecord) (*DNSRecord, error) {
	return c.updateDNSRecord(ctx, site, d)
}
