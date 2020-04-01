package bgpls

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/tools"
)

// NLRI defines BGP-LS NLRI object as collection of BGP-LS TLVs
// https://tools.ietf.org/html/rfc7752#section-3.3
type NLRI struct {
	LS []TLV
}

func (ls *NLRI) String() string {
	var s string

	s += "BGP-LS TLVs:" + "\n"
	for _, tlv := range ls.LS {
		s += tlv.String()
	}

	return s
}

// GetNodeFlags reeturns Flag Bits TLV carries a bit mask describing node attributes.
func (ls *NLRI) GetNodeFlags() uint8 {
	for _, tlv := range ls.LS {
		if tlv.Type != 1024 {
			continue
		}
		return uint8(tlv.Value[0])
	}
	return 0
}

// MarshalJSON defines a method to  BGP-LS TLV object into JSON format
func (ls *NLRI) MarshalJSON() ([]byte, error) {
	var jsonData []byte

	jsonData = append(jsonData, '{')
	jsonData = append(jsonData, []byte("\"BGPLSTLV\":")...)
	jsonData = append(jsonData, '[')
	if ls.LS != nil {
		for i, tlv := range ls.LS {
			b, err := json.Marshal(&tlv)
			if err != nil {
				return nil, err
			}
			jsonData = append(jsonData, b...)
			if i < len(ls.LS)-1 {
				jsonData = append(jsonData, ',')
			}
		}
	}
	jsonData = append(jsonData, ']')
	jsonData = append(jsonData, '}')

	return jsonData, nil
}

// UnmarshalBGPLSNLRI builds Prefix NLRI object
func UnmarshalBGPLSNLRI(b []byte) (*NLRI, error) {
	glog.V(6).Infof("BGPLSNLRI Raw: %s", tools.MessageHex(b))
	bgpls := NLRI{}
	ls, err := UnmarshalBGPLSTLV(b)
	if err != nil {
		return nil, err
	}
	bgpls.LS = ls

	return &bgpls, nil
}
