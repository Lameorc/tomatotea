# TomatoTea
Pomodoro timer CLI written in Go using [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea).

## Installation
Assuming you have a working `go` (v1.21+) installation, installing this application
is as simple as
```sh
VERSION="<your desired version, eg. latest for master>"
go install github.com/Lameorc/tomatotea@$VERSION
```

Alternative, you can build the binary yourself and move it into your `PATH` by cloning
the repository, checking out the desired commit. Then, you can issue
```sh
go build
```
which will result in `tomatotea` binary present in the current working directory.

## Usage
The interface is very simple, simply run the binary (`tomatotea`) and watch the pomodoro intervals
go by!

## Contributing
While not strictly prohibited, code contributions are discouraged as this is intended as more
of an exercise project.
However, bug reports and feature requests are welcome.
