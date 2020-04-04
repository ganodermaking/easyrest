package easyrest

import (
	"net/url"
	"regexp"
	"strings"
	"sync"
)

// router 路由器结构
type router struct {
	mux     sync.RWMutex         // 读写锁
	dynamic map[string][]dynamic // 动态路由
}

// dynamic 路由结构
type dynamic struct {
	regex   *regexp.Regexp // 正则对象
	params  map[int]string // 请求参数
	handler HandlerFunc    // 响应函数
}

// newRouter 创建路由器
func newRouter() *router {
	return &router{
		sync.RWMutex{},
		make(map[string][]dynamic),
	}
}

// 注册动态路由
// 支持路由正则匹配，格式：/user/:id([0-9]+)/:name([a-z]+)
func (r *router) add(method string, pattern string, handler HandlerFunc) bool {
	if len(pattern) == 0 || handler == nil {
		return false
	}

	r.mux.Lock()
	defer r.mux.Unlock()

	params := make(map[int]string) // 请求参数
	var patterns []string          // 正则表达式组成
	pos := 0
	arr := strings.Split(pattern, "/")

	for _, v := range arr {
		if strings.HasPrefix(v, ":") {
			index := strings.Index(v, "(")
			if index != -1 {
				patterns = append(patterns, v[index:])
				params[pos] = v[1:index]
				pos++
				continue
			}
		}
		patterns = append(patterns, v)
	}

	regex, err := regexp.Compile(strings.Join(patterns, "/"))
	if err != nil {
		panic("[router]: wrong pattern \"" + pattern + "\"")
	}

	r.dynamic[method] = append(r.dynamic[method], dynamic{regex, params, handler})
	debugPrintRoute(method, pattern, handler)

	return true
}

func (r *router) handle(c *Context) {
	path := c.Req.URL.Path
	method := c.Req.Method

	if r.dynamic[method] == nil || len(r.dynamic[method]) == 0 {
		return
	}

	for _, handler := range r.dynamic[method] {
		if !handler.regex.MatchString(path) {
			continue
		}

		matches := handler.regex.FindStringSubmatch(path)
		if len(matches[0]) != len(path) {
			continue
		}

		if len(handler.params) > 0 {
			values := c.Req.URL.Query()
			for i, val := range matches[1:] {
				values.Add(handler.params[i], val)
			}
			c.Req.URL.RawQuery = url.Values(values).Encode()
		}

		handler.handler(c)
	}
}
