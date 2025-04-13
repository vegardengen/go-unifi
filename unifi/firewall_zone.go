package unifi

import (
	"context"
	"fmt"
)

type FirewallZone struct {
	ID string `json:"_id,omitempty"`

	// Hidden   bool   `json:"attr_hidden,omitempty"`
	// HiddenID string `json:"attr_hidden_id,omitempty"`
	// NoDelete bool   `json:"attr_no_delete,omitempty"`
	// NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Name        string   `json:"name"`
	DefaultZone bool     `json:default_zone,omitempty`
	NetworkIDs  []string `json:network_ids,omitempty`
	ZoneKey     string   `json:"zone_key,omitempty"`

	// Role string `json:"role"`
}

func (c *Client) ListFirewallZones(ctx context.Context, site string) ([]FirewallZone, error) {
	var respBody []FirewallZone


	err := c.do(ctx, "GET", fmt.Sprintf("site/%s/firewall/zone", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) GetFirewallZone(ctx context.Context, site, id string) (*FirewallZone, error) {
	firewallzones, err := c.ListFirewallZones(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, z := range firewallzones {
		if z.ID == id {
			return &z, nil
		}
	}

	return nil, &NotFoundError{}
}
