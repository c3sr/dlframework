package dlframework

import (
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

const exampleModel = `name: Inception
framework:
  name: TensorFlow
  version: ~1.1.x
version: 3.0
container:
  amd64:
    gpu: raiproject/carml-tensorflow:amd64-cpu
    cpu: raiproject/carml-tensorflow:amd64-gpu
  ppc64le:
    cpu: raiproject/carml-tensorflow:ppc64le-gpu
    gpu: raiproject/carml-tensorflow:ppc64le-gpu
description: >
  An image-classification convolutional network.
  Inception achieves 78.0% top-1 and 93.9% top-5 accuracy on the ILSVRC 2012 validation dataset.
  It consists of fewer than 25M parameters.
references:
  - https://arxiv.org/pdf/1512.00567.pdf
license: TODO
inputs:
  - type: image
    description: the input image
    parameters:
      dimensions: [1, 3, 224, 224]
      mean: [117, 117, 117]
output:
  type: feature
  description: an output image net label
  parameters:
    features_url: http://data.dmlc.ml/mxnet/models/imagenet/synset.txt
model:
  base_url: https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip
  graph_path: tensorflow_inception_graph.pb
  is_archive: true
attributes:
  manifest_author: abduld
  trailing_dataset: ImageNet
`

func TestModelMarshaling(t *testing.T) {
	var model ModelManifest
	err := yaml.Unmarshal([]byte(exampleModel), &model)
	assert.NoError(t, err)
	bts, err := model.Marshal()
	assert.NoError(t, err)
	assert.NotEmpty(t, bts)

	unmarshaledModel := ModelManifest{}
	err = unmarshaledModel.Unmarshal(bts)
	assert.NoError(t, err)
	assert.NotEmpty(t, unmarshaledModel)
	assert.True(t, model.Equal(unmarshaledModel))
}
