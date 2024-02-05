# FileEventLogger

## Features

Creating metrics in prometheus:

- Total number of file create events.
- Total number of file write events.
- Total number of file remove events.

Creating logs to a log file:

- created/uploaded files/folders to the traced folders
- deleted files/folders to the traced folders
- modified files/folders to the traced folders

## Open script

Add all paths where you want to track files/folders:

```go
go run fileWatcher.go { ... }
```


## License

This project is licensed under the [MIT License](LICENSE).
