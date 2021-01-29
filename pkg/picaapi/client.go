package picaapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// PicaClient 哔咔客户端,对http.Client的封装
type PicaClient struct {
	client *http.Client
}

// NewPicaClient 构造一个新的客户端,proxy 和 useIP只需填一个,都不填的话,默认直连
func NewPicaClient(proxy, useIP string) *PicaClient {
	p := &PicaClient{}
	tr := &http.Transport{
		Proxy: func(hr *http.Request) (*url.URL, error) {
			if len(proxy) > 5 {
				return url.Parse(proxy)
			}
			return nil, nil
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			var dialer = &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 300 * time.Second,
			}
			if len(useIP) > 5 {
				addr = useIP
			}
			return dialer.DialContext(ctx, network, addr)
		},
	}
	p.client = &http.Client{
		Transport: tr,
	}
	return p
}

// Send 发送请求
func (c *PicaClient) Send(r *PicaRequest) (*PicaResponse, error) {

	r.sign()

	var bodyReader io.Reader = nil
	if len(r.body) > 0 {
		bodyReader = bytes.NewReader(r.body)
	}

	fullURL := "https://picaapi.picacomic.com" + r.fullPath
	req, _ := http.NewRequest(r.method, fullURL, bodyReader)

	for k, v := range r.headers {
		req.Header.Add(k, v)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return NewPicaResponseFromHTTPResponse(res)
}
