package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sse"
	"net/http"
	"time"
)

func main() {
	h := server.Default()

	h.GET("/sse", SseHandlerDemo)

	h.Spin()

}

func SseHandlerDemo(ctx context.Context, c *app.RequestContext) {
	//模拟测试数据
	times := time.NewTicker(1 * time.Second).C
	// client can tell server last event it received with Last-Event-ID header
	lastEventID := sse.GetLastEventID(c)
	hlog.CtxInfof(ctx, "last event ID: %s", lastEventID)
	// you must set status code and response headers before first render call
	c.SetStatusCode(http.StatusOK)
	s := sse.NewStream(c)
	for t := range times {
		event := &sse.Event{
			Event: "timestamp",
			Data:  []byte(t.Format(time.RFC3339)),
		}
		err := s.Publish(event)
		if err != nil {
			return
		}
	}

}
