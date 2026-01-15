package helpers

import (
	"log"
)

// Color constants
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// LogSuccess logs success messages in green
func LogSuccess(format string, args ...interface{}) {
	log.Printf(ColorGreen+format+ColorReset, args...)
}

// LogError logs error messages in red
func LogError(format string, args ...interface{}) {
	log.Printf(ColorRed+format+ColorReset, args...)
}

// LogWarning logs warning messages in yellow
func LogWarning(format string, args ...interface{}) {
	log.Printf(ColorYellow+format+ColorReset, args...)
}

// LogInfo logs info messages in blue
func LogInfo(format string, args ...interface{}) {
	log.Printf(ColorBlue+format+ColorReset, args...)
}

// LogDebug logs debug messages in cyan
func LogDebug(format string, args ...interface{}) {
	log.Printf(ColorCyan+format+ColorReset, args...)
}

// LogCritical logs critical messages in bold red
func LogCritical(format string, args ...interface{}) {
	log.Printf(ColorBold+ColorRed+format+ColorReset, args...)
}
