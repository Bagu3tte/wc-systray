# WC system tray

A system tray to obtain the status of the toilets in the Franklin Building.

### Installing

#### Windows

- Download the latest version of the executable:
- Copy it in `C:\Users\%username%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`

#### Mac OS

#### Linux

### Dependencies

- `go get ./...`

### Build

#### Windows

- `go build -ldflags -H=windowsgui -o wc-systray.exe *.go`
