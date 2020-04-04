package easyrest

import (
	"net/http"
)

// RouterGroup ...
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodGet, group.prefix+pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodPost, group.prefix+pattern, handler)
}

// DELETE defines the method to add DELETE request
func (group *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodDelete, group.prefix+pattern, handler)
}

// PATCH defines the method to add PATCH request
func (group *RouterGroup) PATCH(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodPatch, group.prefix+pattern, handler)
}

// PUT defines the method to add PUT request
func (group *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodPut, group.prefix+pattern, handler)
}

// OPTIONS defines the method to add OPTIONS request
func (group *RouterGroup) OPTIONS(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodOptions, group.prefix+pattern, handler)
}

// HEAD defines the method to add HEAD request
func (group *RouterGroup) HEAD(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodHead, group.prefix+pattern, handler)
}

// Any defines the method to add Any request
func (group *RouterGroup) Any(pattern string, handler HandlerFunc) {
	group.engine.router.add(http.MethodGet, group.prefix+pattern, handler)
	group.engine.router.add(http.MethodPost, group.prefix+pattern, handler)
	group.engine.router.add(http.MethodDelete, group.prefix+pattern, handler)
	group.engine.router.add(http.MethodPatch, group.prefix+pattern, handler)
	group.engine.router.add(http.MethodPut, group.prefix+pattern, handler)
	group.engine.router.add(http.MethodOptions, group.prefix+pattern, handler)
	group.engine.router.add(http.MethodHead, group.prefix+pattern, handler)
}
