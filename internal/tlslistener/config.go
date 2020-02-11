package tlslistener

import (
	"context"
	"net"

	"github.com/sergejs-katusenoks/lets-proxy2/internal/log"
	zc "github.com/rekby/zapcontext"
	"go.uber.org/zap"
)

type Config struct {
	TLSAddresses []string
	TCPAddresses []string
}

func (c *Config) Apply(ctx context.Context, l *ListenersHandler) error {
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
	l.ListenersForHandleTLS = tlsListeners
	l.Listeners = tcpListeners
	return nil
}
