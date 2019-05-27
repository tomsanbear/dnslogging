package dnslogging

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/mholt/caddy"
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
	dnslogging, err := New()
	if err != nil {
		return plugin.Error(PluginName, err)
	}

	// Pass xpf plugin to our context
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		dnslogging.Next = next
		return dnslogging
	})

	// Setup startup and shutdown behaviour
	// c.OnStartup(func() error {
	// })
	// c.OnShutdown(func() error {
	// })

	return nil
}
