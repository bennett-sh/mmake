package commands

import (
	"fmt"
	"io/fs"
	"mmake/utils/files"
	"mmake/utils/mmakefile"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli/v2"
)

func Compile(ctx *cli.Context) (*mmakefile.MMakeFile, error) {
	mmakefilename, err := mmakefile.GetMMakefile(ctx)

	if err != nil {
		return nil, err
	}

	mmakefile, err := mmakefile.ParseFile(mmakefilename)
	if err != nil {
		return nil, err
	}

	compiler := mmakefile.CC
	flags := mmakefile.CCFlags
	name := mmakefile.Name
	output := path.Join(mmakefile.OutputDirectory, name)

	err = os.MkdirAll(mmakefile.OutputDirectory, fs.ModeDir)
	if err != nil {
		return mmakefile, err
	}

	fmt.Println("Compiling...")

	files, err := files.GetSourceFiles(*mmakefile)
	if err != nil {
		return mmakefile, err
	}

	args := []string{"-o" + output}

	if len(flags) > 0 {
		args = append(args, flags)
	}
	args = append(args, files...)

	for _, include := range mmakefile.Include {
		args = append(args, "-I")
		args = append(args, include)
	}

	cmd := exec.Command(compiler, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Printf("compilation error: %s\n", err.Error())
		return nil, err
	}

	fmt.Println("Done.")

	return mmakefile, nil
}
