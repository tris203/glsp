package protocol

import (
	"github.com/tliron/glsp"
	"sync"
)

type common_handler struct {
	initialized bool
	lock        sync.Mutex
	handlerMap  map[Method]glsp.HandlerInterface
}
