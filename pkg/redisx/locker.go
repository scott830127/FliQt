package redisx

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidislock"
)

func NewLocker(cfg Config) (*Locker, func(), error) {
	locker := new(Locker)
	if len([]string{cfg.Addr}) == 0 {
		return locker, func() {}, errors.New("locker addresses is empty")
	}
	redisLocker, err := rueidislock.NewLocker(rueidislock.LockerOption{
		ClientOption: rueidis.ClientOption{
			InitAddress:  []string{cfg.Addr},
			Username:     "",
			Password:     cfg.Password,
			SelectDB:     cfg.DB,
			DisableCache: true,
			ShuffleInit:  len([]string{cfg.Addr}) > 1,
		},
		NoLoopTracking: true,
	})
	if err != nil {
		return locker, func() {}, err
	}
	locker.locker = redisLocker
	return locker, locker.Close, nil
}

type Locker struct {
	locker rueidislock.Locker
}

// Get a distributed redis lock by key by waiting for it.
func (c *Locker) Get(ctx context.Context, key string) (unlock context.CancelFunc, err error) {
	if c.locker == nil {
		unlock = func() {}
		return
	}
	_, unlock, err = c.locker.WithContext(ctx, key)
	if unlock == nil {
		unlock = func() {}
	}
	if err != nil {
		err = fmt.Errorf("get lock err")
	}
	return
}

// TryGet  a distributed redis lock by key without waiting. It may return ErrNotLocked.
func (c *Locker) TryGet(ctx context.Context, key string) (unlock context.CancelFunc, err error) {
	if c.locker == nil {
		unlock = func() {}
		return
	}
	_, unlock, err = c.locker.TryWithContext(ctx, key)
	if unlock == nil {
		unlock = func() {}
	}
	if err != nil {
		err = fmt.Errorf("get lock err")
	}
	return
}

// AntiReSubmit 防重送
func (c *Locker) AntiReSubmit(ctx context.Context, key string) (unlock context.CancelFunc, err error) {
	if c.locker == nil {
		unlock = func() {}
		return
	}
	_, unlock, err = c.locker.TryWithContext(ctx, key)
	if unlock == nil {
		unlock = func() {}
	}
	if err == nil {
		return
	}
	if errors.Is(err, rueidislock.ErrNotLocked) {
		err = fmt.Errorf("get lock err")
		return
	}
	err = fmt.Errorf("get lock err")
	return
}

func (c *Locker) Close() {
	if c.locker != nil {
		c.locker.Close()
		c.locker = nil
	}
}
