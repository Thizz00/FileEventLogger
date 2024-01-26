package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Open the log file
	file, err := os.OpenFile("Logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file when the main function ends

	// Create a new logger that writes to the file
	logger := log.New(file, "", log.LstdFlags)

	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatal(err)
	}
	defer watcher.Close() // Close the watcher when the main function ends

	// Create a channel to signal when we're done
	done := make(chan bool)

	// Start a goroutine to handle file events and errors
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events: // Receive file events from the watcher
				if !ok {
					return
				}
				// Log the event
				logger.Println("event:", event)
				// If the event is a write event, log the modified file
				if event.Op&fsnotify.Write == fsnotify.Write {
					logger.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors: // Receive errors from the watcher
				if !ok {
					return
				}
				// Log the error
				logger.Println("error:", err)
			}
		}
	}()

	// Add the directory for argument to the watcher
	path := os.Args[1]
	err = watcher.Add(path)
	if err != nil {
		logger.Fatal(err)
	}

	// Wait until we're done
	<-done
}
