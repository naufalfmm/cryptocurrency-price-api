package ws

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Engine struct {
	urlHandlers map[string][]Handler
}

func New() *Engine {
	return &Engine{
		urlHandlers: make(map[string][]Handler),
	}
}

func (e *Engine) Dial(url string, handlers ...Handler) {
	if e.urlHandlers == nil {
		e.urlHandlers = make(map[string][]Handler)
	}

	if _, ok := e.urlHandlers[url]; !ok {
		e.urlHandlers[url] = []Handler{}
	}

	e.urlHandlers[url] = append(e.urlHandlers[url], handlers...)
}

func (e *Engine) Start() error {
	var wg sync.WaitGroup
	wg.Add(len(e.urlHandlers))

	for url, handlers := range e.urlHandlers {
		c, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return err
		}
		defer c.Close()

		done := make(chan bool)
		go func(d chan bool) {
			defer close(done)
			for {
				wctx := context{}
				_, message, err := c.ReadMessage()
				if err != nil {
					wctx.err = err
					close(done)
				}

				wctx.msg = message

				for _, handler := range handlers {
					handler(wctx)
				}
			}
		}(done)

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				wg.Done()
			case <-ticker.C:
			}
		}
	}

	wg.Wait()

	return nil
}
