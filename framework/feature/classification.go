package feature

import (
	"sort"

	"github.com/c3sr/dlframework"
)

func ClassificationType() Option {
	return Type(dlframework.FeatureType_CLASSIFICATION)
}

func Classification(e *dlframework.Classification) Option {
	return func(o *dlframework.Feature) {
		ClassificationType()(o)
		o.Feature = &dlframework.Feature_Classification{
			Classification: e,
		}
	}
}

func ensureClassification(o *dlframework.Feature) *dlframework.Classification {
	if o.Type != dlframework.FeatureType_CLASSIFICATION && !isUnknownType(o) {
		panic("unexpected feature type")
	}
	if o.Feature == nil {
		o.Feature = &dlframework.Feature_Classification{}
	}
	cl, ok := o.Feature.(*dlframework.Feature_Classification)
	if !ok {
		panic("expecting a classification feature")
	}
	if cl.Classification == nil {
		cl.Classification = &dlframework.Classification{}
	}
	ClassificationType()(o)
	return cl.Classification
}

func ClassificationIndex(index int32) Option {
	return func(o *dlframework.Feature) {
		cls := ensureClassification(o)
		cls.Index = index
	}
}

func ClassificationLabel(label string) Option {
	return func(o *dlframework.Feature) {
		cls := ensureClassification(o)
		cls.Label = label
	}
}

func CreateClassificationFeaturesCanonical(probabilities [][]float32, labels []string) []dlframework.Features {
	features := make([]dlframework.Features, len(probabilities))

	for i, _ := range features {
		featureLen := len(probabilities[i])

		rprobs := make([]*dlframework.Feature, featureLen)
		for j := 0; j < featureLen; j++ {
			rprobs[j] = New(
				ClassificationIndex(int32(j)),
				ClassificationLabel(labels[j]),
				Probability(probabilities[i][j]),
			)
		}

		res := dlframework.Features(rprobs)
		sort.Sort(res)

		features[i] = res
	}

	return features
}
