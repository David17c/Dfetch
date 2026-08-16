package config

import (
	"strings"
)

type Values map[string]string

func Fields(format string) []string {
	var result []string

	for i := 0; i < len(format); {
		start := strings.IndexByte(format[i:], '{')
		if start == -1 {
			break
		}

		start += i

		end := strings.IndexByte(format[start:], '}')
		if end == -1 {
			break
		}

		end += start

		result = append(result, format[start+1:end])
		i = end + 1
	}

	return result
}

func Format(format string, values Values) string {
	var output strings.Builder

	for i := 0; i < len(format); {
		start := strings.IndexByte(format[i:], '{')

		if start == -1 {
			output.WriteString(format[i:])
			break
		}

		start += i

		output.WriteString(format[i:start])

		end := strings.IndexByte(format[start:], '}')
		if end == -1 {
			output.WriteString(format[start:])
			break
		}

		end += start

		key := format[start+1 : end]

		if value, ok := values[key]; ok {
			output.WriteString(value)
		} else {
			output.WriteString(format[start : end+1])
		}

		i = end + 1
	}

	return output.String()
}
