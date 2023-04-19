package mm

import (
	"google.golang.org/protobuf/proto"
)

type CacheCaller func() (proto.Message, error)

type IMemoryCache interface {
	Get(key string, value proto.Message, caller CacheCaller, expireSeconds int) error
}
