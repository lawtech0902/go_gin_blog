package article_service

import "strconv"

type Options struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Search bool   `json:"search"`
	Admin  bool   `json:"admin"`
	C      string `json:"c"` // category
	T      string `json:"t"` // tag
	Q      string `json:"q"` // 搜索的关键字
}

var defaultOptions = Options{
	Limit:  10,
	Page:   1,
	C:      "",
	T:      "",
	Q:      "",
	Search: false, // 搜索文章结果不进行缓存
	Admin:  false, // 是否是admin页面请求，如果不是，文章就不包括草稿文章
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	// 初始化默认值
	opt := defaultOptions
	
	for _, o := range opts {
		o(&opt) // 依次调用opts函数列表中的函数，为服务选项（opt变量）赋值
	}
	
	return opt
}

func SetLimitPage(limit, page string) Option {
	return func(o *Options) {
		if limit != "" && page != "" {
			p, _ := strconv.Atoi(page)
			l, _ := strconv.Atoi(limit)
			o.Limit = l
			o.Page = p
		}
	}
}

func SetAdmin(admin string) Option {
	return func(o *Options) {
		if admin != "" {
			o.Admin = true
		}
	}
}

func SetCategory(c string) Option {
	return func(o *Options) {
		o.C = c
	}
}

func SetTag(t string) Option {
	return func(o *Options) {
		o.T = t
	}
}

func SetQ(q string) Option {
	return func(o *Options) {
		o.Q = q
	}
}

func SetSearch(search bool) Option {
	return func(o *Options) {
		o.Search = search
		o.Page = defaultOptions.Page // 如果不是在第一页执行的搜索，比如：page=3，有可能会搜不到数据，必须从第一页开始搜索
	}
}
