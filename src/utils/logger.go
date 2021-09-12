package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"time"
)

type LogrusFormatter struct{}

func (log *LogrusFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	msg := fmt.Sprintf(
		"%s [%s:%d][%s] %s \n",
		timestamp,
		file,
		line,
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
	return []byte(msg), nil
}
