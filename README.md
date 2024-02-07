# FileEventLogger

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/> <img src="https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white"/> <img src="https://img.shields.io/badge/grafana-%23F46800.svg?style=for-the-badge&logo=grafana&logoColor=white"/>

Application allows you to track traffic in the folder(s). if the file is **created**, **deleted**, **changed** this information will be written to the **log file** and **prometheus** to visualize the data on the Grafana graph.


## Run Locally on windows

```bash
git clone https://github.com/Thizz00/FileEventLogger.git
```

Install Prometheus **https://prometheus.io/download/** on Windows.

Install Grafana **https://grafana.com/grafana/download/** on Windows.

Copy **prometheus.exe** and **promtool.exe** to your project.

Open a terminal and type:

```bash
prometheus.exe
```

Open a second terminal and add all paths where you want to track files/folders:

```bash
cd app
go run fileWatcher.go { ... }
```

Type url **https://localhost:9090** target  has been added and the metrics logs:

![App Screenshot](/docs/target.PNG)

Traced files folder/folders is visualized in the "graph" tab:

![App Screenshot](/docs/dashboard.PNG)

Log into grafana under localhost:3000 using login **"admin"** and password **"admin"**, then add source with prometheus.

Prometheus server URL **https://localhost:9090**.

Create your dashboard and add these **3 metrics**:

- file_create_events_total
- file_remove_events_total
- file_write_events_total

![App Screenshot](/docs/grafana.PNG)

## Features

Added possibility of combining prometehus metrics with graphana.

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
