package unifi

import (
	"context"
	"fmt"
	"time"
)

type LastConnection struct {
	Mac      string `json:"mac,omitempty"`
	LastSeen int    `json:"last_seen,omitempty"`
}
type QosProfile struct{}
type PortTable struct {
	PortIdx                       int            `json:"port_idx,omitempty"`
	Media                         string         `json:"media,omitempty"`
	PortPoe                       bool           `json:"port_poe,omitempty"`
	PoeCaps                       int            `json:"poe_caps,omitempty"`
	SpeedCaps                     int            `json:"speed_caps,omitempty"`
	LastConnection                LastConnection `json:"last_connection,omitempty"`
	OpMode                        string         `json:"op_mode,omitempty"`
	Forward                       string         `json:"forward,omitempty"`
	PoeMode                       string         `json:"poe_mode,omitempty"`
	Anomalies                     int            `json:"anomalies,omitempty"`
	Autoneg                       bool           `json:"autoneg,omitempty"`
	Dot1XMode                     string         `json:"dot1x_mode,omitempty"`
	Dot1XStatus                   string         `json:"dot1x_status,omitempty"`
	Enable                        bool           `json:"enable,omitempty"`
	FlowctrlRx                    bool           `json:"flowctrl_rx,omitempty"`
	FlowctrlTx                    bool           `json:"flowctrl_tx,omitempty"`
	FullDuplex                    bool           `json:"full_duplex,omitempty"`
	IsUplink                      bool           `json:"is_uplink,omitempty"`
	Jumbo                         bool           `json:"jumbo,omitempty"`
	MacTableCount                 int            `json:"mac_table_count,omitempty"`
	PoeClass                      string         `json:"poe_class,omitempty"`
	PoeCurrent                    string         `json:"poe_current,omitempty"`
	PoeEnable                     bool           `json:"poe_enable,omitempty"`
	PoeGood                       bool           `json:"poe_good,omitempty"`
	PoePower                      string         `json:"poe_power,omitempty"`
	PoeVoltage                    string         `json:"poe_voltage,omitempty"`
	RxBroadcast                   int            `json:"rx_broadcast,omitempty"`
	RxBytes                       int64          `json:"rx_bytes,omitempty"`
	RxDropped                     int            `json:"rx_dropped,omitempty"`
	RxErrors                      int            `json:"rx_errors,omitempty"`
	RxMulticast                   int            `json:"rx_multicast,omitempty"`
	RxPackets                     int            `json:"rx_packets,omitempty"`
	Satisfaction                  int            `json:"satisfaction,omitempty"`
	SatisfactionReason            int            `json:"satisfaction_reason,omitempty"`
	Speed                         int            `json:"speed,omitempty"`
	StpPathcost                   int            `json:"stp_pathcost,omitempty"`
	StpState                      string         `json:"stp_state,omitempty"`
	TxBroadcast                   int            `json:"tx_broadcast,omitempty"`
	TxBytes                       int64          `json:"tx_bytes,omitempty"`
	TxDropped                     int            `json:"tx_dropped,omitempty"`
	TxErrors                      int            `json:"tx_errors,omitempty"`
	TxMulticast                   int            `json:"tx_multicast,omitempty"`
	TxPackets                     int64          `json:"tx_packets,omitempty"`
	Up                            bool           `json:"up,omitempty"`
	TxBytesR                      float64        `json:"tx_bytes-r,omitempty"`
	RxBytesR                      float64        `json:"rx_bytes-r,omitempty"`
	BytesR                        float64        `json:"bytes-r,omitempty"`
	SettingPreference             string         `json:"setting_preference,omitempty"`
	Name                          string         `json:"name,omitempty"`
	PortSecurityEnabled           bool           `json:"port_security_enabled,omitempty"`
	PortSecurityMacAddress        []any          `json:"port_security_mac_address,omitempty"`
	NativeNetworkconfID           string         `json:"native_networkconf_id,omitempty"`
	TaggedVlanMgmt                string         `json:"tagged_vlan_mgmt,omitempty"`
	MulticastRouterNetworkconfIds []any          `json:"multicast_router_networkconf_ids,omitempty"`
	LldpmedEnabled                bool           `json:"lldpmed_enabled,omitempty"`
	VoiceNetworkconfID            string         `json:"voice_networkconf_id,omitempty"`
	StormctrlBcastEnabled         bool           `json:"stormctrl_bcast_enabled,omitempty"`
	StormctrlBcastRate            int            `json:"stormctrl_bcast_rate,omitempty"`
	StormctrlMcastEnabled         bool           `json:"stormctrl_mcast_enabled,omitempty"`
	StormctrlMcastRate            int            `json:"stormctrl_mcast_rate,omitempty"`
	StormctrlUcastEnabled         bool           `json:"stormctrl_ucast_enabled,omitempty"`
	StormctrlUcastRate            int            `json:"stormctrl_ucast_rate,omitempty"`
	EgressRateLimitKbpsEnabled    bool           `json:"egress_rate_limit_kbps_enabled,omitempty"`
	Isolation                     bool           `json:"isolation,omitempty"`
	StpPortMode                   bool           `json:"stp_port_mode,omitempty"`
	PortKeepaliveEnabled          bool           `json:"port_keepalive_enabled,omitempty"`
	Masked                        bool           `json:"masked,omitempty"`
	AggregatedBy                  bool           `json:"aggregated_by,omitempty"`
	QosProfile                    QosProfile     `json:"qos_profile,omitempty"`
	SfpCompliance                 string         `json:"sfp_compliance,omitempty"`
	SfpFound                      bool           `json:"sfp_found,omitempty"`
	SfpPart                       string         `json:"sfp_part,omitempty"`
	SfpRev                        string         `json:"sfp_rev,omitempty"`
	SfpRxfault                    bool           `json:"sfp_rxfault,omitempty"`
	SfpSerial                     string         `json:"sfp_serial,omitempty"`
	SfpTxfault                    bool           `json:"sfp_txfault,omitempty"`
	SfpVendor                     string         `json:"sfp_vendor,omitempty"`
}
type ConfigNetwork struct {
	IP   string `json:"ip,omitempty"`
	Type string `json:"type,omitempty"`
}
type EthernetTable struct {
	NumPort   int    `json:"num_port,omitempty"`
	OtherMacs []any  `json:"other_macs,omitempty"`
	Name      string `json:"name,omitempty"`
	Mac       string `json:"mac,omitempty"`
}
type LastUplink struct {
	PortIdx          int    `json:"port_idx,omitempty"`
	UplinkMac        string `json:"uplink_mac,omitempty"`
	UplinkDeviceName string `json:"uplink_device_name,omitempty"`
	UplinkRemotePort int    `json:"uplink_remote_port,omitempty"`
	Type             string `json:"type,omitempty"`
}
type EtherLighting struct{}
type SwitchCaps struct {
	FeatureCaps          int `json:"feature_caps,omitempty"`
	MaxCustomMacAcls     int `json:"max_custom_mac_acls,omitempty"`
	MaxQosProfiles       int `json:"max_qos_profiles,omitempty"`
	MaxReservedRoutes    int `json:"max_reserved_routes,omitempty"`
	MaxMirrorSessions    int `json:"max_mirror_sessions,omitempty"`
	MaxAggregateSessions int `json:"max_aggregate_sessions,omitempty"`
	EtherlightCaps       int `json:"etherlight_caps,omitempty"`
	MaxCustomIPAcls      int `json:"max_custom_ip_acls,omitempty"`
	MaxGlobalAcls        int `json:"max_global_acls,omitempty"`
	VlanCaps             int `json:"vlan_caps,omitempty"`
	MaxVlanCount         int `json:"max_vlan_count,omitempty"`
	MaxStaticRoutes      int `json:"max_static_routes,omitempty"`
	MaxClassMaps         int `json:"max_class_maps,omitempty"`
	MaxL3Intf            int `json:"max_l3_intf,omitempty"`
}
type RpsOverride struct{}
type Uplink struct {
	UplinkMac        string  `json:"uplink_mac,omitempty"`
	UplinkDeviceName string  `json:"uplink_device_name,omitempty"`
	UplinkRemotePort int     `json:"uplink_remote_port,omitempty"`
	PortIdx          int     `json:"port_idx,omitempty"`
	Type             string  `json:"type,omitempty"`
	TxBytes          int64   `json:"tx_bytes,omitempty"`
	RxBytes          int64   `json:"rx_bytes,omitempty"`
	TxPackets        int     `json:"tx_packets,omitempty"`
	RxPackets        int64   `json:"rx_packets,omitempty"`
	FullDuplex       bool    `json:"full_duplex,omitempty"`
	IP               string  `json:"ip,omitempty"`
	Mac              string  `json:"mac,omitempty"`
	Name             string  `json:"name,omitempty"`
	Netmask          string  `json:"netmask,omitempty"`
	NumPort          int     `json:"num_port,omitempty"`
	RxDropped        int     `json:"rx_dropped,omitempty"`
	RxErrors         int     `json:"rx_errors,omitempty"`
	RxMulticast      int     `json:"rx_multicast,omitempty"`
	Speed            int     `json:"speed,omitempty"`
	TxDropped        int     `json:"tx_dropped,omitempty"`
	TxErrors         int     `json:"tx_errors,omitempty"`
	Up               bool    `json:"up,omitempty"`
	Media            string  `json:"media,omitempty"`
	MaxSpeed         int     `json:"max_speed,omitempty"`
	UplinkSource     string  `json:"uplink_source,omitempty"`
	TxBytesR         float64 `json:"tx_bytes-r,omitempty"`
	RxBytesR         float64 `json:"rx_bytes-r,omitempty"`
}
type SysStats struct {
	Loadavg1  string `json:"loadavg_1,omitempty"`
	Loadavg15 string `json:"loadavg_15,omitempty"`
	Loadavg5  string `json:"loadavg_5,omitempty"`
	MemBuffer int    `json:"mem_buffer,omitempty"`
	MemTotal  int    `json:"mem_total,omitempty"`
	MemUsed   int    `json:"mem_used,omitempty"`
}
type SystemStats struct {
	CPU    string `json:"cpu,omitempty"`
	Mem    string `json:"mem,omitempty"`
	Uptime string `json:"uptime,omitempty"`
}
type LldpTable struct {
	ChassisID     string `json:"chassis_id,omitempty"`
	IsWired       bool   `json:"is_wired,omitempty"`
	LocalPortIdx  int    `json:"local_port_idx,omitempty"`
	LocalPortName string `json:"local_port_name,omitempty"`
	PortID        string `json:"port_id,omitempty"`
}
type DownlinkTable struct {
	Mac        string `json:"mac,omitempty"`
	PortIdx    int    `json:"port_idx,omitempty"`
	Speed      int    `json:"speed,omitempty"`
	FullDuplex bool   `json:"full_duplex,omitempty"`
}
type Sw struct {
	SiteID            string    `json:"site_id,omitempty"`
	O                 string    `json:"o,omitempty"`
	Oid               string    `json:"oid,omitempty"`
	Sw                string    `json:"sw,omitempty"`
	Time              int64     `json:"time,omitempty"`
	Datetime          time.Time `json:"datetime,omitempty"`
	RxPackets         int       `json:"rx_packets,omitempty"`
	RxBytes           int64     `json:"rx_bytes,omitempty"`
	RxErrors          float64   `json:"rx_errors,omitempty"`
	RxDropped         float64   `json:"rx_dropped,omitempty"`
	RxCrypts          float64   `json:"rx_crypts,omitempty"`
	RxFrags           float64   `json:"rx_frags,omitempty"`
	TxPackets         int       `json:"tx_packets,omitempty"`
	TxBytes           int64     `json:"tx_bytes,omitempty"`
	TxErrors          float64   `json:"tx_errors,omitempty"`
	TxDropped         float64   `json:"tx_dropped,omitempty"`
	TxRetries         float64   `json:"tx_retries,omitempty"`
	RxMulticast       float64   `json:"rx_multicast,omitempty"`
	RxBroadcast       float64   `json:"rx_broadcast,omitempty"`
	TxMulticast       int       `json:"tx_multicast,omitempty"`
	TxBroadcast       float64   `json:"tx_broadcast,omitempty"`
	Bytes             int64     `json:"bytes,omitempty"`
	Duration          int       `json:"duration,omitempty"`
	Port1RxPackets    int       `json:"port_1-rx_packets,omitempty"`
	Port1RxBytes      int64     `json:"port_1-rx_bytes,omitempty"`
	Port1TxPackets    int       `json:"port_1-tx_packets,omitempty"`
	Port1TxBytes      int64     `json:"port_1-tx_bytes,omitempty"`
	Port1RxMulticast  float64   `json:"port_1-rx_multicast,omitempty"`
	Port1RxBroadcast  float64   `json:"port_1-rx_broadcast,omitempty"`
	Port1TxMulticast  float64   `json:"port_1-tx_multicast,omitempty"`
	Port1TxBroadcast  float64   `json:"port_1-tx_broadcast,omitempty"`
	Port3TxPackets    float64   `json:"port_3-tx_packets,omitempty"`
	Port3TxBytes      int       `json:"port_3-tx_bytes,omitempty"`
	Port3TxMulticast  float64   `json:"port_3-tx_multicast,omitempty"`
	Port3TxBroadcast  float64   `json:"port_3-tx_broadcast,omitempty"`
	Port5RxPackets    float64   `json:"port_5-rx_packets,omitempty"`
	Port5RxBytes      int       `json:"port_5-rx_bytes,omitempty"`
	Port5TxPackets    float64   `json:"port_5-tx_packets,omitempty"`
	Port5TxBytes      int       `json:"port_5-tx_bytes,omitempty"`
	Port5RxBroadcast  float64   `json:"port_5-rx_broadcast,omitempty"`
	Port5TxMulticast  float64   `json:"port_5-tx_multicast,omitempty"`
	Port5TxBroadcast  float64   `json:"port_5-tx_broadcast,omitempty"`
	Port6TxPackets    float64   `json:"port_6-tx_packets,omitempty"`
	Port6TxBytes      int       `json:"port_6-tx_bytes,omitempty"`
	Port6TxMulticast  float64   `json:"port_6-tx_multicast,omitempty"`
	Port6TxBroadcast  float64   `json:"port_6-tx_broadcast,omitempty"`
	Port19RxPackets   int       `json:"port_19-rx_packets,omitempty"`
	Port19RxBytes     int64     `json:"port_19-rx_bytes,omitempty"`
	Port19TxPackets   int       `json:"port_19-tx_packets,omitempty"`
	Port19TxBytes     int64     `json:"port_19-tx_bytes,omitempty"`
	Port19RxMulticast float64   `json:"port_19-rx_multicast,omitempty"`
	Port19RxBroadcast float64   `json:"port_19-rx_broadcast,omitempty"`
	Port19TxMulticast float64   `json:"port_19-tx_multicast,omitempty"`
	Port19TxBroadcast float64   `json:"port_19-tx_broadcast,omitempty"`
	Port21RxPackets   int       `json:"port_21-rx_packets,omitempty"`
	Port21RxBytes     int64     `json:"port_21-rx_bytes,omitempty"`
	Port21TxPackets   int       `json:"port_21-tx_packets,omitempty"`
	Port21TxBytes     int64     `json:"port_21-tx_bytes,omitempty"`
	Port21RxMulticast float64   `json:"port_21-rx_multicast,omitempty"`
	Port21RxBroadcast float64   `json:"port_21-rx_broadcast,omitempty"`
	Port21TxMulticast float64   `json:"port_21-tx_multicast,omitempty"`
	Port21TxBroadcast float64   `json:"port_21-tx_broadcast,omitempty"`
	Port22RxPackets   float64   `json:"port_22-rx_packets,omitempty"`
	Port22RxBytes     int       `json:"port_22-rx_bytes,omitempty"`
	Port22TxPackets   float64   `json:"port_22-tx_packets,omitempty"`
	Port22TxBytes     int       `json:"port_22-tx_bytes,omitempty"`
	Port22TxMulticast float64   `json:"port_22-tx_multicast,omitempty"`
	Port22TxBroadcast float64   `json:"port_22-tx_broadcast,omitempty"`
	Port23RxPackets   int       `json:"port_23-rx_packets,omitempty"`
	Port23RxBytes     int64     `json:"port_23-rx_bytes,omitempty"`
	Port23TxPackets   int       `json:"port_23-tx_packets,omitempty"`
	Port23TxBytes     int64     `json:"port_23-tx_bytes,omitempty"`
	Port23RxMulticast float64   `json:"port_23-rx_multicast,omitempty"`
	Port23RxBroadcast float64   `json:"port_23-rx_broadcast,omitempty"`
	Port23TxMulticast float64   `json:"port_23-tx_multicast,omitempty"`
	Port23TxBroadcast float64   `json:"port_23-tx_broadcast,omitempty"`
	Port25RxPackets   int       `json:"port_25-rx_packets,omitempty"`
	Port25RxBytes     int64     `json:"port_25-rx_bytes,omitempty"`
	Port25TxPackets   int       `json:"port_25-tx_packets,omitempty"`
	Port25TxBytes     int64     `json:"port_25-tx_bytes,omitempty"`
	Port25RxMulticast float64   `json:"port_25-rx_multicast,omitempty"`
	Port25RxBroadcast float64   `json:"port_25-rx_broadcast,omitempty"`
	Port25TxMulticast float64   `json:"port_25-tx_multicast,omitempty"`
	Port25TxBroadcast float64   `json:"port_25-tx_broadcast,omitempty"`
	Port26TxPackets   float64   `json:"port_26-tx_packets,omitempty"`
	Port26TxBytes     int       `json:"port_26-tx_bytes,omitempty"`
	Port26TxMulticast float64   `json:"port_26-tx_multicast,omitempty"`
	Port26TxBroadcast float64   `json:"port_26-tx_broadcast,omitempty"`
	Port22RxMulticast float64   `json:"port_22-rx_multicast,omitempty"`
	Port5RxMulticast  float64   `json:"port_5-rx_multicast,omitempty"`
	Port23TxDropped   float64   `json:"port_23-tx_dropped,omitempty"`
	Port22RxBroadcast float64   `json:"port_22-rx_broadcast,omitempty"`
	Port3RxPackets    float64   `json:"port_3-rx_packets,omitempty"`
	Port3RxBytes      float64   `json:"port_3-rx_bytes,omitempty"`
	Port3RxBroadcast  float64   `json:"port_3-rx_broadcast,omitempty"`
	Port1TxDropped    float64   `json:"port_1-tx_dropped,omitempty"`
	Port21TxDropped   float64   `json:"port_21-tx_dropped,omitempty"`
	Port3RxMulticast  float64   `json:"port_3-rx_multicast,omitempty"`
	Port19TxDropped   float64   `json:"port_19-tx_dropped,omitempty"`
	Port20TxPackets   float64   `json:"port_20-tx_packets,omitempty"`
	Port20TxBytes     float64   `json:"port_20-tx_bytes,omitempty"`
	Port20TxMulticast float64   `json:"port_20-tx_multicast,omitempty"`
	Port20TxBroadcast float64   `json:"port_20-tx_broadcast,omitempty"`
	Port20RxPackets   float64   `json:"port_20-rx_packets,omitempty"`
	Port20RxBytes     float64   `json:"port_20-rx_bytes,omitempty"`
	Port20RxBroadcast float64   `json:"port_20-rx_broadcast,omitempty"`
}
type Stat struct {
	Sw Sw `json:"sw,omitempty"`
}
type RpsPortTable struct {
	PortIdx  int    `json:"port_idx,omitempty"`
	Name     string `json:"name,omitempty"`
	PortMode string `json:"port_mode,omitempty"`
}
type Rps struct {
	PowerManagementMode string         `json:"power_management_mode,omitempty"`
	RpsPortTable        []RpsPortTable `json:"rps_port_table,omitempty"`
}
type DeviceV2 struct {
	*Device
	PortTable                           []PortTable     `json:"port_table,omitempty"`
	RequiredVersion                     string          `json:"required_version,omitempty"`
	LicenseState                        string          `json:"license_state,omitempty"`
	BoardRev                            int             `json:"board_rev,omitempty"`
	SetupID                             string          `json:"setup_id,omitempty"`
	HwCaps                              int             `json:"hw_caps,omitempty"`
	RebootDuration                      int             `json:"reboot_duration,omitempty"`
	SlimcfgCaps                         int             `json:"slimcfg_caps,omitempty"`
	ManufacturerID                      int             `json:"manufacturer_id,omitempty"`
	Sysid                               int             `json:"sysid,omitempty"`
	LcmOrientationOverride              int             `json:"lcm_orientation_override,omitempty"`
	IP                                  string          `json:"ip,omitempty"`
	Fw2Caps                             int             `json:"fw2_caps,omitempty"`
	Version                             string          `json:"version,omitempty"`
	AdoptionCompleted                   bool            `json:"adoption_completed,omitempty"`
	UnsupportedReason                   int             `json:"unsupported_reason,omitempty"`
	Shortname                           string          `json:"shortname,omitempty"`
	AnonID                              string          `json:"anon_id,omitempty"`
	LastConnectionNetworkID             string          `json:"last_connection_network_id,omitempty"`
	AdoptedAt                           int             `json:"adopted_at,omitempty"`
	FwCaps                              int             `json:"fw_caps,omitempty"`
	ExternalID                          string          `json:"external_id,omitempty"`
	TwoPhaseAdopt                       bool            `json:"two_phase_adopt,omitempty"`
	ConnectedAt                         int             `json:"connected_at,omitempty"`
	InformIP                            string          `json:"inform_ip,omitempty"`
	Cfgversion                          string          `json:"cfgversion,omitempty"`
	ProvisionedAt                       int             `json:"provisioned_at,omitempty"`
	InformURL                           string          `json:"inform_url,omitempty"`
	UpgradeDuration                     int             `json:"upgrade_duration,omitempty"`
	EthernetTable                       []EthernetTable `json:"ethernet_table,omitempty"`
	ServiceMac                          string          `json:"service_mac,omitempty"`
	BleCaps                             int             `json:"ble_caps,omitempty"`
	SysErrorCaps                        int             `json:"sys_error_caps,omitempty"`
	Ipv6                                []string        `json:"ipv6,omitempty"`
	LastUplink                          LastUplink      `json:"last_uplink,omitempty"`
	DisconnectedAt                      int             `json:"disconnected_at,omitempty"`
	Architecture                        string          `json:"architecture,omitempty"`
	XAesGcm                             bool            `json:"x_aes_gcm,omitempty"`
	HasFan                              bool            `json:"has_fan,omitempty"`
	ModelIncompatible                   bool            `json:"model_incompatible,omitempty"`
	XAuthkey                            string          `json:"x_authkey,omitempty"`
	XSSHHostkeyFingerprint              string          `json:"x_ssh_hostkey_fingerprint,omitempty"`
	Satisfaction                        int             `json:"satisfaction,omitempty"`
	ModelInEol                          bool            `json:"model_in_eol,omitempty"`
	Anomalies                           int             `json:"anomalies,omitempty"`
	HasTemperature                      bool            `json:"has_temperature,omitempty"`
	SwitchCaps                          SwitchCaps      `json:"switch_caps,omitempty"`
	ModelInLts                          bool            `json:"model_in_lts,omitempty"`
	KernelVersion                       string          `json:"kernel_version,omitempty"`
	Serial                              string          `json:"serial,omitempty"`
	CredentialCaps                      int             `json:"credential_caps,omitempty"`
	Default                             bool            `json:"default,omitempty"`
	DiscoveredVia                       string          `json:"discovered_via,omitempty"`
	AdoptIP                             string          `json:"adopt_ip,omitempty"`
	AdoptURL                            string          `json:"adopt_url,omitempty"`
	LastSeen                            int             `json:"last_seen,omitempty"`
	MinInformIntervalSeconds            int             `json:"min_inform_interval_seconds,omitempty"`
	Upgradable                          bool            `json:"upgradable,omitempty"`
	AdoptableWhenUpgraded               bool            `json:"adoptable_when_upgraded,omitempty"`
	Rollupgrade                         bool            `json:"rollupgrade,omitempty"`
	KnownCfgversion                     string          `json:"known_cfgversion,omitempty"`
	Uptime                              int             `json:"uptime,omitempty"`
	Uptime0                             int             `json:"_uptime,omitempty"`
	Locating                            bool            `json:"locating,omitempty"`
	StartConnectedMillis                int64           `json:"start_connected_millis,omitempty"`
	PrevNonBusyState                    int             `json:"prev_non_busy_state,omitempty"`
	NextInterval                        int             `json:"next_interval,omitempty"`
	SysStats                            SysStats        `json:"sys_stats,omitempty"`
	SystemStats                         SystemStats     `json:"system-stats,omitempty"`
	SSHSessionTable                     []any           `json:"ssh_session_table,omitempty"`
	LldpTable                           []LldpTable     `json:"lldp_table,omitempty"`
	DisplayableVersion                  string          `json:"displayable_version,omitempty"`
	ConnectionNetworkID                 string          `json:"connection_network_id,omitempty"`
	ConnectionNetworkName               string          `json:"connection_network_name,omitempty"`
	StartupTimestamp                    int             `json:"startup_timestamp,omitempty"`
	IsAccessPoint                       bool            `json:"is_access_point,omitempty"`
	SafeForAutoupgrade                  bool            `json:"safe_for_autoupgrade,omitempty"`
	FanLevel                            int             `json:"fan_level,omitempty"`
	GeneralTemperature                  int             `json:"general_temperature,omitempty"`
	Overheating                         bool            `json:"overheating,omitempty"`
	TotalMaxPower                       int             `json:"total_max_power,omitempty"`
	DownlinkTable                       []DownlinkTable `json:"downlink_table,omitempty"`
	UplinkDepth                         int             `json:"uplink_depth,omitempty"`
	DownlinkLldpMacs                    []any           `json:"downlink_lldp_macs,omitempty"`
	DhcpServerTable                     []any           `json:"dhcp_server_table,omitempty"`
	Watched                             bool            `json:"watched,omitempty"`
	ConnectRequestIP                    string          `json:"connect_request_ip,omitempty"`
	ConnectRequestPort                  string          `json:"connect_request_port,omitempty"`
	WatchedTimeout                      int             `json:"watched_timeout,omitempty"`
	TotalUsedPower                      float64         `json:"total_used_power,omitempty"`
	DetailedStates                      DetailedStates  `json:"detailed_states,omitempty"`
	Ipv4LeaseExpirationTimestampSeconds int             `json:"ipv4_lease_expiration_timestamp_seconds,omitempty"`
	Stat                                Stat            `json:"stat,omitempty"`
	TxBytes                             int64           `json:"tx_bytes,omitempty"`
	RxBytes                             int64           `json:"rx_bytes,omitempty"`
	Bytes                               int64           `json:"bytes,omitempty"`
	NumSta                              int             `json:"num_sta,omitempty"`
	UserNumSta                          int             `json:"user-num_sta,omitempty"`
	GuestNumSta                         int             `json:"guest-num_sta,omitempty"`
	Rps                                 Rps             `json:"rps,omitempty"`
	XHasSSHHostkey                      bool            `json:"x_has_ssh_hostkey,omitempty"`
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
