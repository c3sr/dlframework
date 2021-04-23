package feature

import (
	"sort"

	"github.com/c3sr/dlframework"
	"github.com/c3sr/utils"
)

func InstanceSegmentType() Option {
	return Type(dlframework.FeatureType_INSTANCESEGMENT)
}

func InstanceSegment(e *dlframework.InstanceSegment) Option {
	return func(o *dlframework.Feature) {
		InstanceSegmentType()(o)
		o.Feature = &dlframework.Feature_InstanceSegment{
			InstanceSegment: e,
		}
	}
}

func ensureInstanceSegment(o *dlframework.Feature) *dlframework.InstanceSegment {
	if o.Type != dlframework.FeatureType_INSTANCESEGMENT && !isUnknownType(o) {
		panic("unexpected feature type")
	}
	if o.Feature == nil {
		o.Feature = &dlframework.Feature_InstanceSegment{}
	}
	iseg, ok := o.Feature.(*dlframework.Feature_InstanceSegment)
	if !ok {
		panic("expecting a classification feature")
	}
	if iseg.InstanceSegment == nil {
		iseg.InstanceSegment = &dlframework.InstanceSegment{}
	}
	InstanceSegmentType()(o)
	return iseg.InstanceSegment
}

func InstanceSegmentXmin(xmin float32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Xmin = xmin
	}
}

func InstanceSegmentXmax(xmax float32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Xmax = xmax
	}
}

func InstanceSegmentYmin(ymin float32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Ymin = ymin
	}
}

func InstanceSegmentYmax(ymax float32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Ymax = ymax
	}
}
func InstanceSegmentIndex(index int32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Index = index
	}
}

func InstanceSegmentLabel(label string) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Label = label
	}
}

func InstanceSegmentMaskType(masktype string) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.MaskType = masktype
	}
}

func InstanceSegmentHeight(height int32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Height = height
	}
}

func InstanceSegmentWidth(width int32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.Width = width
	}
}

func InstanceSegmentIntMask(mask []int32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.IntMask = mask
	}
}

func InstanceSegmentFloatMask(mask []float32) Option {
	return func(o *dlframework.Feature) {
		iseg := ensureInstanceSegment(o)
		iseg.FloatMask = mask
	}
}

func CreateInstanceSegmentFeaturesCanonical(probabilities [][]float32, classes [][]float32, boxes [][][4]float32, masks [][][][]float32, labels []string) []dlframework.Features {
	features := make([]dlframework.Features, len(probabilities))

	for i, _ := range features {
		featureLen := len(probabilities[i])
		rprobs := make([]*dlframework.Feature, featureLen)

		for j := 0; j < featureLen; j++ {
			mask := masks[i][j]
			masktype := "float"
			height := len(mask)
			width := len(mask[0])
			rprobs[j] = New(
				InstanceSegmentType(),
				InstanceSegmentXmin(boxes[i][j][1]),
				InstanceSegmentXmax(boxes[i][j][3]),
				InstanceSegmentYmin(boxes[i][j][0]),
				InstanceSegmentYmax(boxes[i][j][2]),
				InstanceSegmentIndex(int32(classes[i][j])),
				InstanceSegmentLabel(labels[int32(classes[i][j])]),
				InstanceSegmentMaskType(masktype),
				InstanceSegmentFloatMask(utils.FlattenFloat32Slice(mask)),
				InstanceSegmentHeight(int32(height)),
				InstanceSegmentWidth(int32(width)),
				Probability(probabilities[i][j]),
			)
		}

		sort.Sort(dlframework.Features(rprobs))
		features[i] = rprobs
	}

	return features
}
