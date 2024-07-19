// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"fmt"
)

func (c *Client) ListDNSRecord(ctx context.Context, site string) ([]DNSRecord, error) {
	var respBody []DNSRecord

	err := c.do(ctx, "GET", fmt.Sprintf("%s/site/%s/static-dns", c.apiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) GetDNSRecord(ctx context.Context, site, id string) (*DNSRecord, error) {
	var respBody DNSRecord

	err := c.do(ctx, "GET", fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiV2Path, site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	// if len(respBody.Data) != 1 {
	// 	return nil, &NotFoundError{}
	// }

	// d := respBody.Data[0]
	return &respBody, nil
}

func (c *Client) DeleteDNSRecord(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateDNSRecord(ctx context.Context, site string, d *DNSRecord) (*DNSRecord, error) {
	var respBody DNSRecord
	err := c.do(ctx, "POST", fmt.Sprintf("%s/site/%s/static-dns", c.apiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

func (c *Client) UpdateDNSRecord(ctx context.Context, site string, d *DNSRecord) (*DNSRecord, error) {
	var respBody DNSRecord

	err := c.do(ctx, "PUT", fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	// if len(respBody) != nil {
	// 	return nil, &NotFoundError{}
	// }

	return &respBody, nil
}
