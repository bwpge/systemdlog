// Package systemdlog is a library for writing log priority levels to stdout for systemd's journal.
package systemdlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
	ctx    context.Context
)

func init() {
	ctx = context.Background()
	SetLogLevel(LevelDebug)
}

// Emergencyf logs a formatted message with the 'emerg' (0) priority level.
//
// Description: System is unusable
//
// This level should not be used by applications.
func Emergencyf(f string, v ...any) {
	logf(LevelEmergency, f, v...)
}

// Emergency logs a message with the 'emerg' (0) priority level.
//
// Description: System is unusable
//
// This level should not be used by applications.
func Emergency(v ...any) {
	logln(LevelEmergency, v...)
}

// Alertf logs a formatted message with the 'alert' (1) priority level.
//
// Description: Should be corrected immediately
func Alertf(f string, v ...any) {
	logf(LevelAlert, f, v...)
}

// Alert logs a message with the 'alert' (1) priority level.
//
// Description: Should be corrected immediately
func Alert(v ...any) {
	logln(LevelAlert, v...)
}

// Criticalf logs a formatted message with the 'crit' (2) priority level.
//
// Description: Critical conditions; vital subsystem goes out of work, data loss, etc.
func Criticalf(f string, v ...any) {
	logf(LevelCritical, f, v...)
}

// Critical logs a message with the 'crit' (2) priority level.
//
// Description: Critical conditions; vital subsystem goes out of work, data loss, etc.
func Critical(v ...any) {
	logln(LevelCritical, v...)
}

// Errorf logs a formatted message with the 'err' (3) priority level.
//
// Description: Error conditions; non-fatal error reported
func Errorf(f string, v ...any) {
	logf(LevelError, f, v...)
}

// Error logs a message with the 'err' (3) priority level.
//
// Description: Error conditions; non-fatal error reported
func Error(v ...any) {
	logln(LevelError, v...)
}

// Warnf logs a formatted message with the 'warning' (4) priority level.
//
// Description: May indicate that an error will occur if action is not taken
func Warnf(f string, v ...any) {
	logf(LevelWarning, f, v...)
}

// Warn logs a message with the 'warning' (4) priority level.
//
// Description: May indicate that an error will occur if action is not taken
func Warn(v ...any) {
	logln(LevelWarning, v...)
}

// Noticef logs a formatted message with the 'notice' (5) priority level.
//
// Description: Events that are unusual, but not error conditions
func Noticef(f string, v ...any) {
	logf(LevelNotice, f, v...)
}

// Notice logs a message with the 'notice' (5) priority level.
//
// Description: Events that are unusual, but not error conditions
func Notice(v ...any) {
	logln(LevelNotice, v...)
}

// Infof logs a formatted message with the 'info' (6) priority level.
//
// Description: Normal operational messages that require no action
func Infof(f string, v ...any) {
	logf(LevelInfo, f, v...)
}

// Info logs a message with the 'info' (6) priority level.
//
// Description: Normal operational messages that require no action
func Info(v ...any) {
	logln(LevelInfo, v...)
}

// Debugf logs a formatted message with the 'debug' (7) priority level.
//
// Description: Messages which may need to be enabled first, only useful for debugging
func Debugf(f string, v ...any) {
	logf(LevelDebug, f, v...)
}

// Debug logs a formatted message with the 'debug' (7) priority level.
//
// Description: Messages which may need to be enabled first, only useful for debugging
func Debug(v ...any) {
	logln(LevelDebug, v...)
}

// Fatal is equivalent to systemdlog.Critical followed by a call to os.Exit(1).
//
// Drop-in replacement for log.Fatal.
func Fatal(v ...any) {
	logln(LevelCritical, v...)
	os.Exit(1)
}

// Fatalf is equivalent to systemdlog.Criticalf followed by a call to os.Exit(1).
//
// Drop-in replacement for log.Fatalf.
func Fatalf(f string, v ...any) {
	logf(LevelCritical, f, v...)
	os.Exit(1)
}

// Log is a simple pass-through function to the underlying slog.Logger's Log method
func Log(level PriorityLevel, msg string, v ...any) {
	logger.Log(ctx, level.Level(), msg, v...)
}

func SetLogLevel(level PriorityLevel) {
	opts := slog.HandlerOptions{Level: level}
	handler := NewSystemdHandler(os.Stdout, opts)
	logger = slog.New(handler)
}

// SetLevelToStringFunc sets the function called by PriorityLevel.String().
// Useful for changing the level output style at runtime with e.g., based on a
// command line option.
func SetLevelToStringFunc(f func(PriorityLevel) string) {
	priorityLevelToString = f
}

func logln(level slog.Leveler, v ...any) {
	logger.Log(ctx, level.Level(), fmt.Sprint(v...))
}

func logf(level slog.Leveler, f string, v ...any) {
	logger.Log(ctx, level.Level(), fmt.Sprintf(f, v...))
}
