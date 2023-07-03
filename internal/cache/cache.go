package cache

import "github.com/coocood/freecache"

// Cache 缓存单例
var Cache *freecache.Cache

// init
// @description: 初始化缓存
func init() {
	if Cache == nil {
		cacheSize := 100 * 1024 * 1024
		Cache = freecache.NewCache(cacheSize)
	}
}
