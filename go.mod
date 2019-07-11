module github.com/tomsanbear/dnslogging

go 1.12

require (
	github.com/avast/retry-go v2.3.0+incompatible
	github.com/caddyserver/caddy v1.0.1
	github.com/coredns/coredns v1.5.2
	github.com/miekg/dns v1.1.15
	github.com/nats-io/nats-server/v2 v2.0.0 // indirect
	github.com/nats-io/nats.go v1.8.1
	zombiezen.com/go/capnproto2 v2.17.0+incompatible
)

replace github.com/coredns/coredns => github.com/tomsanbear/coredns v1.5.1-0.20190704065049-f4a6bea4b414
