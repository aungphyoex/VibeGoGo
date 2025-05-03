package logger

import (
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

var log = logrus.New()

func Init() {
    log.SetFormatter(&logrus.TextFormatter{
        ForceColors:     true,
        FullTimestamp:   true,
        TimestampFormat: "2006-01-02 15:04:05",
    })
    log.SetOutput(os.Stdout)
    log.SetLevel(logrus.InfoLevel)
}

func Info(args ...interface{}) {
    log.Info(args...)
}

func Error(args ...interface{}) {
    log.Error(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
    return log.WithFields(fields)
}

// Example for logging HTTP requests with color and details
func RequestLog(route string, duration time.Duration, status int, method string, remote string) {
    log.WithFields(logrus.Fields{
        "route":    route,
        "duration": duration,
        "status":   status,
        "method":   method,
        "remote":   remote,
    }).Info("Request handled")
}