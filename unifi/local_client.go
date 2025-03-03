package unifi

import (
	"context"
	"fmt"
)

type LocalClient struct {
	Blocked                       bool        `json:"blocked,omitempty"`
	DisplayName                   string      `json:"display_name,omitempty"`
	Fingerprint                   Fingerprint `json:"fingerprint,omitempty"`
	FirstSeen                     int         `json:"first_seen,omitempty"`
	FixedIP                       string      `json:"fixed_ip,omitempty"`
	Hostname                      string      `json:"hostname,omitempty"`
	ID                            string      `json:"id,omitempty"`
	IsAllowedInVisualProgramming  bool        `json:"is_allowed_in_visual_programming,omitempty"`
	IsGuest                       bool        `json:"is_guest,omitempty"`
	IsWired                       bool        `json:"is_wired,omitempty"`
	LastConnectionNetworkID       string      `json:"last_connection_network_id,omitempty"`
	LastConnectionNetworkName     string      `json:"last_connection_network_name,omitempty"`
	LastIP                        string      `json:"last_ip,omitempty"`
	LastIpv6                      []string    `json:"last_ipv6,omitempty"`
	LastSeen                      int         `json:"last_seen,omitempty"`
	LastUplinkMac                 string      `json:"last_uplink_mac,omitempty"`
	LastUplinkName                string      `json:"last_uplink_name,omitempty"`
	LocalDNSRecord                string      `json:"local_dns_record,omitempty"`
	LocalDNSRecordEnabled         bool        `json:"local_dns_record_enabled,omitempty"`
	Mac                           string      `json:"mac,omitempty"`
	ModelName                     string      `json:"model_name,omitempty"`
	Name                          string      `json:"name,omitempty"`
	Noted                         bool        `json:"noted,omitempty"`
	Oui                           string      `json:"oui,omitempty"`
	SiteID                        string      `json:"site_id,omitempty"`
	Status                        string      `json:"status,omitempty"`
	SwPort                        int         `json:"sw_port,omitempty"`
	Tags                          []any       `json:"tags,omitempty"`
	Type                          string      `json:"type,omitempty"`
	UnifiDevice                   bool        `json:"unifi_device,omitempty"`
	UplinkMac                     string      `json:"uplink_mac,omitempty"`
	UseFixedip                    bool        `json:"use_fixedip,omitempty"`
	UserID                        string      `json:"user_id,omitempty"`
	UsergroupID                   string      `json:"usergroup_id,omitempty"`
	VirtualNetworkOverrideEnabled bool        `json:"virtual_network_override_enabled,omitempty"`
	VirtualNetworkOverrideID      string      `json:"virtual_network_override_id,omitempty"`
	WiredRateMbps                 int         `json:"wired_rate_mbps,omitempty"`
}

func (c *Client) GetLocalClient(ctx context.Context, site string, mac string) (*LocalClient, error) {
	var respBody LocalClient

	err := c.do(ctx, "GET", fmt.Sprintf("%s/site/%s/clients/local/%s?includeUnifiDevices=true&includeUnifiDevices=true", c.apiV2Path, site, mac), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
