package feature

import "github.com/c3sr/dlframework"

func RawType() Option {
	return Type(dlframework.FeatureType_RAW)
}

func Raw(e *dlframework.Raw) Option {
	return func(o *dlframework.Feature) {
		RawType()(o)
		o.Feature = &dlframework.Feature_Raw{
			Raw: e,
		}
	}
}
