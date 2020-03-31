package cfull

import "fmt"

const (
	eos = "\033[0m"             // End of string
	fg  = "\033[38;2;%v;%v;%vm" // ANSI template for text foreground
	bg  = "\033[48;2;%v;%v;%vm" // ANSI template for text background
)

type RGB []uint8

type Color struct {
	ColorFG RGB
	ColorBG RGB
}

func SetRGB(text string, color Color) (result interface{}) {
	var temp []interface{}

	if color.ColorFG != nil {
		// If color has bad values return simple text
		if len(color.ColorFG) != 3 {
			return text
		}

		temp = append(temp, fmt.Sprintf(fg, color.ColorFG[0], color.ColorFG[1], color.ColorFG[2]))
	}

	if color.ColorBG != nil {
		// If color has bad values return simple text
		if len(color.ColorBG) != 3 {
			return text
		}
		temp = append(temp, fmt.Sprintf(bg, color.ColorBG[0], color.ColorBG[1], color.ColorBG[2]))
	}

	temp = append(temp, text, eos)
	var format string

	for i := 0; i < len(temp); i++ {
		format += "%s"
	}

	result = fmt.Sprintf(format, temp...)

	return
}
