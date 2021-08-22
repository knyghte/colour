package colour

// Constants representing the different ASCII colour codes.
const (
	colourReset  = "\033[0m"
	colourRed    = "\033[31m"
	colourGreen  = "\033[32m"
	colourYellow = "\033[33m"
	colourBlue   = "\033[34m"
	colourPurple = "\033[35m"
	colourCyan   = "\033[36m"
	colourGray   = "\033[37m"
	colourWhite  = "\033[97m"
)

// Red will colourize this text red.
func Red(text string) string {
	return colourRed + text + colourReset
}

// Green will colourize this text green.
func Green(text string) string {
	return colourGreen + text + colourReset
}

// Yellow will colourize this text yellow.
func Yellow(text string) string {
	return colourYellow + text + colourReset
}

// Blue will colourize this text blue.
func Blue(text string) string {
	return colourBlue + text + colourReset
}

// Purple will colourize this text purple.
func Purple(text string) string {
	return colourPurple + text + colourReset
}

// Cyan will colourize this text cyan.
func Cyan(text string) string {
	return colourCyan + text + colourReset
}

// Gray will colourize this text gray.
func Gray(text string) string {
	return colourGray + text + colourReset
}

// White will colourize this text white.
func White(text string) string {
	return colourWhite + text + colourWhite
}
