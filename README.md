# key2sACN
This is a simple command-line tool for Linux and Windows that uses a keyboard as a source for DMX data.
Every key triggers a DMX channel: if the key is hold down, the channel is at 100% and if the key is released, it goes back to 0%.

Note: the DMX channels are different on every operating system and sometimes even keyboards.

The DMX data can only be send out via the network protocol sACN. Therefore some knowledge in this area is useful. Note: on Windows you can not use multicast at the moment. So you **have to** use a destination ip address: eg `key2sacn.exe -destination="192.168.2.45"`.

Use `./key2sacn -h` or `key2sacn.exe -h` as a quick help. You can exit a running program with `Ctrl+C`.
Linux: you have to use root rights: `sudo ./key2sacn`.

Note: Linux can distinguish between mutiple keyboards, but this program uses only one at a time and 
automaticly takes the first keyboard. There is currently no way to specify the keyboard to use.
For more fine-grained options on Linux use [key2sACN](https://github.com/Hundemeier/key2sACN).


This program was originally intended to use with a headless raspberry and a keyboard in an sACN environment. 
Then this combination is useful if you want some buttons (mind the n-Key rollover!) away from your console.

This program is written in go and some compiled binaries are available under [releases](https://github.com/Hundemeier/key2sACN-simple/releases).

It uses the [keylogger](https://github.com/MarinX/keylogger) by MarinX and this sACN library: [sACN](https://github.com/Hundemeier/go-sacn).