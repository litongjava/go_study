package main

import "testing"

func TestLogrus(t *testing.T) {
	InitLog()
	Log.Infoln("测试日志文件")
}
