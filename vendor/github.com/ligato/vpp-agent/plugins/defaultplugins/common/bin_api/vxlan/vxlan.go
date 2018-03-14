// Code generated by govpp binapi-generator DO NOT EDIT.
// Package vxlan represents the VPP binary API of the 'vxlan' VPP module.
// Generated from '/usr/share/vpp/api/vxlan.api.json'
package vxlan

import "git.fd.io/govpp.git/api"

// VxlanAddDelTunnel represents the VPP binary API message 'vxlan_add_del_tunnel'.
//
type VxlanAddDelTunnel struct {
	IsAdd          uint8
	IsIpv6         uint8
	SrcAddress     []byte `struc:"[16]byte"`
	DstAddress     []byte `struc:"[16]byte"`
	McastSwIfIndex uint32
	EncapVrfID     uint32
	DecapNextIndex uint32
	Vni            uint32
}

func (*VxlanAddDelTunnel) GetMessageName() string {
	return "vxlan_add_del_tunnel"
}
func (*VxlanAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanAddDelTunnel) GetCrcString() string {
	return "79be0753"
}
func NewVxlanAddDelTunnel() api.Message {
	return &VxlanAddDelTunnel{}
}

// VxlanAddDelTunnelReply represents the VPP binary API message 'vxlan_add_del_tunnel_reply'.
//
type VxlanAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*VxlanAddDelTunnelReply) GetMessageName() string {
	return "vxlan_add_del_tunnel_reply"
}
func (*VxlanAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanAddDelTunnelReply) GetCrcString() string {
	return "3965e5df"
}
func NewVxlanAddDelTunnelReply() api.Message {
	return &VxlanAddDelTunnelReply{}
}

// VxlanTunnelDump represents the VPP binary API message 'vxlan_tunnel_dump'.
//
type VxlanTunnelDump struct {
	SwIfIndex uint32
}

func (*VxlanTunnelDump) GetMessageName() string {
	return "vxlan_tunnel_dump"
}
func (*VxlanTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*VxlanTunnelDump) GetCrcString() string {
	return "7d29e867"
}
func NewVxlanTunnelDump() api.Message {
	return &VxlanTunnelDump{}
}

// VxlanTunnelDetails represents the VPP binary API message 'vxlan_tunnel_details'.
//
type VxlanTunnelDetails struct {
	SwIfIndex      uint32
	SrcAddress     []byte `struc:"[16]byte"`
	DstAddress     []byte `struc:"[16]byte"`
	McastSwIfIndex uint32
	EncapVrfID     uint32
	DecapNextIndex uint32
	Vni            uint32
	IsIpv6         uint8
}

func (*VxlanTunnelDetails) GetMessageName() string {
	return "vxlan_tunnel_details"
}
func (*VxlanTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*VxlanTunnelDetails) GetCrcString() string {
	return "fa28d42c"
}
func NewVxlanTunnelDetails() api.Message {
	return &VxlanTunnelDetails{}
}

// SwInterfaceSetVxlanBypass represents the VPP binary API message 'sw_interface_set_vxlan_bypass'.
//
type SwInterfaceSetVxlanBypass struct {
	SwIfIndex uint32
	IsIpv6    uint8
	Enable    uint8
}

func (*SwInterfaceSetVxlanBypass) GetMessageName() string {
	return "sw_interface_set_vxlan_bypass"
}
func (*SwInterfaceSetVxlanBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceSetVxlanBypass) GetCrcString() string {
	return "da63ecfd"
}
func NewSwInterfaceSetVxlanBypass() api.Message {
	return &SwInterfaceSetVxlanBypass{}
}

// SwInterfaceSetVxlanBypassReply represents the VPP binary API message 'sw_interface_set_vxlan_bypass_reply'.
//
type SwInterfaceSetVxlanBypassReply struct {
	Retval int32
}

func (*SwInterfaceSetVxlanBypassReply) GetMessageName() string {
	return "sw_interface_set_vxlan_bypass_reply"
}
func (*SwInterfaceSetVxlanBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceSetVxlanBypassReply) GetCrcString() string {
	return "c4609ab5"
}
func NewSwInterfaceSetVxlanBypassReply() api.Message {
	return &SwInterfaceSetVxlanBypassReply{}
}
