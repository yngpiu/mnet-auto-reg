package main

import "time"

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func sleepSeconds(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}
