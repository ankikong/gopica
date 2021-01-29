package picaapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ComicOrder 漫画顺序
type ComicOrder string

const (
	// ComicOrderDEFALUE 默认顺序
	ComicOrderDEFALUE ComicOrder = "ua"
	// ComicOrderNEWEST 按时间新到旧
	ComicOrderNEWEST ComicOrder = "dd"
	// ComicOrderOLDEST 按时间旧到新
	ComicOrderOLDEST ComicOrder = "da"
	// ComicOrderPOINT 最多指名
	ComicOrderPOINT ComicOrder = "vd"
	// ComicOrderLOVE 最多收藏
	ComicOrderLOVE ComicOrder = "ld"
)

// PicaRequest 构造哔咔请求
type PicaRequest struct {
	headers  map[string]string
	path     string
	method   string
	params   url.Values
	body     []byte
	fullPath string
	order    ComicOrder
	client   *http.Client
}

// NewPicaRequest 构建一个新的请求
func NewPicaRequest() (r *PicaRequest) {
	r = &PicaRequest{}
	r.headers = map[string]string{
		"accept":            "application/vnd.picacomic.com.v1+json",
		"app-channel":       "1",
		"app-uuid":          "defaultUuid",
		"app-platform":      "android",
		"app-version":       "2.2.1.2.3.3",
		"app-build-version": "44",
		"User-Agent":        "okhttp/3.8.1",
		"image-quality":     "original", // 默认原画
	}

	return r
}

// SetImageQuality 设置获取图片的清晰度,默认是原图
func (r *PicaRequest) SetImageQuality(q string) *PicaRequest {
	r.AddHeader("image-quality", q)
	return r
}

// SetResultOrder 设置漫画顺序
func (r *PicaRequest) SetResultOrder(order ComicOrder) *PicaRequest {
	r.order = order
	return r
}

// AddHeader 添加http头,相同key会覆盖
func (r *PicaRequest) AddHeader(k, v string) *PicaRequest {
	r.headers[k] = v
	return r
}

// SetToken 身份验证token
func (r *PicaRequest) SetToken(token string) *PicaRequest {
	r.headers["authorization"] = token
	return r
}

// SetURLParam 设置url参数
func (r *PicaRequest) SetURLParam(k, v string) *PicaRequest {
	r.params.Set(k, v)
	return r
}

// SetPage 设置跳转页面
func (r *PicaRequest) SetPage(page uint64) *PicaRequest {
	if page == 0 {
		page = 1
	}
	r.params.Set("page", strconv.FormatUint(page, 10))
	return r
}

// Path 设置http的path
func (r *PicaRequest) Path(p string) *PicaRequest {
	r.path = p
	return r
}

// Method http请求method
func (r *PicaRequest) Method(m string) *PicaRequest {
	r.method = m
	return r
}

// SetBytesBody 设定post时的body
func (r *PicaRequest) SetBytesBody(b string) *PicaRequest {
	r.body = []byte(b)
	return r
}

// SetSimpleMapBody 设置简单单层的json
func (r *PicaRequest) SetSimpleMapBody(data map[string]interface{}) *PicaRequest {
	r.body, _ = json.Marshal(data)
	return r
}

// sign 私有校验方法
func (r *PicaRequest) sign() {
	var ts = strconv.FormatInt(time.Now().Unix(), 10)
	var nonce = "b1ab87b4800d4d4590a11701b8551afa"
	var apiKey = "C69BAF41DA5ABD1FFEDC6D2FEA56B"
	var secret = "~d}$Q7$eIni=V)9\\RK/P.RM4;9[7|@/CA}b~OW!3?EV`:<>M7pddUBL5n|0/*Cn"

	r.AddHeader("time", ts)
	r.AddHeader("nonce", nonce)
	r.AddHeader("api-key", apiKey)

	var raw strings.Builder
	raw.WriteString(r.path)
	p := r.params.Encode()
	if len(p) > 0 {
		raw.WriteString("?")
		raw.WriteString(p)
	}
	r.fullPath = raw.String()

	raw.WriteString(ts)
	raw.WriteString(nonce)
	raw.WriteString(r.method)
	raw.WriteString(apiKey)

	finalRaw := strings.ToLower(raw.String())
	finalRaw = finalRaw[1:]
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(finalRaw))
	r.AddHeader("signature", hex.EncodeToString(h.Sum(nil)))
}
