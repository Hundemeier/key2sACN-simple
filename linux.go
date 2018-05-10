// +build linux

package main

import "github.com/MarinX/keylogger"

func getKeylogger() chan KeyEvent {
	devs, err := keylogger.NewDevices()
	logErr(err)

	//keyboard device file, on your system it will be diffrent!
	rd := keylogger.NewKeyLogger(devs[len(devs)-1])

	in, err := rd.Read()
	logErr(err)

	formatted := make(chan KeyEvent)
	//function that formattes to the formatted channel
	go func() {
		for i := range in {
			if i.Type == keylogger.EV_KEY {
				event := KeyEvent{
					Code:  i.Code,
					Value: i.Value,
				}
				formatted <- event
			}
		}
	}()
	return formatted
}
