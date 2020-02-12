package tlslistener

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/maxatome/go-testdeep"

	"github.com/sergejs-katusenoks/lets-proxy2/internal/th"
)

func TestProxyListenerType(t *testing.T) {
	listener := listenerType{connections: make(chan net.Conn)}

	// test proxy
	conn1 := &net.TCPConn{}
	go func() {
		err := listener.Put(conn1)
		if err != nil {
			t.Error(err)
		}
	}()

	rconn1, err := listener.Accept()
	if err != nil {
		t.Error()
	}
	if rconn1.(*net.TCPConn) != conn1 {
		t.Error()
	}

	conn2 := &net.TCPConn{}
	go func() {
		err := listener.Put(conn2)
		if err != nil {
			t.Error(err)
		}
	}()

	rconn2, err := listener.Accept()
	if err != nil {
		t.Error()
	}
	if rconn2.(*net.TCPConn) != conn2 {
		t.Error()
	}

	// test reject accept on close
	go func() {
		time.Sleep(time.Millisecond)
		err := listener.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	rconn3, err := listener.Accept()
	if rconn3 != nil {
		t.Error(rconn3)
	}
	if err == nil {
		t.Error()
	}

	// second close
	err = listener.Close()
	if err == nil {
		t.Error()
	}
}

func TestProxyTLS(t *testing.T) {
	defer time.Sleep(time.Second / 10)

	var body []byte
	var resp *http.Response
	ctx, flush := th.TestContext()
	defer flush()

	td := testdeep.NewT(t)

	td.FailureIsFatal()
	listenerForTLS1, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	td.CmpNoError(err)
	listenerForTLS2, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	td.CmpNoError(err)
	listenerForTCP1, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	td.CmpNoError(err)
	listenerForTCP2, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	td.CmpNoError(err)
	td.FailureIsFatal(false)

	proxy := ListenersHandler{
		GetCertificate:        dummyGetCertificate,
		ListenersForHandleTLS: []net.Listener{listenerForTLS1, listenerForTLS2},
		Listeners:             []net.Listener{listenerForTCP1, listenerForTCP2},
	}

	err = proxy.Start(ctx)
	td.CmpNoError(err)

	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		localAddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)
		_, err := proxy.GetConnectionContext(r.RemoteAddr, localAddr.String())
		td.CmpNoError(err)
		reqBytes, err := ioutil.ReadAll(r.Body)
		td.CmpNoError(err)
		if len(reqBytes) == 0 {
			_, _ = w.Write([]byte{3, 2, 1})
		} else {
			_, _ = w.Write(bytes.Repeat(reqBytes, 2))
		}
	})
	httpServer := http.Server{
		Handler: mux,
	}
	defer func() {
		_ = httpServer.Shutdown(context.Background())
	}()

	go func() {
		err := httpServer.Serve(&proxy)
		if err != nil {
			if err == http.ErrServerClosed {
				err = nil
			} else if strings.Contains(err.Error(), "listener closed") {
				err = nil
			}
		}
		td.CmpNoError(err)
	}()

	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				//nolint:gosec
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err = httpClient.Get("https://" + listenerForTLS1.Addr().String())
	td.CmpNoError(err)
	body, err = ioutil.ReadAll(resp.Body)
	td.CmpNoError(err)
	_ = resp.Body.Close()
	td.CmpDeeply(body, []byte{3, 2, 1})

	resp, err = httpClient.Get("https://" + listenerForTLS2.Addr().String())
	td.CmpNoError(err)
	body, err = ioutil.ReadAll(resp.Body)
	td.CmpNoError(err)
	_ = resp.Body.Close()
	td.CmpDeeply(body, []byte{3, 2, 1})

	resp, err = httpClient.Get("http://" + listenerForTCP1.Addr().String())
	td.CmpNoError(err)
	body, err = ioutil.ReadAll(resp.Body)
	td.CmpNoError(err)
	_ = resp.Body.Close()
	td.CmpDeeply(body, []byte{3, 2, 1})

	resp, err = httpClient.Get("http://" + listenerForTCP2.Addr().String())
	td.CmpNoError(err)
	body, err = ioutil.ReadAll(resp.Body)
	td.CmpNoError(err)
	_ = resp.Body.Close()
	td.CmpDeeply(body, []byte{3, 2, 1})

	_ = listenerForTLS1.Close()
	_ = listenerForTLS2.Close()
	_ = listenerForTCP1.Close()
	_ = listenerForTCP2.Close()
}

func dummyGetCertificate(info *tls.ClientHelloInfo) (certificate *tls.Certificate, e error) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)

	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(123),
		DNSNames:     []string{info.ServerName}, NotAfter: time.Now().Add(time.Hour),
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, certTemplate, certTemplate, key.Public(), key)
	if err != nil {
		panic(err)
	}
	leaf, err := x509.ParseCertificate(certBytes)
	certificate = &tls.Certificate{
		Leaf:        leaf,
		PrivateKey:  key,
		Certificate: [][]byte{certBytes},
	}
	return certificate, err
}
