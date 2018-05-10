// +build windows

package main

import (
	"syscall"
	"time"
)

func getKeylogger() chan KeyEvent {
	user32 := syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState := user32.NewProc("GetAsyncKeyState")
	formatted := make(chan KeyEvent) //channel
	go func() {
		//store last information of key. So only react on changes. TRUE: down
		last := [256]bool{}
		for {
			time.Sleep(time.Millisecond)      //polling with 1000Hz
			for key := 0; key <= 255; key++ { //go through all keys
				Val, _, _ := getAsyncKeyState.Call(uintptr(key))
				if Val>>15&1 == 1 {
					//Key DOWN
					if !last[key] {
						//only react if the last state was not down
						last[key] = true //set state to down
						formatted <- KeyEvent{
							Code:  uint16(key),
							Value: 1,
						}
					}
				} else if Val>>15&1 == 0 {
					//Key UP
					if last[key] {
						//only react if the last state was down
						last[key] = false //set state to down
						formatted <- KeyEvent{
							Code:  uint16(key),
							Value: 0,
						}
					}
				}
			}
		}
	}()
	return formatted
}
