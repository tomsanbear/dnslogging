package dnslogging

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/miekg/dns"
)

func TestLogMarshal(t *testing.T) {
	responseTestMsg := new(dns.Msg)
	responseTestMsg.SetQuestion("google.com.", dns.TypeA)

	testLog := &Log{
		Req:  responseTestMsg,
		Resp: responseTestMsg,
	}
	_, err := testLog.pack()
	assert.NoError(t, err)
}
