package dnslogging

import (
	"testing"

	"github.com/caddyserver/caddy"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		input      string
		shouldErr  bool
		expNatsURL string
	}{
		// positive
		{`dnslogging {
			nats_url nats://1.2.3.4:5222
		}`, false, "nats://1.2.3.4:5222"},
		{`dnslogging`, false, "nats://127.0.0.1:4222"},
		// negative
		{`dnslogging {
			nats_url
		}`, true, ""},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		dl, err := parseDNSLogging(c)
		if test.shouldErr {
			assert.Error(t, err, i)
		} else {
			assert.NoError(t, err, i)
			assert.Equal(t, test.expNatsURL, dl.natsURL, i)
		}
	}
}
