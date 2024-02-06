package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const logFilePath = "Logs/logs.log"

var (
	writeEvents = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "file_write_events_total",
			Help: "Total number of file write events.",
		},
		[]string{"file"},
	)
	removeEvents = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "file_remove_events_total",
			Help: "Total number of file remove events.",
		},
		[]string{"file"},
	)
	createEvents = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "file_create_events_total",
			Help: "Total number of file create events.",
		},
		[]string{"file"},
	)
)

func main() {
	prometheus.MustRegister(writeEvents, removeEvents, createEvents)
	createLogsDirectory()

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatal("Error creating file watcher:", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go handleFileEvents(logger, watcher)
	addDirectoriesToWatcher(logger, watcher)
	startHTTPServer()
	<-done
}

func createLogsDirectory() {
	if _, err := os.Stat("Logs"); os.IsNotExist(err) {
		if err := os.Mkdir("Logs", 0755); err != nil {
			log.Fatal("Error creating 'Logs' directory:", err)
		}
	}
}

func handleFileEvents(logger *log.Logger, watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			handleFileEvent(logger, event)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Println("Error watching files:", err)
		}
	}
}

func handleFileEvent(logger *log.Logger, event fsnotify.Event) {
	logger.Println("Event:", event.Name)

	switch {
	case event.Op&fsnotify.Write == fsnotify.Write:
		logger.Println("Modified file:", event.Name)
		writeEvents.WithLabelValues(event.Name).Inc()
	case event.Op&fsnotify.Remove == fsnotify.Remove:
		logger.Println("Removed file:", event.Name)
		removeEvents.WithLabelValues(event.Name).Inc()
	case event.Op&fsnotify.Create == fsnotify.Create:
		logger.Println("Created file:", event.Name)
		createEvents.WithLabelValues(event.Name).Inc()
	}
}

func addDirectoriesToWatcher(logger *log.Logger, watcher *fsnotify.Watcher) {
	for _, arg := range os.Args[1:] {
		if err := filepath.Walk(arg, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				logger.Println("Error walking path:", err)
				return nil
			}
			if info.IsDir() {
				if err := watcher.Add(path); err != nil {
					logger.Println("Error adding directory to watcher:", err)
				}
			}
			return nil
		}); err != nil {
			logger.Fatal("Error walking directory:", err)
		}
	}
}

func startHTTPServer() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())

		if err := http.ListenAndServe(":9091", nil); err != nil {
			log.Fatal("Error starting HTTP server:", err)
		}
	}()
}
