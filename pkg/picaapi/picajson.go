package picaapi

// Thumb 所有图片格式,链接格式为 FileServer + "/static/" + Path
type Thumb struct {
	FileServer   string `json:"fileServer"`
	Path         string `json:"path"`
	OriginalName string `json:"originalName"`
}

// Category 每个分区的信息
type Category struct {
	Active bool   `json:"active"`
	IsWeb  bool   `json:"isWeb"`
	Link   string `json:"link"`
	Title  string `json:"title"`
	Thumb  Thumb  `json:"thumb"`
}

// Comic 搜索和分区中,每一条记录的数据格式
type Comic struct {
	Author     string   `json:"author"`
	Categories []string `json:"categories"`
	EpsCount   uint32   `json:"epsCount"`
	Finished   bool     `json:"finished"`
	LikesCount uint32   `json:"likesCount"`
	PagesCount uint32   `json:"pagesCount"`
	Thumb      Thumb    `json:"thumb"`
	Title      string   `json:"title"`
	ID         string   `json:"_id"`
}

// User 用户数据格式
type User struct {
	Avatar     Thumb    `json:"avatar"`
	Character  string   `json:"character"`
	Characters []string `json:"characters"`
	Exp        uint32   `json:"exp"`
	Gender     string   `json:"gender"`
	Level      int32    `json:"level"`
	Name       string   `json:"name"`
	Role       string   `json:"role"`
	Slogan     string   `json:"slogan"`
	Title      string   `json:"title"`
	Verified   bool     `json:"verified"`
	ID         string   `json:"_id"`
}

// ComicDetail 漫画详细信息的格式
type ComicDetail struct {
	AllowComment  bool     `json:"allowComment"`
	AllowDownload bool     `json:"allowDownload"`
	Author        string   `json:"author"`
	Categories    []string `json:"categories"`
	ChineseTeam   string   `json:"chineseTeam"`
	CommentsCount string   `json:"commentsCount"`
	CreatedAt     string   `json:"created_at"`
	Description   string   `json:"description"`
	EpsCount      uint32   `json:"epsCount"`
	Finished      bool     `json:"finished"`
	IsFavourite   bool     `json:"isFavourite"`
	IsLiked       bool     `json:"isLiked"`
	LikesCount    uint32   `json:"likesCount"`
	PagesCount    uint32   `json:"pagesCount"`
	Tags          []string `json:"tags"`
	Thumb         Thumb    `json:"thumb"`
	Title         string   `json:"title"`
	TotalLikes    uint32   `json:"totalLikes"`
	TotalViews    uint32   `json:"totalViews"`
	UpdatedAt     string   `json:"updated_at"`
	ViewsCount    uint32   `json:"viewsCount"`
	Creator       User     `json:"_creator"`
	ID            string   `json:"_id"`
}

// Ep 获取comic时会返回的简单数据
type Ep struct {
	ID        string `json:"id"`
	Order     uint32 `json:"order"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
}

type common struct {
	Limit uint32 `json:"limit"`
	Page  uint32 `json:"page"`
	Pages uint32 `json:"pages"`
	Total uint32 `json:"total"`
}

// Image 一集每张图片的格式
type Image struct {
	ID    string `json:"id"`
	Media Thumb  `json:"media"`
}

// PicaJSON 哔咔的所有json格式,统一到一个struct里面
type PicaJSON struct {
	// Code 相应状态码
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data 返回数据
	Data struct {
		// Categories 首页目录
		Categories []Category `json:"categories"` // adddd
		// Comic 获取一部漫画详情时返回
		Comic ComicDetail `json:"comic"`
		// Comics 分区和搜索时返回的数据格式
		Comics struct {
			Docs []Comic `json:"docs"`
			*common
		} `json:"comics"`
		// Eps 漫画获取分集的时候返回
		Eps struct {
			Docs []Ep `json:"docs"`
			*common
		} `json:"eps"`
		// Pages 获取一集的所有图片时返回
		Pages struct {
			Docs []Image `json:"docs"`
			*common
		} `json:"pages"`
		// 登陆时返回的数据
		Token string `json:"token"`
	} `json:"data"`
}
