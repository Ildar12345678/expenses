package log

import (
	"log"
	"os"
	"fmt"
	"sync"
)

var (
	logger *log.Logger
	lock   = &sync.Mutex{}
)

func NewLog() (*log.Logger, error) {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()
		if logger == nil {
			file, err := os.OpenFile("server_log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				return nil, fmt.Errorf("error while creating log file: %s", err.Error())
			}
			
			logger := log.New(file, "", log.LstdFlags)
			
			logger.SetFlags(log.LstdFlags | log.Llongfile)
		}
	}
	return logger, nil
}
