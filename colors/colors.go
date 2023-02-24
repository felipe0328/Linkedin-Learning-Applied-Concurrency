package colors

import "fmt"

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func SprintFColor(color, value string, params ...any) string{
	return fmt.Sprintf("%s%s%s", color, fmt.Sprintf(value, params...), Reset)
}

func PrintlnColor(color, value string) {
	fmt.Printf("%s%s%s\n", color, value, Reset)
}

func PrintFColor(color, value string, params ...any) {
	fmt.Printf("%s%s%s", color, fmt.Sprintf(value, params...), Reset)
}
