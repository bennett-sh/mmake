package commandutils

import (
	"bytes"
	"fmt"
	"strings"
)

type CommandWriter struct{}

func (e CommandWriter) Write(p []byte) (int, error) {
	fmt.Println(bytes.NewBuffer(p).String())

	return len(p), nil
}

func SplitArguments(args string) []string {
	result := []string{}
	argsSplit := strings.Split(args, " ")
	i, stringBuilder := 0, new(strings.Builder)

	for _, item := range argsSplit {
		if i == 1 {
			stringBuilder.WriteString(" ")
		}
		stringBuilder.WriteString(item)

		if i == 1 {
			result = append(result, stringBuilder.String())
		}

		i = (i + 1) % 2
	}

	return result
}
