## Draco encoder GUI

Build command for windows from Linux
```
sudo apt-get install mingw-w64
export CC=x86_64-w64-mingw32-gcc
export CGO_ENABLED=1
GOOS=windows GOARCH=amd64 go build -o dracogui.exe
```