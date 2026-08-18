package format

import (
	"strings"
)

type Values map[string]string

func Fields(format string) []string {
	var result []string

	for i := 0; i < len(format); {
		if i+1 < len(format) && format[i] == '$' && format[i+1] == '{' {
			end := strings.IndexByte(format[i+2:], '}')
			if end == -1 {
				break
			}

			i = i + 2 + end + 1
			continue
		}

		if format[i] == '{' {
			end := strings.IndexByte(format[i+1:], '}')
			if end == -1 {
				break
			}

			end += i + 1

			result = append(result, format[i+1:end])
			i = end + 1
			continue
		}

		i++
	}

	return result
}

func Format(format string, values Values) string {
	var output strings.Builder

	for i := 0; i < len(format); {
		if i+1 < len(format) && format[i] == '$' && format[i+1] == '{' {
			end := strings.IndexByte(format[i+2:], '}')
			if end == -1 {
				output.WriteString(format[i:])
				break
			}

			end += i + 2

			output.WriteString(format[i : end+1])
			i = end + 1
			continue
		}

		if format[i] == '{' {
			end := strings.IndexByte(format[i+1:], '}')
			if end == -1 {
				output.WriteString(format[i:])
				break
			}

			end += i + 1

			key := format[i+1 : end]

			if value, ok := values[key]; ok {
				output.WriteString(value)
			} else {
				output.WriteString(format[i : end+1])
			}

			i = end + 1
			continue
		}

		output.WriteByte(format[i])
		i++
	}

	return output.String()
}
