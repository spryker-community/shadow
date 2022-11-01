package cmd

import (
	"github.com/pkg/errors"
	"github.com/symfony-cli/console"
	"shadow/internal/filesystem"
	"shadow/internal/io"
	"shadow/internal/project"
)

type initCommand struct {
	cc *console.Command
}

func newInitCommand() *initCommand {
	cmd := &initCommand{}

	cmd.cc = &console.Command{
		Name:   "init",
		Usage:  "Init directories",
		Action: adoptActionFunc(cmd.initAction, false),
	}

	return cmd
}

func (i *initCommand) initAction(_ *console.Context, prj *project.Project) error {
	if exists, _ := filesystem.DirExists(prj.Fs, prj.ShadowDir); exists {
		return errors.Errorf(`Module directory already exists at "%s"`, prj.ShadowDir)
	}

	if err := prj.Fs.Mkdir(prj.ShadowDir, 0755); err != nil {
		return errors.Wrapf(err, `Unable to create module directory at "%s"`, prj.ShadowDir)
	}

	io.Write(`<info>Successfully created module directory at "%s"</info>`, prj.ShadowDir)

	return nil
}
