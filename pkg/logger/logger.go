package logger

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, fields ...interface{})
	Info(message string, fields ...interface{})
	Warn(message string, fields ...interface{})
	Error(message interface{}, fields ...interface{})
	Fatal(message interface{}, fields ...interface{})
}

// Logger -.
type Logger struct {
	mu       sync.Mutex
	logger   zerolog.Logger
	logFile  *lumberjack.Logger
	level    zerolog.Level
	filePath string
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string) *Logger {
	logger := &Logger{}
	logger.initLogger(level)

	// Start a goroutine to monitor date changes and rotate logs at midnight
	go logger.autoRotateLogs(level)

	return logger
}

// initLogger -.
func (l *Logger) initLogger(level string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Define the log file with the current date
	currentDate := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/%s.log", currentDate)

	// Define log rotation settings
	l.logFile = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // MB
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   false,
	}

	// Save the current log file path for compression later
	l.filePath = filename

	// Set log level
	l.setLogLevel(level)

	// Multi-writer: log to file and console
	multiWriter := zerolog.MultiLevelWriter(l.logFile, os.Stdout)

	// Initialize the logger
	skipFrameCount := 3
	l.logger = zerolog.New(multiWriter).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
		Logger()
}

// compressLogFile -.
func (l *Logger) compressLogFile() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	oldLogFile := l.filePath
	compressedFile := fmt.Sprintf("%s.gz", oldLogFile)

	inputFile, err := os.Open(oldLogFile)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(compressedFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := gzip.NewWriter(outputFile)
	defer writer.Close()

	_, err = io.Copy(writer, inputFile)
	if err != nil {
		return err
	}

	// Remove the original uncompressed log file after compression
	return os.Remove(oldLogFile)
}

// autoRotateLogs -.
func (l *Logger) autoRotateLogs(level string) {
	for {
		// Get the current time
		now := time.Now()

		// Calculate time until midnight
		nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		duration := time.Until(nextMidnight)

		// Wait until midnight
		time.Sleep(duration)

		// Compress the old log file before rotating
		if err := l.compressLogFile(); err != nil {
			fmt.Printf("Failed to compress log file: %v\n", err)
		}

		// Reinitialize the logger with the new day's log file
		l.initLogger(level)
	}
}

// setLogLevel -.
func (l *Logger) setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		l.level = zerolog.DebugLevel
	case "info":
		l.level = zerolog.InfoLevel
	case "warn":
		l.level = zerolog.WarnLevel
	case "error":
		l.level = zerolog.ErrorLevel
	case "fatal":
		l.level = zerolog.FatalLevel
	default:
		l.level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l.level)
}

// Debug -.
func (l *Logger) Debug(message interface{}, fields ...interface{}) {
	l.msg("debug", message, fields...)
}

// Info -.
func (l *Logger) Info(message string, fields ...interface{}) {
	l.log(message, fields...)
}

// Warn -.
func (l *Logger) Warn(message string, fields ...interface{}) {
	l.log(message, fields...)
}

// Error -.
func (l *Logger) Error(message interface{}, fields ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, fields...)
	}

	l.msg("error", message, fields...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, fields ...interface{}) {
	l.msg("fatal", message, fields...)

	os.Exit(1)
}

// Fatalf -.
func (l *Logger) Fatalf(message interface{}, args ...interface{}) {
	l.msgf("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, fields ...interface{}) {
	l.logger.Info().Fields(fields).Msg(message)
}

func (l *Logger) msg(level string, message interface{}, fields ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), fields...)
	case string:
		l.log(msg, fields...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), fields...)
	}
}

func (l *Logger) logf(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

func (l *Logger) msgf(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.logf(msg.Error(), args...)
	case string:
		l.logf(msg, args...)
	default:
		l.logf(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
