package gackup

import (
	"log"
	"os"
	"path/filepath"

	"github.com/richardwilkes/toolbox/errs"
)

type FileStatus uint8

const (
	StatusMissing FileStatus = iota
	StatusLink
	StatusRegular
)

func (fs FileStatus) String() string {
	switch fs {
	case StatusLink:
		return "Link"
	case StatusMissing:
		return "Missing"
	case StatusRegular:
		return "Regular"
	}
	return ""
}

type FileAction uint8

const (
	ActionSkip FileAction = iota
	ActionSkipBecauseMissing
	ActionCopyAndLink
	ActionLink
	ActionRelink
	ActionUnhandled
)

func (f FileAction) String() string {
	switch f {
	case ActionSkip:
		return "already linked, skipping"
	case ActionCopyAndLink:
		return "copy and link"
	case ActionLink:
		return "link"
	case ActionRelink:
		return "change link destination"
	case ActionSkipBecauseMissing:
		return "Skip because missing"
	case ActionUnhandled:
		return "This combination is currently unhandled"
	}
	return ""
}

type FileDetails struct {
	source, destination, linkDestination string
	sourceStatus, destinationStatus      FileStatus
	action                               FileAction
}

func (f *FileDetails) Action() FileAction {
	switch f.destinationStatus {
	case StatusRegular, StatusLink:
		switch f.sourceStatus {
		case StatusLink:
			if f.linkDestination == f.destination {
				return ActionSkip
			}
			return ActionRelink
		case StatusRegular, StatusMissing:
			return ActionLink
		}

	case StatusMissing:
		switch f.sourceStatus {
		case StatusRegular, StatusLink:
			return ActionCopyAndLink
		case StatusMissing:
			return ActionSkipBecauseMissing
		}
	}
	log.Printf("unhandled: %s %s\n", f.sourceStatus, f.destinationStatus)
	return ActionUnhandled
}

func Status(f string) (FileStatus, error) {
	stat, err := os.Lstat(f)
	if err != nil {
		if !os.IsNotExist(err) {
			return StatusMissing, errs.Wrap(err)
		}
		return StatusMissing, nil
	}
	m := stat.Mode()
	if m.IsRegular() || m.IsDir() {
		return StatusRegular, nil
	}
	return StatusLink, nil
}

func NewFileDetails(fn, base, configDir string) (*FileDetails, error) {
	source := filepath.Join(base, fn)
	destination := filepath.Join(base, configDir, fn)

	srcStatus, err := Status(source)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	var linkDestination string
	if srcStatus == StatusLink {
		linkDestination, err = os.Readlink(source)
		if err != nil {
			return nil, errs.Wrap(err)
		}
	}

	destinationStatus, err := Status(destination)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return &FileDetails{
		source:            source,
		sourceStatus:      srcStatus,
		linkDestination:   linkDestination,
		destination:       destination,
		destinationStatus: destinationStatus,
	}, nil
}
