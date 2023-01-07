package commands

import (
	"fmt"
	"mmake/utils/commandutils"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	mmakefile, err := Compile(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("running %s...\n", mmakefile.Name)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command(path.Join(cwd, fmt.Sprintf(mmakefile.OutputFormat, mmakefile.Name)))

	fmt.Println((ctx.String("arguments")))

	cmd.Args = commandutils.SplitArguments(ctx.String("arguments"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		return err
	}

	// FIXME: passing args doesn't work

	return nil
}
