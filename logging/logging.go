package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var globalLogger Logger

type Logger struct {
	logger        *log.Logger
	fileDirectory string
	dateLayout    string
	fileDate      string
	outputToFile  bool
}

const (
	defaultDateLayout = "2006-01-02"
	defaultDirectory  = "log"
	logToFile         = "LOG_TO_FILE"
)

func InitLogger() {
	globalLogger.initLogger(defaultDirectory, defaultDateLayout)
}

func (logger *Logger) initLogger(dir string, dateLayout string) {
	logger.fileDirectory = dir
	logger.dateLayout = dateLayout
	logger.fileDate = time.Now().UTC().Format(logger.dateLayout)
	logger.outputToFile = true
	if logOutputEnv, ok := os.LookupEnv(logToFile); ok {
		outputToFile, err := strconv.ParseBool(logOutputEnv)
		if err == nil {
			logger.outputToFile = outputToFile
		}
	}
	logger.setOutput()
}

func (logger *Logger) setOutput() {
	if logger.outputToFile {
		currentDate := time.Now().UTC().Format(logger.dateLayout)
		// if log to file and no date or logger set
		if logger.fileDate != currentDate || logger.logger == nil {
			logger.fileDate = currentDate
			// we need to create a new logger and log file
			f, err := logger.createLogFile()
			if err != nil {
				logger.logger = log.New(os.Stdout, "", 0)
				logger.outputToFile = false
				log.Printf("failed to create log file %v, logging to stdout", err)
				return
			}
			logger.logger = log.New(f, "", 0)
		}
	} else {
		if logger.logger == nil {
			logger.logger = log.New(os.Stdout, "", 0)
		}
	}
}

// Functions accessible outside the package that
// don't require a logger method
func Logf(format string, messages ...interface{}) {
	globalLogger.logf(format, messages...)
}
func Log(message string) {
	globalLogger.log(message)
}

func (logger *Logger) log(message string) {
	logger.setOutput()
	logger.logger.Printf("%v %v \n", logPrefix(), message)
}

func (logger *Logger) logf(format string, messages ...interface{}) {
	logger.log(fmt.Sprintf(format, messages...))
}

func logPrefix() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05")
}

func (logger *Logger) createLogFile() (*os.File, error) {
	if _, err := os.Stat(logger.fileDirectory); os.IsNotExist(err) {
		err = os.Mkdir(logger.fileDirectory, 0750)
		if err != nil {
			return nil, err
		}
	}
	// Join the filepath, pass into OpenFile as name
	// 0600 is for a user readable+writable file
	// 0644 is for a user readable+writable file and group readable file
	return os.OpenFile(filepath.Join(logger.fileDirectory, fmt.Sprintf("scs-log-%v.log", logger.fileDate)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
}
