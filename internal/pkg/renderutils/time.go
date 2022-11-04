package renderutils

import (
	"log"
	"strconv"
	"strings"
)

func ParseTime(time string) int {
	parts := strings.Split(time, ":")
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Cannot parse time %s", time)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Cannot parse time %s", time)
	}

	return hours*60 + minutes
}
