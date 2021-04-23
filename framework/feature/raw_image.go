package feature

import (
	"github.com/c3sr/dlframework"
)

func RawImageType() Option {
	return Type(dlframework.FeatureType_RAW_IMAGE)
}

func RawImage(e *dlframework.RawImage) Option {
	return func(o *dlframework.Feature) {
		RawImageType()(o)
		o.Feature = &dlframework.Feature_RawImage{
			RawImage: e,
		}
	}
}

func ensureRawImage(o *dlframework.Feature) *dlframework.RawImage {
	if o.Type != dlframework.FeatureType_RAW_IMAGE && !isUnknownType(o) {
		panic("unexpected feature type")
	}
	if o.Feature == nil {
		o.Feature = &dlframework.Feature_RawImage{}
	}
	img, ok := o.Feature.(*dlframework.Feature_RawImage)
	if !ok {
		panic("expecting an raw image feature")
	}
	if img.RawImage == nil {
		img.RawImage = &dlframework.RawImage{}
	}
	RawImageType()(o)
	return img.RawImage
}

func RawImageID(id string) Option {
	return func(o *dlframework.Feature) {
		img := ensureRawImage(o)
		img.ID = id
	}
}

func RawImageWidth(width int) Option {
	return func(o *dlframework.Feature) {
		img := ensureRawImage(o)
		img.Width = int32(width)
	}
}

func RawImageHeight(height int) Option {
	return func(o *dlframework.Feature) {
		img := ensureRawImage(o)
		img.Height = int32(height)
	}
}

func RawImageChannels(channels int) Option {
	return func(o *dlframework.Feature) {
		img := ensureRawImage(o)
		img.Channels = int32(channels)
	}
}

func RawImageFloat32Data(data []float32) Option {
	return func(o *dlframework.Feature) {
		img := ensureRawImage(o)
		img.FloatList = data
		img.DataType = "float32"
	}
}

func RawImageInt8Data(data []int8) Option {
	return func(o *dlframework.Feature) {
		buf := make([]int32, len(data))
		for ii, val := range data {
			buf[ii] = int32(val)
		}
		RawImageInt32Data(buf)(o)
		img := ensureRawImage(o)
		img.DataType = "int8"
	}
}

func RawImageUInt8Data(data []uint8) Option {
	return func(o *dlframework.Feature) {
		buf := make([]int32, len(data))
		for ii, val := range data {
			buf[ii] = int32(val)
		}
		RawImageInt32Data(buf)(o)

		img := ensureRawImage(o)
		img.DataType = "uint8"
	}
}

func RawImageInt32Data(data []int32) Option {
	return func(o *dlframework.Feature) {
		img := ensureRawImage(o)
		img.DataType = "int32"
		img.CharList = data
	}
}

func RawImageData(data interface{}) Option {
	switch v := data.(type) {
	case []int8:
		return RawImageInt8Data(v)
	case []uint8:
		return RawImageUInt8Data(v)
	case []int32:
		return RawImageInt32Data(v)
	case []float32:
		return RawImageFloat32Data(v)
	}
	panic("invalid RawImageData type")
}

func CreateRawImageFeaturesCanonical(images [][][][]float32) []dlframework.Features {
	features := make([]dlframework.Features, len(images))

	for i, _ := range features {
		cur := images[i]
		height := len(cur)
		width := len(cur[0])
		channels := len(cur[0][0])

		pixels := make([]float32, height*width*channels)

		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				for c := 0; c < channels; c++ {
					pixels[h*width*channels+w*channels+c] = cur[h][w][c]
				}
			}
		}

		features[i] = dlframework.Features{
			New(
				RawImageType(),
				RawImageWidth(width),
				RawImageHeight(height),
				RawImageChannels(channels),
				RawImageData(pixels),
			),
		}
	}

	return features
}
