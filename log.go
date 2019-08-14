package dnslogging

import (
	"encoding/json"

	"github.com/miekg/dns"
)

type Log struct {
	Req  *dns.Msg
	Resp *dns.Msg
}

func (l *Log) pack() ([]byte, error) {
	buffer, err := json.Marshal(l)
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
