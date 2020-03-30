package base

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/tools"
)

// LinkDescriptorTLV defines Link Descriptor TLVs object
// https://tools.ietf.org/html/rfc7752#section-3.2.2
type LinkDescriptorTLV struct {
	Type   uint16
	Length uint16
	Value  []byte
}

func (tlv *LinkDescriptorTLV) String() string {
	var s string
	switch tlv.Type {
	case 258:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d (Link Local/Remote Identifiers)\n", tlv.Type)
		lri, err := UnmarshalLocalRemoteIdentifierTLV(tlv.Value)
		if err != nil {
			s += err.Error() + "\n"
			break
		}
		s += lri.String()
	case 259:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d (IPv4 interface address)\n", tlv.Type)
		s += fmt.Sprintf("      IPv4 interface address: %s\n", net.IP(tlv.Value).To4().String())
	case 260:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d (IPv4 neighbor address)\n", tlv.Type)
		s += fmt.Sprintf("      IPv4 neighbor address: %s\n", net.IP(tlv.Value).To4().String())
	case 261:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d (IPv6 interface address)\n", tlv.Type)
		s += fmt.Sprintf("      IPv6 interface address: %s\n", net.IP(tlv.Value).To16().String())
	case 262:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d (IPv6 neighbor address)\n", tlv.Type)
		s += fmt.Sprintf("      IPv6 neighbor address: %s\n", net.IP(tlv.Value).To16().String())
	case 263:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d (Multi-Topology Identifier)\n", tlv.Type)
		mit, err := UnmarshalMultiTopologyIdentifierTLV(tlv.Value)
		if err != nil {
			s += err.Error() + "\n"
			break
		}
		s += mit.String()
	default:
		s += fmt.Sprintf("   Link Descriptor TLV Type: %d\n", tlv.Type)
		s += fmt.Sprintf("   Link Descriptor TLV Length: %d\n", tlv.Length)
		s += "      Value: "
		s += tools.MessageHex(tlv.Value)
		s += "\n"
	}

	return s
}

// MarshalJSON defines a method to Marshal Link Descriptor TLV object into JSON format
func (tlv *LinkDescriptorTLV) MarshalJSON() ([]byte, error) {
	var jsonData []byte
	var b []byte

	jsonData = append(jsonData, '{')
	jsonData = append(jsonData, []byte("\"Type\":")...)
	jsonData = append(jsonData, []byte(fmt.Sprintf("%d,", tlv.Type))...)
	jsonData = append(jsonData, []byte("\"Description\":")...)
	switch tlv.Type {
	case 258:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"Local/Remote Identifiers\","))...)
		jsonData = append(jsonData, []byte("\"identifiersLocalRemote\":")...)
		lri, err := UnmarshalLocalRemoteIdentifierTLV(tlv.Value)
		if err != nil {
			return nil, err
		}
		b, err = json.Marshal(lri)
		if err != nil {
			return nil, err
		}
		jsonData = append(jsonData, b...)
	case 259:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"IPv4 interface address\","))...)
		jsonData = append(jsonData, []byte("\"ipv4InterfaceAddress\":")...)
		jsonData = append(jsonData, tools.RawBytesToJSON(tlv.Value)...)
	case 260:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"IPv4 neighbor address\","))...)
		jsonData = append(jsonData, []byte("\"ipv4NeighborAddress\":")...)
		jsonData = append(jsonData, tools.RawBytesToJSON(tlv.Value)...)
	case 261:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"IPv6 interface address\","))...)
		jsonData = append(jsonData, []byte("\"ipv6InterfaceAddress\":")...)
		jsonData = append(jsonData, tools.RawBytesToJSON(tlv.Value)...)
	case 262:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"IPv6 neighbor address\","))...)
		jsonData = append(jsonData, []byte("\"ipv6NeighborAddress\":")...)
		jsonData = append(jsonData, tools.RawBytesToJSON(tlv.Value)...)
	case 263:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"Multi-Topology Identifier\","))...)
		jsonData = append(jsonData, []byte("\"multiTopologyIdentifier\":")...)
		mit, err := UnmarshalMultiTopologyIdentifierTLV(tlv.Value)
		if err != nil {
			return nil, err
		}
		b, err = json.Marshal(mit)
		if err != nil {
			return nil, err
		}
		jsonData = append(jsonData, b...)
	default:
		jsonData = append(jsonData, []byte(fmt.Sprintf("\"Unknown Link TLV\","))...)
		jsonData = append(jsonData, []byte("\"Value\":")...)
		jsonData = append(jsonData, tools.RawBytesToJSON(tlv.Value)...)
	}
	jsonData = append(jsonData, '}')

	return jsonData, nil
}

// UnmarshalLinkDescriptorTLV builds Link Descriptor TLVs object
func UnmarshalLinkDescriptorTLV(b []byte) ([]LinkDescriptorTLV, error) {
	glog.V(6).Infof("LinkDescriptorTLV Raw: %s", tools.MessageHex(b))
	ltlvs := make([]LinkDescriptorTLV, 0)
	for p := 0; p < len(b); {
		ltlv := LinkDescriptorTLV{}
		ltlv.Type = binary.BigEndian.Uint16(b[p : p+2])
		p += 2
		ltlv.Length = binary.BigEndian.Uint16(b[p : p+2])
		p += 2
		ltlv.Value = make([]byte, ltlv.Length)
		copy(ltlv.Value, b[p:p+int(ltlv.Length)])
		ltlvs = append(ltlvs, ltlv)
		p += int(ltlv.Length)
	}

	return ltlvs, nil
}
