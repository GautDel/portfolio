package utils

import (
	"fmt"
	"time"
)

func DateFormat(dateStr string, formatStr string) (string, error) {
	// Parse with the correct format string to match "2025-04-03 10:01:09.596Z"
	t, err := time.Parse("2006-01-02 15:04:05.999Z", dateStr)
	if err != nil {
		return "", fmt.Errorf("Failed to parse date %s: %v", dateStr, err)
	}

	// Format the time according to the desired format
	return t.Format(formatStr), nil
}
