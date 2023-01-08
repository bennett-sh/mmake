package files

import (
	"fmt"
	"mmake/utils/mmakefile"
	"path"
	"path/filepath"
	"strings"
)

func GetSourceFiles(mmakefile mmakefile.MMakeFile) ([]string, error) {
	sourceFiles := []string{}

	for _, pattern := range strings.Split(mmakefile.SourceOptions.Directories, " ") {
		for _, namepattern := range strings.Split(mmakefile.SourceOptions.Files, " ") {
			files, err := filepath.Glob(path.Join(pattern, namepattern))
			if err != nil {
				return []string{}, err
			}

			sourceFiles = append(sourceFiles, files...)
		}
	}

	return sourceFiles, nil
}

func GetOutputFile(mmakefile mmakefile.MMakeFile) string {
	return path.Join(
		mmakefile.OutputDirectory,
		fmt.Sprintf(
			mmakefile.OutputFormat,
			mmakefile.Name,
		),
	)
}
