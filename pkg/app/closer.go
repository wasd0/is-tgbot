package app

import (
	"context"
	"fmt"
	"is-tgbot/pkg/logger"
	"strings"
	"sync"
)

type Callback func(context.Context) error

type Closer struct {
	mutex  sync.Mutex
	buffer []Callback
}

func (c *Closer) Add(callback Callback) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.buffer = append(c.buffer, callback)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var (
		msgs     = make([]string, 0, len(c.buffer))
		complete = make(chan struct{}, 1)
	)

	go func() {
		for _, callback := range c.buffer {
			if err := callback(ctx); err != nil {
				msgs = append(msgs, fmt.Sprintf("[Closer] callback error: %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-complete:
		logger.Log().Infof("[Closer] Processed %d messages", len(c.buffer))
		break
	case <-ctx.Done():
		return fmt.Errorf("[Closer]: shutdown error:  %v", ctx.Err())
	}

	if len(msgs) > 0 {
		return fmt.Errorf(strings.Join(msgs, "\n"))
	}

	return nil
}
