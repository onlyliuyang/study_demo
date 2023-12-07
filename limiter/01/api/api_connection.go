package api

import (
	"context"
	"golang.org/x/time/rate"
	"time"
)

type ApiConnection struct {
	diskLimit    RateLimiter
	networkLimit RateLimiter
	apiLimit     RateLimiter
}

func Open() *ApiConnection {
	secordLimit := rate.NewLimiter(Per(2, time.Second), 1)
	minuteLimit := rate.NewLimiter(Per(10, time.Minute), 10)
	return &ApiConnection{
		apiLimit:     MultiLimiter(secordLimit, minuteLimit),
		diskLimit:    MultiLimiter(rate.NewLimiter(rate.Limit(1), 1)),
		networkLimit: MultiLimiter(rate.NewLimiter(Per(3, time.Second), 3)),
	}
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	////return rate.Limit()
	//switch second {
	//case time.Second:
	//	//return rate.Every(2 * time.Second)
	//	//return rate.Limit(i)
	//	//return rate.Limit(i)
	//	return rate.Every(2 * time.Second)
	//case time.Minute:
	//	//return rate.Every(10 * time.Minute)
	//	//return rate.Limit(i / 60)
	//	return rate.Every(10 / 60 * time.Second)
	//}
	//return rate.Limit(0)
	return rate.Every(duration / time.Duration(eventCount))
}

func (a *ApiConnection) ReadFile(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *ApiConnection) ResolveAddress(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx); err != nil {
		return err
	}
	return nil
}
