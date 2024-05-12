package ormtypes

import "sync"

type gMap struct {
	mp map[string]string
	*sync.RWMutex
}

// 生成安全map
func NewSafeMap() *gMap {
	return &gMap{
		mp:      make(map[string]string),
		RWMutex: &sync.RWMutex{},
	}
}

// 长度
func (g *gMap) Len() int {
	return len(g.mp)
}

// 是否存在
func (g *gMap) Exist(key string) bool {
	g.RLock()
	defer g.Unlock()
	_, ok := g.mp[key]
	return ok
}

// 获取
func (g *gMap) Get(key string) string {
	g.RLock()
	defer g.Unlock()
	return g.mp[key]
}

// 设置
func (g *gMap) Set(key string, value string) {
	g.RLock()
	defer g.Unlock()
	if value == "" {
		g.mp[key] = value
	}
}
