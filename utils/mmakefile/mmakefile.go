package mmakefile

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/urfave/cli/v2"
)

func GetMMakefile(ctx *cli.Context) (string, error) {
	mmakefile := "MMakefile"

	if ctx.Args().Len() > 0 {
		firstArg := ctx.Args().First()
		_, statErr := os.Stat(firstArg)

		if statErr == nil {
			mmakefile = firstArg
		} else if errors.Is(statErr, fs.ErrNotExist) {
			return "", cli.Exit(fmt.Sprintf("mmakefile not found (%s)", firstArg), 1)
		} else {
			cli.Exit("no mmakefile found", 1)
		}
	}

	if _, err := os.Stat(mmakefile); errors.Is(err, fs.ErrNotExist) {
		return "", cli.Exit("no mmakefile found", 1)
	}

	return mmakefile, nil
}