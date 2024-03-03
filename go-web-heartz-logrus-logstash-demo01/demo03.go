package main

import (
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	log := logrus.New()
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}
	//formatter := logrustash.DefaultFormatter(logrus.Fields{"appName": "myappName"})
	//formatter := logrus.JSONFormatter{
	//  TimestampFormat: "2006-01-02 15:04:05",
	//}
	formatter := logrus.TextFormatter{}
	hook := logrustash.New(conn, &formatter)

	log.Hooks.Add(hook)
	log.Info("Hello World!")
}
