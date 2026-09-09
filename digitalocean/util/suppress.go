package util

import (
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CaseSensitive implements a schema.SchemaDiffSuppressFunc that ignores case
func CaseSensitive(_, old, new string, _ *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

// Float32Precision suppresses diffs caused by float64 config values being
// stored/returned as float32 by the API (e.g. 0.001 vs 0.0010000000474974513).
func Float32Precision(_, old, new string, _ *schema.ResourceData) bool {
	oldF, errOld := strconv.ParseFloat(old, 64)
	newF, errNew := strconv.ParseFloat(new, 64)
	if errOld != nil || errNew != nil {
		return false
	}
	return float32(oldF) == float32(newF)
}
