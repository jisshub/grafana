package postgres

import (
	sdkproxy "github.com/grafana/grafana-plugin-sdk-go/backend/proxy"
)

// createPostgresProxyDriver creates and registers a new sql driver that uses a postgres connector and updates the dialer to
// route connections through the secure socks proxy
func createPostgresProxyDriver(cnnstr string, opts *sdkproxy.Options) (string, error) {
	panic("FIXME: not implemented")
	return "", nil
}
