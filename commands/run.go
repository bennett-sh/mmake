package commands

import (
	"fmt"
	"mmake/utils"
	"os"
	"os/exec"
	"path"

	"github.com/google/shlex"
	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	mmakefile, err := Compile(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Running %s...\n", mmakefile.Name)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command(
		path.Join(
			cwd,
			utils.GetOutputFile(*mmakefile),
		),
	)

	runArgsSplit, err := shlex.Split(ctx.String("arguments"))
	if err != nil {
		return err
	}

	cmd.Args = runArgsSplit
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
