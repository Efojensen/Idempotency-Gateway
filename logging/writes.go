package logging

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func WriteLogToFile(log string, key string) error {
	utcNow := time.Now().UTC().Format("2006-01-02")

	fileName := fmt.Sprintf("%s.log", utcNow)
	logMsg := fmt.Sprint(log + " from user with key:" + key + " at " + time.Now().Format("15:04:05"))

	file, err := os.OpenFile(
		fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if _, err := writer.WriteString(logMsg + "\n"); err != nil {
		return err
	}
	return writer.Flush()
}

func WriteAttackLogToFile(log string, key string) error {
	utcNow := time.Now().UTC().Format("2006-01-02")

	fileName := fmt.Sprintf("attacks-[%s].log", utcNow)
	logMsg := fmt.Sprint(log + key + " at " + time.Now().Format("15:04:05"))

	file, err := os.OpenFile(
		fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if _, err := writer.WriteString(logMsg + "\n"); err != nil {
		return err
	}
	return writer.Flush()
}