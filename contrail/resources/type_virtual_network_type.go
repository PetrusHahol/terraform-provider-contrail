//
// Automatically generated. DO NOT EDIT.
//

package resources

type VirtualNetworkType struct {
	AllowTransit           bool   `json:"allow_transit,omitempty"`
	NetworkId              int    `json:"network_id,omitempty"`
	VxlanNetworkIdentifier int    `json:"vxlan_network_identifier,omitempty"`
	ForwardingMode         string `json:"forwarding_mode,omitempty"`
	Rpf                    string `json:"rpf,omitempty"`
	MirrorDestination      bool   `json:"mirror_destination,omitempty"`
}
