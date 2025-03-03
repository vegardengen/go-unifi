package unifi

import (
	"context"
	"fmt"
)

type LastConnection struct {
	Mac      string `json:"mac,omitempty"`
	LastSeen int    `json:"last_seen,omitempty"`
}
type PortTable struct {
	PortIdx        int            `json:"port_idx,omitempty"`
	Media          string         `json:"media,omitempty"`
	PortPoe        bool           `json:"port_poe,omitempty"`
	PoeCaps        int            `json:"poe_caps,omitempty"`
	SpeedCaps      int            `json:"speed_caps,omitempty"`
	LastConnection LastConnection `json:"last_connection,omitempty"`
	OpMode         string         `json:"op_mode,omitempty"`
	Forward        string         `json:"forward,omitempty"`
	PoeMode        string         `json:"poe_mode,omitempty"`
	FullDuplex     bool           `json:"full_duplex,omitempty"`
	IsUplink       bool           `json:"is_uplink,omitempty"`
	MacTableCount  int            `json:"mac_table_count,omitempty"`
	PoeClass       string         `json:"poe_class,omitempty"`
	PoeCurrent     string         `json:"poe_current,omitempty"`
	PoeEnable      bool           `json:"poe_enable,omitempty"`
	PoeGood        bool           `json:"poe_good,omitempty"`
	PoePower       string         `json:"poe_power,omitempty"`
	PoeVoltage     string         `json:"poe_voltage,omitempty"`
}
type DeviceV2 struct {
	PortTable []PortTable `json:"port_table,omitempty"`
}

func (c *Client) GetDeviceByMACv2(ctx context.Context, site, mac string) (*DeviceV2, error) {
	var respBody struct {
		Meta meta       `json:"meta"`
		Data []DeviceV2 `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/stat/device/%s", site, mac), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}
