package memory

import (
	"sync"

	"github.com/coocood/freecache"
	"google.golang.org/protobuf/proto"
)

type call struct {
	wg  sync.WaitGroup
	val proto.Message
	err error
}

type Cache struct {
	doMu  sync.Mutex
	doMap map[string]*call

	storage *freecache.Cache
}

func NewMemoryCache() *Cache {
	cacheSize := 4096 * 1024 * 1024 // 4GB

	c := &Cache{
		doMap:   map[string]*call{},
		storage: freecache.NewCache(cacheSize),
	}

	return c
}
