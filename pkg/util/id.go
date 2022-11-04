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

func IsSameObject(a model.ObjectID, b model.ObjectID) bool {
	return (sanitize(a.ExtId) != "" && sanitize(a.ExtId) == sanitize(b.ExtId)) ||
		(sanitize(a.Db640) != "" && sanitize(a.Db640) == sanitize(b.Db640)) ||
		(a.Ifopt.Country == b.Ifopt.Country && a.Ifopt.State == b.Ifopt.State && a.Ifopt.Stop == b.Ifopt.Stop)
}

func sanitize(value string) string {
	value = strings.Trim(value, " ")
	return strings.ReplaceAll(value, "  ", " ")
}
