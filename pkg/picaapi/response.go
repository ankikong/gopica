package picaapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PicaResponse 哔咔请求后的返回结果
type PicaResponse struct {
	// Content 请求后,http的body
	Content []byte
	// StatusCode http的状态码
	StatusCode int
	// Header 请求后的响应头
	Header http.Header
}

// NewPicaResponseFromHTTPResponse 从http相应构建一个新的PicaResponse
func NewPicaResponseFromHTTPResponse(h *http.Response) (*PicaResponse, error) {
	r := &PicaResponse{}
	var err error
	r.Content, err = ioutil.ReadAll(h.Body)
	if err != nil {
		return nil, err
	}
	r.Header = h.Header
	return r, nil
}

// JSON 把数据格式化为json
func (p *PicaResponse) JSON() PicaJSON {
	var ret PicaJSON
	json.Unmarshal(p.Content, &ret)
	return ret
}
