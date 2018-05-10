# key2sACN
This is a simple command-line tool for Linux that uses a keyboard as a source for DMX data.
Every key triggers a DMX channel: if the key is hold down, the channel is at 100% and if the key is released, it goes back to 0%.

The DMX data can only be send out via the network protocol sACN. Therefore some knowledge in this area is useful.

Use `./key2sacn -h` as a quick help. You can exit a running program with `Ctrl+C`.

This program is written in go and some compiled binaries are available under [releases](https://github.com/Hundemeier/key2sACN/releases).