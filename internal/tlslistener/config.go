package tlslistener

import (
	"context"
	"net"

	zc "github.com/rekby/zapcontext"
	"github.com/sergejs-katusenoks/lets-proxy2/internal/log"
	"go.uber.org/zap"
)

type Config struct {
	TLSAddresses  []string
	TCPAddresses  []string
	MinVersionTLS uint16
}

func (c Config) Apply(ctx context.Context, l *ListenersHandler) error {
	logger := zc.L(ctx)

	var tlsListeners = make([]net.Listener, 0, len(c.TLSAddresses))
	for _, addr := range c.TLSAddresses { //nolint:wsl
		listener, err := net.Listen("tcp", addr)
		log.DebugError(logger, err, "Start listen tls binding", zap.String("address", addr))
		if err != nil {
			return err
		}

		tlsListeners = append(tlsListeners, listener)
	}

	var tcpListeners = make([]net.Listener, 0, len(c.TCPAddresses))

	for _, addr := range c.TCPAddresses {
		listener, err := net.Listen("tcp", addr)
		log.DebugError(logger, err, "Start listen tcp binding", zap.String("address", addr))
		if err != nil {
			return err
		}

		tcpListeners = append(tcpListeners, listener)
	}

	// TODO: read min version from config

	l.ListenersForHandleTLS = tlsListeners
	l.Listeners = tcpListeners
	l.MinVersionTLS = 0x0303
	return nil
}
