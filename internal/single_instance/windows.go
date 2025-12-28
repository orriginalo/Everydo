//go:build windows

package singleinstance

import (
	"fmt"
	"runtime"
	"time"

	"github.com/natefinch/npipe"
)

const pipeName = `\\.\pipe\everydo_single_instance`

func CheckSingleInstance(onSecondInstance func()) bool {
	if runtime.GOOS != "windows" {
		return true
	}
	ln, err := npipe.Listen(pipeName)
	if err != nil {
		// второй экземпляр
		conn, err := npipe.DialTimeout(pipeName, time.Second)
		if err == nil {
			conn.Write([]byte("show"))
			conn.Close()
		}
		return false
	}

	// первый экземпляр — слушаем
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			buf := make([]byte, 16)
			conn.Read(buf)
			conn.Close()
			onSecondInstance()
		}
	}()

	fmt.Println("Single instance started")
	return true
}
