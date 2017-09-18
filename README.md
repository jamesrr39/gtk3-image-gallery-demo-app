
## Installation Instructions

### For Ubuntu 16.04
```
sudo apt update && sudo apt install -y libgtk-3-dev libcairo2-dev libglib2.0-dev
go get github.com/gotk3/gotk3/gtk
go install -tags gtk_3_18 github.com/gotk3/gotk3/gtk
```

Build:
`make`
