package utils

import (
	"fmt"
	"mmake/utils/mmakefile"
	"path"
)

func GetOutputFile(mmakefile mmakefile.MMakeFile) string {
	return path.Join(
		mmakefile.OutputDirectory,
		fmt.Sprintf(
			mmakefile.OutputFormat,
			mmakefile.Name,
		),
	)
}
