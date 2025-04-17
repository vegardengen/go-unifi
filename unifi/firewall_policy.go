package unifi

import (
	"context"
	"fmt"
)

type FirewallPolicy struct {
	ID string `json:"_id,omitempty"`

	// Hidden   bool   `json:"attr_hidden,omitempty"`
	// HiddenID string `json:"attr_hidden_id,omitempty"`
	// NoDelete bool   `json:"attr_no_delete,omitempty"`
	// NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Action              string   `json:"action"`
	ConnectionStateType string   `json:"connection_state_type"`
	ConnectionStates    []string `json:"connection_states"`
	CreateAllowRespond  bool     `json:"create_allow_respond"`
	Description         string   `json:"description"`
	Destination         struct {
		IPGroupID          string   `json:"ip_group_id"`
		Ips                []string `json:"ips"`
		MatchOppositeIps   bool     `json:"match_opposite_ips"`
		MatchOppositePorts bool     `json:"match_opposite_ports"`
		MatchingTarget     string   `json:"matching_target"`
		NetworkIDs         [] string   `json:"network_ids"`
		Port               string   `json:"port"`
		PortGroupID        string   `json:"port_group_id"`
		PortMatchingType   string   `json:"port_matching_type"`
		Regions            []string `json:"regions"`
		ZoneID             string   `json:"zone_id"`
	} `json:"destination"`
	Enabled               bool   `json:"enabled"`
	IcmpTypename          string `json:"icmp_typename"`
	IcmpV6Typename        string `json:"icmp_v6_typename"`
	Index                 int64  `json:"index"`
	IPVersion             string `json:"ip_version"`
	Logging               bool   `json:"logging"`
	MatchIPSec            bool   `json:"match_ip_sec"`
	MatchIPSecType        string `json:"match_ip_sec_type"`
	MatchOppositeProtocol bool   `json:"match_opposite_protocol"`
	Name                  string `json:"name"`
	OriginID              string `json:"origin_id"`
	OriginType            string `json:"origin_type"`
	Predefined            bool   `json:"predefined"`
	Protocol              string `json:"protocol"`
	Schedule              struct {
		Mode         string        `json:"mode"`
		DateStart     string  `json:"date_start"`
		DateEnd     string  `json:"date_end"`
		RepeatOnDays []string `json:"repeat_on_days"`
		TimeAllDay   bool          `json:"time_all_day"`
		TimeRangeStart string      `json:"time_range_start"`
		TimeRangeEnd string      `json:"time_range_end"`
	} `json:"schedule"`
	Source struct {
		ClientMacs            []string `json:"client_macs"`
		Ips                   []string `json:"ips"`
		MatchMac              bool     `json:"match_mac"`
		MatchOppositeIps      bool     `json:"match_opposite_ips"`
		MatchOppositeNetworks bool     `json:"match_opposite_networks"`
		MatchOppositePorts    bool     `json:"match_opposite_ports"`
		MatchingTarget        string   `json:"matching_target"`
		MatchingTargetType    string   `json:"matching_target_type"`
		NetworkIDs            []string `json:"network_ids"`
		Port                  string   `json:"port"`
		PortMatchingType      string   `json:"port_matching_type"`
		ZoneID                string   `json:"zone_id"`
	} `json:"source"`

	// Role string `json:"role"`
}

func (c *Client) ListFirewallPolicy(ctx context.Context, site string) ([]FirewallPolicy, error) {
	var respBody []FirewallPolicy

	err := c.do_versioned(ctx, "V2", "GET", fmt.Sprintf("site/%s/firewall-policies", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) GetFirewallPolicy(ctx context.Context, site, id string) (*FirewallPolicy, error) {

	var respBody FirewallPolicy
	err := c.do_versioned(ctx, "V2", "GET", fmt.Sprintf("site/%s/firewall-policies/%s", site,id), nil, &respBody)
        if err != nil {
	   return nil, &NotFoundError{}
        }

	new := respBody
        return &new, nil
}

func (c *Client) DeleteFirewallPolicy(ctx context.Context, site, id string) error {
	var respBody FirewallPolicy
        err := c.do_versioned(ctx, "V2", "DELETE", fmt.Sprintf("site/%s/firewall-policies/%s", site, id), nil, &respBody)

        if err != nil {
                return err
        }
        return nil
}

func (c *Client) CreaateFirewallPolicy(ctx context.Context, site string, d *FirewallPolicy) (*FirewallPolicy, error) {
        var respBody FirewallPolicy

        err := c.do_versioned(ctx, "V2", "POST", fmt.Sprintf("site/%s/firewall-policies", site), d, &respBody)
        if err != nil {
                return nil, err
        }


        new := respBody

        return &new, nil
}

func (c *Client) UpdateFirewallPolicy(ctx context.Context, site string, d *FirewallPolicy) (*FirewallPolicy, error) {
        var respBody FirewallPolicy

	err := c.do_versioned(ctx, "V2", "PUT", fmt.Sprintf("site/%s/firewall-policies/%s", site, d.ID), d, &respBody)
        if err != nil {
                return nil, err
        }

        new := respBody

        return &new, nil
}

