package feature

import "github.com/c3sr/dlframework"

func GeoLocationType() Option {
	return Type(dlframework.FeatureType_GEOLOCATION)
}

func GeoLocation(e *dlframework.GeoLocation) Option {
	return func(o *dlframework.Feature) {
		GeoLocationType()(o)
		o.Feature = &dlframework.Feature_Geolocation{
			Geolocation: e,
		}
	}
}
