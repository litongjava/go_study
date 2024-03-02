package main

import (
	"context"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	// SetLevel sets the level of logs below which logs will not be output.
	hlog.SetLevel(hlog.LevelDebug)
	f, err := os.Create("hertz.log")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// SetOutput sets the output of default logger. By default, it is stderr.
	//hlog.SetOutput(f)
	// if you want to output the log to the file and the stdout at the same time, you can use the following codes
	fileWriter := io.MultiWriter(f, os.Stdout)
	hlog.SetOutput(fileWriter)

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		// it will be output
		hlog.Info("Hello, hertz")
		// it will not be output
		hlog.Trace("Hello, hertz")
		c.String(consts.StatusOK, "Hello hertz!")
	})

	h.Spin()
}
