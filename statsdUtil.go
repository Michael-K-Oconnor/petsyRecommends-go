package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var queue = make(chan string, 100)
var statsDhost = os.Getenv("STATSD_HOST")

func init() {
	go statsdSender()
}

func StatCount(metric string, value int) {
	queue <- fmt.Sprintf("%s:%d|c", metric, value)
}

func StatTime(metric string, took time.Duration) {
	queue <- fmt.Sprintf("%s:%d|ms", metric, took/1e6)
}

func StatGauge(metric string, value int) {
	queue <- fmt.Sprintf("%s:%d|g", metric, value)
}

func statsdSender() {
	for s := range queue {
		if conn, err := net.Dial("udp", statsDhost+":8125"); err == nil {
			io.WriteString(conn, s)
			conn.Close()
		}
	}
}
