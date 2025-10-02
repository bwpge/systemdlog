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
func Emergencyf(f string, v ...any) {
	logf(lEmergency, f, v...)
}

// Emergency logs a message with the 'emerg' (0) priority level.
//
// Description: System is unusable
//
// This level should not be used by applications.
func Emergency(v ...any) {
	log(lEmergency, v...)
}

// Alertf logs a formatted message with the 'alert' (1) priority level.
//
// Description: Should be corrected immediately
func Alertf(f string, v ...any) {
	logf(lAlert, f, v...)
}

// Alert logs a message with the 'alert' (1) priority level.
//
// Description: Should be corrected immediately
func Alert(v ...any) {
	log(lAlert, v...)
}

// Criticalf logs a formatted message with the 'crit' (2) priority level.
//
// Description: Critical conditions; vital subsystem goes out of work, data loss, etc.
func Criticalf(f string, v ...any) {
	logf(lCritical, f, v...)
}

// Critical logs a message with the 'crit' (2) priority level.
//
// Description: Critical conditions; vital subsystem goes out of work, data loss, etc.
func Critical(v ...any) {
	log(lCritical, v...)
}

// Errorf logs a formatted message with the 'err' (3) priority level.
//
// Description: Error conditions; non-fatal error reported
func Errorf(f string, v ...any) {
	logf(lError, f, v...)
}

// Error logs a message with the 'err' (3) priority level.
//
// Description: Error conditions; non-fatal error reported
func Error(v ...any) {
	log(lError, v...)
}

// Warnf logs a formatted message with the 'warning' (4) priority level.
//
// Description: May indicate that an error will occur if action is not taken
func Warnf(f string, v ...any) {
	logf(lWarning, f, v...)
}

// Warn logs a message with the 'warning' (4) priority level.
//
// Description: May indicate that an error will occur if action is not taken
func Warn(v ...any) {
	log(lWarning, v...)
}

// Noticef logs a formatted message with the 'notice' (5) priority level.
//
// Description: Events that are unusual, but not error conditions
func Noticef(f string, v ...any) {
	logf(lNotice, f, v...)
}

// Notice logs a message with the 'notice' (5) priority level.
//
// Description: Events that are unusual, but not error conditions
func Notice(v ...any) {
	log(lNotice, v...)
}

// Infof logs a formatted message with the 'info' (6) priority level.
//
// Description: Normal operational messages that require no action
func Infof(f string, v ...any) {
	logf(lInformational, f, v...)
}

// Info logs a message with the 'info' (6) priority level.
//
// Description: Normal operational messages that require no action
func Info(v ...any) {
	log(lInformational, v...)
}

// Debugf logs a formatted message with the 'debug' (7) priority level.
//
// Description: Messages which may need to be enabled first, only useful for debugging
func Debugf(f string, v ...any) {
	logf(lDebug, f, v...)
}

// Debug logs a formatted message with the 'debug' (7) priority level.
//
// Description: Messages which may need to be enabled first, only useful for debugging
func Debug(v ...any) {
	log(lDebug, v...)
}

// Fatal is equivalent to systemdlog.Critical followed by a call to os.Exit(1).
//
// Drop-in replacement for log.Fatal.
func Fatal(v ...any) {
	log(lCritical, v...)
	os.Exit(1)
}

// Fatalf is equivalent to systemdlog.Criticalf followed by a call to os.Exit(1).
//
// Drop-in replacement for log.Fatalf.
func Fatalf(f string, v ...any) {
	logf(lCritical, f, v...)
	os.Exit(1)
}

func logf(p priorityLevel, f string, v ...any) {
	fmt.Printf(p.String()+f+"\n", v...)
}

func log(p priorityLevel, v ...any) {
	fmt.Print(p.String())
	fmt.Print(v...)
	fmt.Print("\n")
}
