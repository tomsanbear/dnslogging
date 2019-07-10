module github.com/tomsanbear/dnslogging

go 1.12

require (
	github.com/avast/retry-go v2.3.0+incompatible
	github.com/coredns/coredns v1.5.0
	github.com/mholt/caddy v1.0.0
	github.com/miekg/dns v1.1.15
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/go-nats-streaming v0.4.5 // indirect
	github.com/nats-io/nats.go v1.8.1
	github.com/nats-io/stan.go v0.4.5 // indirect
	zombiezen.com/go/capnproto2 v2.17.0+incompatible
)

replace github.com/coredns/coredns => github.com/tomsanbear/coredns v1.5.1-0.20190710194947-7fe93df26a46
