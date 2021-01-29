package session

import (
	"errors"
	"io/ioutil"

	"github.com/ankikong/gopica/pkg/picaapi"
)

// PicaSession 保存一次会话
type PicaSession struct {
	token  string
	ip     string
	proxy  string
	client *picaapi.PicaClient
}

// NewPicaSession 构造一个新的PicaSession
func NewPicaSession(proxy, useIP string) *PicaSession {
	p := &PicaSession{
		client: picaapi.NewPicaClient(proxy, useIP),
	}
	return p
}

// Save 保存token到指定文件
func (s *PicaSession) Save(path string) error {
	return ioutil.WriteFile(path, []byte(s.token), 0644)
}

// Load 从指定文件中加载token
func (s *PicaSession) Load(path string) error {
	content, err := ioutil.ReadFile(path)
	if err == nil {
		s.token = string(content)
		if len(s.token) <= 0 {
			return errors.New("empty token")
		}
	}
	return err
}

// Login 用账号密码登录
func (s *PicaSession) Login(username, password string) (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Method("POST").Path("/auth/sign-in").AddHeader("Content-Type", "application/json; charset=UTF-8")
	req.SetSimpleMapBody(map[string]interface{}{
		"email":    username,
		"password": password,
	})
	res, err := s.client.Send(req)
	if err != nil {
		return nil, err
	}
	s.token = res.JSON().Data.Token
	return res, nil
}

// Search 搜索
func (s *PicaSession) Search(keyword string, order picaapi.ComicOrder, page uint64) (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Path("/comics/advanced-search")
	req.Method("POST").AddHeader("Content-Type", "application/json; charset=UTF-8")
	req.SetSimpleMapBody(map[string]interface{}{
		"keyword": keyword,
		"sort":    string(order),
	})
	req.AddHeader("authorization", s.token)
	req.SetPage(page)
	return s.client.Send(req)
}

// GetBlock 构建分区请求
func (s *PicaSession) GetBlock(blockName string, order picaapi.ComicOrder, page uint64) (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Method("GET")
	req.Path("/comics")
	req.SetURLParam("c", blockName)
	req.SetResultOrder(order)
	req.SetPage(page)
	req.AddHeader("authorization", s.token)
	return s.client.Send(req)
}

// GetCategory 获取主页目录
func (s *PicaSession) GetCategory() (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Method("GET")
	req.Path("/categories")
	req.AddHeader("authorization", s.token)
	return s.client.Send(req)
}

// GetEps 获取漫画的分集,支持Page
func (s *PicaSession) GetEps(id string, page uint64) (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Method("GET")
	req.Path(`comics/` + id + `/eps`)
	req.SetPage(page)
	req.AddHeader("authorization", s.token)
	return s.client.Send(req)
}

// GetImages 获取一集漫画,图片的链接,支持Page
func (s *PicaSession) GetImages(id, index string, page uint64) (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Method("GET")
	req.Path(`/comics/` + id + `/order/` + index + `/pages`)
	req.SetPage(page)
	req.AddHeader("authorization", s.token)
	return s.client.Send(req)
}

// GetMangaDetail 获取漫画详情
func (s *PicaSession) GetMangaDetail(id string) (*picaapi.PicaResponse, error) {
	req := picaapi.NewPicaRequest()
	req.Method("GET")
	req.Path("/comics/" + id)
	req.AddHeader("authorization", s.token)
	return s.client.Send(req)
}
