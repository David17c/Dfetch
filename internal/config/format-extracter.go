package config

import "strings"

func ExtractFormat(format string) []string {
	var result []string

	for {
		start := strings.IndexByte(format, '{')
		if start == -1 {
			break
		}

		end := strings.IndexByte(format[start:], '}')
		if end == -1 {
			break
		}
		end += start

		value := format[start+1 : end]

		result = append(result, value)

		format = format[end+1:]
	}

	return result
}
