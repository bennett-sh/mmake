package compilation

import (
	"mmake/utils/mmakefile"
	"path"
	"path/filepath"
	"strings"
)

func GetFiles(mmakefile mmakefile.MMakeFile) ([]string, error) {
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
