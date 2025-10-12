package systemdlog

import (
	"fmt"
	"log/slog"
)

type PriorityLevel int

// see: https://wiki.archlinux.org/title/Systemd/Journal#Priority_level
const (
	LevelEmergency PriorityLevel = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug
)

func (p PriorityLevel) Level() slog.Level {
	// need to reverse the mapping to match slog style levels. this maps
	// similar to slog levels where default is info and debug is negative
	return slog.Level(LevelInfo - p)
}

func priorityFromSlog(l slog.Level) PriorityLevel {
	// essentially just reverse the previous mapping
	return PriorityLevel(int(LevelInfo) - int(l))
}

func (p PriorityLevel) String() string {
	return fmt.Sprintf("<%d>", p)
}
