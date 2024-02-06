# FileEventLogger

## Run Locally on windows

```bash
git clone https://github.com/Thizz00/FileEventLogger.git
```

Install **https://prometheus.io/download/** OS Windows

Copy prometheus.exe,promtool.exe to your project 

Open a terminal and type:

```bash
prometheus.exe
```

Open a second terminal and add all paths where you want to track files/folders:

```go
go run fileWatcher.go { ... }
```

Type url **https://localhost:9090** target  has been added and the metrics logs:

![App Screenshot](/docs/target.PNG)

## Features

Creating metrics in prometheus:

- Total number of file create events.
- Total number of file write events.
- Total number of file remove events.

Creating logs to a log file:

- created/uploaded files/folders to the traced folders
- deleted files/folders to the traced folders
- modified files/folders to the traced folders


## License

This project is licensed under the [MIT License](LICENSE).
