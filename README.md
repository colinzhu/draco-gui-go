# STL to DRC Converter GUI

This is a simple GUI application built with the Fyne library in Go. It allows you to convert a batch of STL files in a folder to DRC format.

## Features

- Select a folder containing STL files for conversion
- Set the quality parameter for the conversion
- View the output of the conversion process

## Installation

To install this application, you need to have Go installed on your machine. Once you have Go installed, you can clone this repository and build the application:

### Build command for Linux
```bash
git clone https://github.com/colinzhu/draco-gui-go.git
cd draco-gui-go
go build -o draco-gui
cp draco-gui ./linux
```


### Build command for windows from Linux
```
sudo apt-get install mingw-w64
export CC=x86_64-w64-mingw32-gcc
export CGO_ENABLED=1
GOOS=windows GOARCH=amd64 go build -o draco-gui.exe
copy draco-gui.exe win/
```

## Usage
To use this application, run the `draco-gui` executable:
```bash
# linux
./draco-gui

# windows
draco-gui.exe
```

Then, click the "Select folder..." button to select a folder containing STL files. Enter the quality parameter for the conversion in the text box, and click the "Convert" button to start the conversion process. The output of the conversion process will be displayed in the application window.


