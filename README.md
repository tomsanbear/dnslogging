[![Build Status](https://travis-ci.org/tomsanbear/dnslogging.svg?branch=master)](https://travis-ci.org/tomsanbear/dnslogging) [![codecov](https://codecov.io/gh/tomsanbear/dnslogging/branch/master/graph/badge.svg)](https://codecov.io/gh/tomsanbear/dnslogging) [![GitHub version](https://badge.fury.io/gh/tomsanbear%2Fdnslogging.svg)](https://badge.fury.io/gh/tomsanbear%2Fdnslogging) ![GitHub issues](https://img.shields.io/github/issues/tomsanbear/dnslogging.svg) ![GitHub pull requests](https://img.shields.io/github/issues-pr/tomsanbear/dnslogging.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/tomsanbear/dnslogging)](https://goreportcard.com/report/github.com/tomsanbear/dnslogging)

# CoreDNS dnslogging

This CoreDNS plugin forwards all DNS transactions to a Nats server specified by the config. Currently the format sent over the line is dictated by the cap'n proto files.

WIP Document