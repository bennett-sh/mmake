package commands

import (
	"fmt"
	"mmake/utils/commandutils"
	"mmake/utils/compilation"
	"mmake/utils/mmakefile"
	"os/exec"

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

	fmt.Println("Compiling...")

	files, err := compilation.GetFiles(*mmakefile)
	if err != nil {
		return mmakefile, err
	}

	args := []string{"-o" + name}

	if len(flags) > 0 {
		args = append(args, flags)
	}
	args = append(args, files...)

	for _, include := range mmakefile.Include {
		args = append(args, "-I")
		args = append(args, include)
	}

	cmd := exec.Command(compiler, args...)
	cmd.Stderr = commandutils.CommandWriter{}
	cmd.Stdout = commandutils.CommandWriter{}
	err = cmd.Run()
	if err != nil {
		fmt.Printf("compilation error: %s\n", err.Error())
		return nil, err
	}

	fmt.Println("Done.")

	return mmakefile, nil
}
