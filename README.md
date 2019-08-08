# WC system tray

A system tray to obtain the status of the toilets in the Franklin Building.

![example](https://bitbucket.org/fkiene/wc-systray/downloads/demo.png)

### Installing

#### Windows

- Download the latest version of the executable: https://bitbucket.org/fkiene/wc-systray/downloads/wc-systray.exe
- Copy it in `C:\Users\%username%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`

#### Mac OS

#### Linux

- `sudo apt-get install libgtk-3-dev libappindicator3-dev`

### Build

#### Windows

- `go build -ldflags -H=windowsgui -o wc-systray.exe *.go`

#### From Windows to Mac OS (in progress)

- `GOOS=darwin GOARCH=amd64 go build -o wc-systray.app/Contents/MacOS/wc-systray *.go`

#### From Windows to linux (in progress)

- `GOOS=linux GOARCH=amd64 go build -o wc-systray *.go`
