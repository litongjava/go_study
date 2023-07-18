package middleware

import (
  "context"
  "net/http"
  "time"
)

type TimeoutMiddleware struct {
  Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if tm.Next != nil {
    tm.Next = http.DefaultServeMux
  }

  //从请求中获取context
  ctx := r.Context()
  //创建一个新的context,设置3 second的超时
  ctx, _ = context.WithTimeout(ctx, 3*time.Second)
  //使用新创建的context代替请求中的context
  r.WithContext(ctx)
  //创建一个channel,如果这个请求可以在3s内完成,该channel会收到一个信号,信号的内容是一个空的struct
  ch := make(chan struct{})
  go func() {
    //执行方法
    tm.Next.ServeHTTP(w, r)
    //执行完成后向channel发送一个信号,信号是一个空的struct,注意,这里发送的是空的struct{},发送成功之后表示一切都处理正常
    ch <- struct{}{}
  }()
  //使用select创建一个竞争的状态
  select {
  case <-ch:
    //如果按时处理完,会从ch中得到一个信号,然后return
    return
  case <-ctx.Done():
    //如果是从ctx.Done()的channel中得到信号,这表示该请求超时了,然后返回响应的状态码408
    w.WriteHeader(http.StatusRequestTimeout)
  }
  ctx.Done()
}
