package helper

import (
	"fmt"
)

// Log contact
type Log interface {
	LogError(err error)
}

// LogImpl struct for config
type LogImpl struct {
}

// NewLog for init
func NewLog() Log {
	return &LogImpl{}
}

func (x LogImpl) LogError(err error) {
	fmt.Println("[ERROR] ", err)
}
