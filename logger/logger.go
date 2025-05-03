package logger

import (
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

var Log = logrus.New()

func InitLogger() {
    Log.SetFormatter(&logrus.TextFormatter{
        ForceColors:   true,
        FullTimestamp: true,
        TimestampFormat: "2006-01-02 15:04:05",
    })
    Log.SetOutput(os.Stdout)
    Log.SetLevel(logrus.InfoLevel)
}

func LogRequest(route string, duration time.Duration, status int, extra map[string]interface{}) {
    entry := Log.WithFields(logrus.Fields{
        "route":    route,
        "duration": duration,
        "status":   status,
    })
    for k, v := range extra {
        entry = entry.WithField(k, v)
    }
    entry.Info("Request handled")
}