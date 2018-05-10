# key2sACN
This is a simple command-line tool for Linux that uses a keyboard as a source for DMX data.
Every key triggers a DMX channel: if the key is hold down, the channel is at 100% and if the key is released, it goes back to 0%.

The DMX data can only be send out via the network protocol sACN. Therefore some knowledge in this area is useful.

Use `./key2sacn -h` as a quick help. You can exit a running program with `Ctrl+C`.

This program is intended to use with a headless raspberry and a keyboard in an sACN environment. Then this 
combination is useful if you want some buttons (mind the n-Key rollover!) away from your console.

This program is written in go and some compiled binaries are available under [releases](https://github.com/Hundemeier/key2sACN/releases).

It uses the [keylogger](https://github.com/MarinX/keylogger) by MarinX and this sACN library: [sACN](https://github.com/Hundemeier/go-sacn).