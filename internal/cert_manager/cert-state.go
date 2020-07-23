//nolint:golint
package cert_manager

import (
	"context"
	"crypto/tls"
	"sync"

	zc "github.com/rekby/zapcontext"

	"github.com/sergejs-katusenoks/lets-proxy2/internal/cache"
	"github.com/sergejs-katusenoks/lets-proxy2/internal/log"

	"go.uber.org/zap"
)

type certState struct {
	mu sync.RWMutex

	issueContext       context.Context // nil if no issue process now
	issueContextCancel func()
	cert               *tls.Certificate
	useAsIs            bool // certificate locked by flag. It deny renew and some internal checks.
	lastError          error
}

// Try to lock state for issue certificate.
// It return true if success.
// It return true if state already locked for issue
func (s *certState) StartIssue(ctx context.Context) (res bool) {
	defer func() {
		// defer for log outside of lock mutex
		zc.L(ctx).Debug("StartAutoRenew issue lock", zap.Bool("result", res))
	}()

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.issueContext == nil {
		s.issueContext, s.issueContextCancel = context.WithCancel(context.Background())
		return true
	}
	return false
}

// Must call after StartIssue issue certificate
func (s *certState) FinishIssue(ctx context.Context, cert *tls.Certificate, lastError error) {
	logger := zc.L(ctx)

	if cert != nil && lastError != nil || cert == nil && lastError == nil {
		logger.DPanic("Must be cert exactly one: cert or last error. Cert set as nil.", zap.Error(lastError),
			log.Cert(cert))
		cert = nil
	}

	s.mu.Lock()
	s.cert = cert
	oldContext, oldCancel := s.issueContext, s.issueContextCancel
	s.issueContext, s.issueContextCancel, s.lastError = nil, nil, lastError
	s.mu.Unlock()

	if oldContext == nil {
		zc.L(ctx).DPanic("Finish issue certificate without start it.")
	} else {
		logger.Debug("Finish issue lock.")
		oldCancel()
	}
}

func (s *certState) WaitFinishIssue(ctx context.Context) (cert *tls.Certificate, err error) {
	logger := zc.L(ctx)
	logger.Debug("StartAutoRenew waiting to finish certificate issue.")

	s.mu.RLock()
	issueContext, cert, err := s.issueContext, s.cert, s.lastError
	s.mu.RUnlock()

	if issueContext == nil {
		return cert, err
	}
	select {
	case <-ctx.Done():
		err = ctx.Err()
		logger.Warn("Certificate issue waiting context canceled.", zap.Error(err))
		return nil, err
	case <-issueContext.Done():
		cert, err = s.Cert()
		logger.Debug("Waiting for certificate issue finished", log.Cert(cert))
		return cert, err
	}
}

func (s *certState) Cert() (cert *tls.Certificate, lastError error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cert = s.cert
	if cert == nil && s.lastError == nil {
		lastError = cache.ErrCacheMiss
	} else {
		lastError = s.lastError
	}

	return cert, lastError
}

func (s *certState) CertSet(ctx context.Context, locked bool, cert *tls.Certificate) {
	zc.L(ctx).Debug("Store certificate in local state", log.Cert(cert))

	s.mu.Lock()
	s.cert = cert
	s.useAsIs = locked
	s.lastError = nil
	s.mu.Unlock()
}

func (s *certState) GetUseAsIs() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.useAsIs
}
