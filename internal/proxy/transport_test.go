package proxy

import (
	"net/http"
	"testing"

	"github.com/sergejs-katusenoks/lets-proxy2/internal/th"

	"github.com/maxatome/go-testdeep"
)

func TestTransport_GetTransport(t *testing.T) {
	ctx, flush := th.TestContext()
	defer flush()

	td := testdeep.NewT(t)

	tr := Transport{}
	r, _ := http.NewRequest(http.MethodGet, "http://www.ru", nil)
	r = r.WithContext(ctx)
	httpTransport := tr.getTransport(r)
	td.True(httpTransport == defaultTransport) // equal pointers

	tr = Transport{IgnoreHttpsCertificate: false}
	r, _ = http.NewRequest(http.MethodGet, "https://www.ru", nil)
	r = r.WithContext(ctx)
	httpTransport = tr.getTransport(r)
	td.True(httpTransport != defaultTransport) // different pointers
	td.Cmp(httpTransport.TLSClientConfig.ServerName, "www.ru")
	td.Cmp(httpTransport.TLSClientConfig.InsecureSkipVerify, false)

	tr = Transport{IgnoreHttpsCertificate: true}
	r, _ = http.NewRequest(http.MethodGet, "https://www.ru", nil)
	r = r.WithContext(ctx)
	httpTransport = tr.getTransport(r)
	td.True(httpTransport != defaultTransport) // different pointers
	td.Cmp(httpTransport.TLSClientConfig.ServerName, "www.ru")
	td.Cmp(httpTransport.TLSClientConfig.InsecureSkipVerify, true)
}
