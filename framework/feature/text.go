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

func ensureText(o *dlframework.Feature) *dlframework.Text {
	if o.Type != dlframework.FeatureType_TEXT && !isUnknownType(o) {
		panic("unexpected feature type")
	}
	if o.Feature == nil {
		o.Feature = &dlframework.Feature_Text{}
	}
	text, ok := o.Feature.(*dlframework.Feature_Text)
	if !ok {
		panic("expecting a text feature")
	}
	if text.Text == nil {
		text.Text = &dlframework.Text{}
	}
	TextType()(o)
	return text.Text
}

func TextData(data []byte) Option {
	return func(o *dlframework.Feature) {
		text := ensureText(o)
		text.Data = data
	}
}

func CreateTextFeaturesCanonical(feat [][]byte) []dlframework.Features {
	features := make([]dlframework.Features, len(feat))

	for i, _ := range features {
		cur := feat[i]

		features[i] = dlframework.Features{
			New(
				TextType(),
				TextData(cur),
			),
		}
	}

	return features
}
