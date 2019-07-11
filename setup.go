package dnslogging

import (
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyfile"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

// PluginName is the name of our plugin
const PluginName string = "dnslogging"

func init() {
	caddy.RegisterPlugin(PluginName, caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {

	// Normal Setup
	dnslogging, err := parseDNSLogging(c)
	if err != nil {
		return plugin.Error(PluginName, err)
	}
	err = dnslogging.Initialize()
	if err != nil {
		return err
	}

	// Pass xpf plugin to our context
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		dnslogging.Next = next
		return nil
	})

	// Setup startup and shutdown behaviour
	c.OnShutdown(func() error {
		dnslogging.Close()
		return nil
	})

	return nil
}

func parseDNSLogging(c *caddy.Controller) (*DNSLogging, error) {
	var (
		dl  *DNSLogging
		err error
		i   int
	)
	if err != nil {
		return dl, err
	}

	for c.Next() {
		if i > 0 {
			return nil, plugin.ErrOnce
		}
		i++
		dl, err = parseDNSLoggingStanza(&c.Dispenser)
		if err != nil {
			return nil, err
		}
	}

	return dl, nil
}

func parseDNSLoggingStanza(c *caddyfile.Dispenser) (dl *DNSLogging, err error) {
	dl, err = New()

	// xpf stanza if present
	for c.NextBlock() {
		if err := parseDNSLoggingBlock(c, dl); err != nil {
			return dl, err
		}
	}

	return dl, err
}

func parseDNSLoggingBlock(c *caddyfile.Dispenser, dl *DNSLogging) error {
	switch c.Val() {
	case "nats_url":
		if arg := c.NextArg(); !arg {
			return c.Errf("missing rr_type argument")
		}
		dl.natsURL = c.Val()
	default:
		return c.Errf("unknown property '%s'", c.Val())
	}
	return nil
}
