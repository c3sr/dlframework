package feature

import (
	"github.com/c3sr/dlframework"
)

func isUnknownType(o *dlframework.Feature) bool {
	return o.Type == dlframework.FeatureType_UNKNOWN
}
