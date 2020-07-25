package service

import (
	"fmt"
	"github.com/aka-achu/go-web/logging"
	"os"
)

type Sample struct{}

func (*Sample) GetHostName() (string, error) {
	if hostname, err := os.Hostname(); err != nil {
		logging.AppLogger.Errorf("Failed to fetch the hostname. Error- %s", err.Error())
		return "", err
	} else {
		logging.AppLogger.Info("Successfully fetched the hostname")
		return fmt.Sprintf("Hello from Host- %s", hostname), nil
	}
}
