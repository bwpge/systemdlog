// Package systemdlog is a library for writing log priority levels to stdout for systemd's journal.
package systemdlog

import (
	"fmt"
	"os"
)

// Emergencyf logs a formatted message with the 'emerg' (0) priority level.
//
// Description: System is unusable
//
// This level should not be used by applications.
func Emergencyf(f string, args ...any) {
	logf(lEmergency, f, args...)
}

// Emergency logs a message with the 'emerg' (0) priority level.
//
// Description: System is unusable
//
// This level should not be used by applications.
func Emergency(msg string) {
	logf(lEmergency, msg)
}

// Alertf logs a formatted message with the 'alert' (1) priority level.
//
// Description: Should be corrected immediately
func Alertf(f string, args ...any) {
	logf(lAlert, f, args...)
}

// Alert logs a message with the 'alert' (1) priority level.
//
// Description: Should be corrected immediately
func Alert(msg string) {
	logf(lAlert, msg)
}

// Criticalf logs a formatted message with the 'crit' (2) priority level.
//
// Description: Critical conditions; vital subsystem goes out of work, data loss, etc.
func Criticalf(f string, args ...any) {
	logf(lCritical, f, args...)
}

// Critical logs a message with the 'crit' (2) priority level.
//
// Description: Critical conditions; vital subsystem goes out of work, data loss, etc.
func Critical(msg string) {
	logf(lCritical, msg)
}

// Errorf logs a formatted message with the 'err' (3) priority level.
//
// Description: Error conditions; non-fatal error reported
func Errorf(f string, args ...any) {
	logf(lError, f, args...)
}

// Error logs a message with the 'err' (3) priority level.
//
// Description: Error conditions; non-fatal error reported
func Error(msg string) {
	logf(lError, msg)
}

// Warnf logs a formatted message with the 'warning' (4) priority level.
//
// Description: May indicate that an error will occur if action is not taken
func Warnf(f string, args ...any) {
	logf(lWarning, f, args...)
}

// Warn logs a message with the 'warning' (4) priority level.
//
// Description: May indicate that an error will occur if action is not taken
func Warn(msg string) {
	logf(lWarning, msg)
}

// Noticef logs a formatted message with the 'notice' (5) priority level.
//
// Description: Events that are unusual, but not error conditions
func Noticef(f string, args ...any) {
	logf(lNotice, f, args...)
}

// Notice logs a message with the 'notice' (5) priority level.
//
// Description: Events that are unusual, but not error conditions
func Notice(msg string) {
	logf(lNotice, msg)
}

// Infof logs a formatted message with the 'info' (6) priority level.
//
// Description: Normal operational messages that require no action
func Infof(f string, args ...any) {
	logf(lInformational, f, args...)
}

// Info logs a message with the 'info' (6) priority level.
//
// Description: Normal operational messages that require no action
func Info(msg string) {
	logf(lInformational, msg)
}

// Debugf logs a formatted message with the 'debug' (7) priority level.
//
// Description: Messages which may need to be enabled first, only useful for debugging
func Debugf(f string, args ...any) {
	logf(lDebug, f, args...)
}

// Debug logs a formatted message with the 'debug' (7) priority level.
//
// Description: Messages which may need to be enabled first, only useful for debugging
func Debug(msg string) {
	logf(lDebug, msg)
}

// Fatal is equivalent to systemdlog.Critical followed by a call to os.Exit(1).
//
// Drop-in replacement for log.Fatal.
func Fatal(err error) {
	logf(lCritical, err.Error())
	os.Exit(1)
}

// Print is equivalent to systemdlog.Info.
func Print(msg string) {
	logf(lInformational, msg)
}

// Printf is equivalent to systemdlog.Infof.
func Printf(f string, args ...any) {
	logf(lInformational, f, args...)
}

func logf(p priorityLevel, f string, args ...any) {
	fmt.Printf(p.String()+f+"\n", args...)
}
