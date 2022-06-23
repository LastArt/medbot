package pkg

import (
	"log"
	"os"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func Logger(some error) {
	f, err := os.OpenFile(LogJournalPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "ℹ️ LOG:", log.LstdFlags)
	logger.Println(some)
}
