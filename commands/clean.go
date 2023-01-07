package commands

import (
	"mmake/utils"
	"mmake/utils/mmakefile"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func Clean(ctx *cli.Context) error {
	mmakefilename, err := mmakefile.GetMMakefile(ctx)
	if err != nil {
		return err
	}

	mmakefile, err := mmakefile.ParseFile(mmakefilename)
	if err != nil {
		return err
	}

	// if build dir is project dir
	if path.Clean(mmakefile.OutputDirectory) == "." {
		err = os.Remove(utils.GetOutputFile(*mmakefile))
	} else {
		err = os.RemoveAll(mmakefile.OutputDirectory)
	}

	if err != nil {
		return err
	}

	return nil
}
