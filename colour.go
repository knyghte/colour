package colour

// Constants representing the different ASCII colour codes.
const (
	colourReset        = "\033[0m"
	colourBlack        = "\033[30m"
	colourBrightBlack  = "\033[90m"
	colourRed          = "\033[31m"
	colourBrightRed    = "\033[91m"
	colourGreen        = "\033[32m"
	colourBrightGreen  = "\033[92m"
	colourYellow       = "\033[33m"
	colourBrightYellow = "\033[93m"
	colourBlue         = "\033[34m"
	colourBrightBlue   = "\033[94m"
	colourPurple       = "\033[35m"
	colourBrightPurple = "\033[95m"
	colourCyan         = "\033[36m"
	colourBrightCyan   = "\033[96m"
	colourGray         = "\033[37m"
	colourWhite        = "\033[97m"
)

// Black will colourize this text black.
func Black(text string) string {
	return colourBlack + text + colourReset
}

// BrightBlack will colourize this text bright black.
func BrightBlack(text string) string {
	return colourBrightBlack + text + colourReset
}

// Red will colourize this text red.
func Red(text string) string {
	return colourRed + text + colourReset
}

// BrightRed will colourize this text bright red.
func BrightRed(text string) string {
	return colourBrightRed + text + colourReset
}

// Green will colourize this text green.
func Green(text string) string {
	return colourGreen + text + colourReset
}

// BrightGreen will colourize this text bright green.
func BrightGreen(text string) string {
	return colourBrightGreen + text + colourReset
}

// Yellow will colourize this text yellow.
func Yellow(text string) string {
	return colourYellow + text + colourReset
}

// BrightYellow will colourize this text bright yellow.
func BrightYellow(text string) string {
	return colourBrightYellow + text + colourReset
}

// Blue will colourize this text blue.
func Blue(text string) string {
	return colourBlue + text + colourReset
}

// BrightBlue will colourize this text bright blue.
func BrightBlue(text string) string {
	return colourBrightBlue + text + colourReset
}

// Purple will colourize this text purple.
func Purple(text string) string {
	return colourPurple + text + colourReset
}

// BrightPurple will colourize this text bright purple.
func BrightPurple(text string) string {
	return colourBrightPurple + text + colourReset
}

// Cyan will colourize this text cyan.
func Cyan(text string) string {
	return colourCyan + text + colourReset
}

// BrightCyan will colourize this text bright cyan.
func BrightCyan(text string) string {
	return colourBrightCyan + text + colourReset
}

// Gray will colourize this text gray.
func Gray(text string) string {
	return colourGray + text + colourReset
}

// White will colourize this text white.
func White(text string) string {
	return colourWhite + text + colourWhite
}
