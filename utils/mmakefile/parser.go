package mmakefile

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

type MMakeFile struct {
	Name            string        `yaml:"name"`
	CC              string        `yaml:"cc"`
	CCFlags         string        `yaml:"ccFlags"`
	OutputFormat    string        `yaml:"outputFormat"`
	OutputDirectory string        `yaml:"outputDirectory"`
	SourceOptions   SourceOptions `yaml:"sourceOptions"`
	Include         []string      `yaml:"include"`
}

type SourceOptions struct {
	Files       string `yaml:"files" default:""`
	Directories string `yaml:"directories" default:""`
}

func detectDefaultOutputFormat() string {
	switch runtime.GOOS {
	case "windows":
		return "%s.exe"

	default:
		return "./%s"
	}
}

func ParseFile(file string) (*MMakeFile, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	mmakefile := &MMakeFile{}
	err = yaml.Unmarshal(buf, mmakefile)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", file, err)
	}

	if len(mmakefile.OutputFormat) < 1 {
		mmakefile.OutputFormat = detectDefaultOutputFormat()
	}

	if len(mmakefile.SourceOptions.Directories) < 1 {
		mmakefile.SourceOptions.Directories = "source/ src/ sources/ source/** src/** sources/**"
	}

	if len(mmakefile.SourceOptions.Files) < 1 {
		mmakefile.SourceOptions.Files = "*.cpp *.cc *.c *.c++"
	}

	if len(mmakefile.Include) < 1 {
		mmakefile.Include = []string{"include", "headers"}
	}

	if len(mmakefile.OutputDirectory) < 1 {
		mmakefile.OutputDirectory = "bin"
	}

	if len(mmakefile.CC) < 1 {
		mmakefile.CC = "g++"
	}

	return mmakefile, nil
}
