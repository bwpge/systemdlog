package systemdlog

import "strconv"

type priorityLevel int

// see: https://wiki.archlinux.org/title/Systemd/Journal#Priority_level
const (
	lEmergency priorityLevel = iota
	lAlert
	lCritical
	lError
	lWarning
	lNotice
	lInformational
	lDebug
)

func (p priorityLevel) String() string {
	return "<" + strconv.Itoa(int(p)) + ">"
}
