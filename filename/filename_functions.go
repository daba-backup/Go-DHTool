package filename

import (
	"strings"
)

func ReplaceWindowsDelimiterWithLinuxDelimiter(filename string) string {
	return strings.Replace(filename, "\\", "/", -1)
}
func GetFileDirectory(filename string) string {
	split := strings.Split(filename, "/")
	split_num := len(split)

	directory := ""
	for i := 0; i < split_num-1; i++ {
		directory += split[i]
		if i != split_num-2 {
			directory += "/"
		}
	}

	return directory
}
func GetFileExtension(filename string) string {
	split := strings.Split(filename, ".")
	split_num := len(split)

	extension := split[split_num-1]
	return extension
}
func GetFilenameWithoutExtension(filename string) string {
	split := strings.Split(filename, ".")
	split_num := len(split)

	ret := ""
	for i := 0; i < split_num-1; i++ {
		ret += split[i]
	}

	return ret
}
func GetFilenameWithoutDirectory(filename string) string {
	split := strings.Split(filename, "/")
	split_num := len(split)

	return split[split_num-1]
}
