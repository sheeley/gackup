package gackup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/richardwilkes/toolbox/errs"
)

type Config struct {
	BaseDir, ConfigDir              string
	ForceRelink, Verbose, ShowSkips bool
}

var DefaultConfig = &Config{
	BaseDir:   os.Getenv("HOME"),
	ConfigDir: "Documents/config",
}

func (c *Config) Check() error {
	var err error
	if c.BaseDir == "" {
		err = errs.Append(errs.New("BaseDir must not be empty"))
	}
	if c.ConfigDir == "" {
		err = errs.Append(errs.New("ConfigDig must not be empty"))
	}
	return err
}

func New(files []string, c *Config) (*Backup, error) {
	b := &Backup{}

	if c == nil {
		c = DefaultConfig
	}

	if err := c.Check(); err != nil {
		return nil, errs.Wrap(err)
	}

	b.config = c

	for _, fn := range files {
		fd, err := NewFileDetails(fn, b.config.BaseDir, b.config.ConfigDir)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		if b.config.Verbose {
			fmt.Printf("%+v\n", fd)
		}
		b.fds = append(b.fds, fd)
	}

	return b, nil
}

type Backup struct {
	fds    []*FileDetails
	config *Config
}

func (b *Backup) Proposed() (string, error) {
	return b.do(false, true)
}

func (b *Backup) Move() (string, error) {
	return b.do(true, false)
}

func (b *Backup) do(doActions, showActions bool) (string, error) {
	o := strings.Builder{}

	for _, fd := range b.fds {

		a := fd.Action()
		if b.config.Verbose {
			o.WriteString(a.String() + " " + fd.source + "\n")
		}
		if a == ActionSkip {
			if !b.config.ForceRelink {
				if b.config.Verbose || b.config.ShowSkips {
					o.WriteString(fmt.Sprintf("SKIP: %s\n", fd.source))
				}
				continue
			}
			a = ActionLink
		}

		link := (a == ActionLink || a == ActionRelink)
		if a == ActionCopyAndLink {
			link = true
			if b.config.Verbose || showActions {
				o.WriteString(fmt.Sprintf("MOVE: %s -> %s\n", fd.source, fd.destination))
			}

			if doActions {
				d := filepath.Dir(fd.destination)
				err := os.MkdirAll(d, 0700)
				if err != nil {
					o.WriteString(d)
					return o.String(), errs.Wrap(err)
				}
				err = os.Rename(fd.source, fd.destination)
				if err != nil {
					o.WriteString(fmt.Sprintf(fd.source, fd.destination))
					return o.String(), errs.Wrap(err)
				}
			}
		}

		if link {
			if b.config.Verbose || showActions {
				actionDescription := "LINK"
				if a == ActionRelink {
					actionDescription = "RELINK"
				}
				o.WriteString(fmt.Sprintf("%s: %s -> %s\n", actionDescription, fd.source, fd.destination))
			}

			if doActions {
				err := os.Remove(fd.source)
				if err != nil && !os.IsNotExist(err) {
					return o.String(), errs.Wrap(err)
				}
				err = os.Symlink(fd.destination, fd.source)
				if err != nil {
					o.WriteString(fmt.Sprintf(fd.source, fd.destination))
					return o.String(), errs.Wrap(err)
				}
			}
		}
	}

	return o.String(), nil
}
