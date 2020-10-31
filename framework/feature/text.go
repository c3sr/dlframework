package feature

import "github.com/c3sr/dlframework"

func TextType() Option {
	return Type(dlframework.FeatureType_TEXT)
}

func Text(e *dlframework.Text) Option {
	return func(o *dlframework.Feature) {
		TextType()(o)
		o.Feature = &dlframework.Feature_Text{
			Text: e,
		}
	}
}
