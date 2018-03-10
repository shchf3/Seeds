package model

//go:generate reform

// SsNode represents a row in ss_node table.
//reform:ss_node
type SsNode struct {
	ID                     int32   `reform:"id,pk"`
	Name                   string  `reform:"name"`
	Type                   int32   `reform:"type"`
	Server                 string  `reform:"server"`
	Method                 string  `reform:"method"`
	Info                   string  `reform:"info"`
	Status                 string  `reform:"status"`
	Sort                   int32   `reform:"sort"`
	CustomMethod           int8    `reform:"custom_method"`
	TrafficRate            float32 `reform:"traffic_rate"`
	NodeClass              int32   `reform:"node_class"`
	NodeSpeedlimit         string  `reform:"node_speedlimit"`
	NodeConnector          int32   `reform:"node_connector"`
	NodeBandwidth          int64   `reform:"node_bandwidth"`
	NodeBandwidthLimit     int64   `reform:"node_bandwidth_limit"`
	BandwidthlimitResetday int32   `reform:"bandwidthlimit_resetday"`
	NodeHeartbeat          int64   `reform:"node_heartbeat"`
	NodeIp                 *string `reform:"node_ip"`
	NodeGroup              int32   `reform:"node_group"`
	CustomRss              int32   `reform:"custom_rss"`
	MuOnly                 *int32  `reform:"mu_only"`
}
