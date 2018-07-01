Set-Variable GOOS=linux
Set-Variable GOARCH=amd64
go build -o build/key2sACN-linux-64

Set-Variable GOOS=linux
Set-Variable GOARCH=386
go build -o build/key2sACN-linux-32

Set-Variable GOOS=windows
Set-Variable GOARCH=amd64
go build -o build/key2sACN-windows-64.exe

Set-Variable GOOS=windows
Set-Variable GOARCH=386
go build -o build/key2sACN-windows-32.exe

Set-Variable GOOS=darwin
Set-Variable GOARCH=amd64
go build -o build/key2sACN-macos-64

Set-Variable GOOS=darwin
Set-Variable GOARCH=386
go build -o build/key2sACN-macos-32