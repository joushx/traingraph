package util

import (
	"strings"

	"github.com/joushx/traingraph/pkg/model"
)

func NewIfOpt(value string) model.Ifopt {
	result := model.Ifopt{}
	parts := strings.Split(value, ":")

	if len(parts) > 0 {
		result.Country = parts[0]
	}

	if len(parts) > 1 {
		result.State = parts[1]
	}

	if len(parts) > 2 {
		result.Stop = parts[2]
	}

	if len(parts) > 3 {
		result.Area = parts[3]
	}

	if len(parts) > 4 {
		result.Platform = parts[4]
	}

	return result
}
