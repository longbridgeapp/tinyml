package tinyml

import (
	"fmt"
	"strings"
)

func parseSecurityTag(val string) (counterID string, name string, err error) {
	parts := strings.Split(val, "#")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Invalid securityTag value: %v", val)
	}

	counterID = strings.TrimSpace(strings.ToUpper(parts[0]))
	name = strings.TrimSpace(parts[1])

	return counterID, name, nil
}
