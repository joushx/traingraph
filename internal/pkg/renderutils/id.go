package renderutils

import (
	"strings"

	"github.com/joushx/traingraph/pkg/model"
)

func IsSameObject(a model.ObjectID, b model.ObjectID) bool {
	return (sanitize(a.ExtId) != "" && sanitize(a.ExtId) == sanitize(b.ExtId)) ||
		(sanitize(a.Db640) != "" && sanitize(a.Db640) == sanitize(b.Db640)) ||
		(a.Ifopt.Country == b.Ifopt.Country && a.Ifopt.State == b.Ifopt.State && a.Ifopt.Stop == b.Ifopt.Stop)
}

func sanitize(value string) string {
	value = strings.Trim(value, " ")
	return strings.ReplaceAll(value, "  ", " ")
}
