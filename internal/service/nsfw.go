package service

import (
	"errors"
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"

	"github.com/tensorflow/tensorflow/tensorflow/go/op"

	"github.com/kshedden/gonpy"
)

const (
	TENSOR     = 1
	BASE64JPEG = 2
)

type NsfwService struct {
	Weights []float64
	Predictions  tf.Output
}

func NewNsfwService() *NsfwService {
	return &NsfwService{}
}

func (n *NsfwService) Build(weightsPath string, inputType int) (err error) {
	var (
		input       tf.Output
		inputTensor tf.Output
	)

	r, err := gonpy.NewFileReader(weightsPath)
	if err != nil {
		fmt.Println("NewFileReader err ", err.Error())
		return
	}
	n.Weights, err = r.GetFloat64()
	if err != nil {
		fmt.Printf("GetFloat64 err: %s. data type :%s \n", err.Error(),	r.Dtype)
		return
	}
	fmt.Println("n.Weights,",	n.Weights)
	switch inputType {
	case TENSOR:
		s := op.NewScope()
		input = op.Placeholder(s, tf.Float, op.PlaceholderShape(tf.MakeShape(0, 224, 224, 3)))
		inputTensor = input
	case BASE64JPEG:
		s := op.NewScope()
		input = op.Placeholder(s, tf.String, op.PlaceholderShape(tf.MakeShape(0)))
		// TODO 这里没有传入需要的数据
		inputTensor = op.DecodeBase64(s, tf.Output{})
	default:
		err = errors.New(fmt.Sprintf("Invalid input type %d", inputType))
		return
	}
	n.prediction(inputTensor)
	return
}

func (n *NsfwService) prediction(x tf.Output) {
	s := op.NewScope()
	x = op.Pad(s, x, op.Const(s, [][]int64{{0, 0}, {3, 3}, {3, 3}, {0, 0}}))
	x = n.conv2d("conv_1", x, 64, 7, 2, "valid", false)
	x = n.batchNorm("bn_1", x, false)
	x = op.Relu(s, x)
	x = op.MaxPool3D(s, x, []int64{2}, []int64{3}, "valid")
	x = n.convBlock(0, 0, x, []int64{32, 32, 128}, 3, 1)
	x = n.identityBlock(0, 1, x, []int64{32, 32, 128}, 3)
	x = n.identityBlock(0, 2, x, []int64{32, 32, 128}, 3)

	x = n.convBlock(1, 0, x, []int64{64, 64, 256}, 3, 2)
	x = n.identityBlock(1, 1, x, []int64{64, 64, 256}, 3)
	x = n.identityBlock(1, 2, x, []int64{64, 64, 256}, 3)
	x = n.identityBlock(1, 3, x, []int64{64, 64, 256}, 3)

	x = n.convBlock(2, 0, x, []int64{128, 128, 512}, 3, 2)
	x = n.identityBlock(2, 1, x, []int64{128, 128, 512}, 3)
	x = n.identityBlock(2, 2, x, []int64{128, 128, 512}, 3)
	x = n.identityBlock(2, 3, x, []int64{128, 128, 512}, 3)
	x = n.identityBlock(2, 4, x, []int64{128, 128, 512}, 3)
	x = n.identityBlock(2, 5, x, []int64{128, 128, 512}, 3)

	x = n.convBlock(3, 0, x, []int64{256, 256, 1024}, 3, 2)
	x = n.identityBlock(3, 1, x, []int64{256, 256, 1024}, 3)
	x = n.identityBlock(3, 2, x, []int64{256, 256, 1024}, 3)

	x = op.AvgPool3D(s, x, []int64{7}, []int64{1}, "valid")
	x = op.Reshape(s, x, op.Const(s, []int64{-1, 1024}))

	logits := n.fullyConnected("fc_nsfw", x, 2)
	n.Predictions = op.Softmax(s, logits)
}

func (n *NsfwService) getWeights(layerName, fieldName string) (x tf.Output) {

	return
}

func (n *NsfwService) fullyConnected(name string, inputs tf.Output, numOutputs int) (x tf.Output) {
	return
}

func (n *NsfwService) conv2d(name string, inputs tf.Output, filterDepth, kernelSize, stride int, padding string, trainable bool) (x tf.Output) {
	return
}

func (n *NsfwService) batchNorm(name string, inputs tf.Output, training bool) (x tf.Output) {
	return
}

func (n *NsfwService) convBlock(stage, block int, inputs tf.Output, filterDepths []int64, kernelSize, stride int) (x tf.Output) {
	return
}

func (n *NsfwService) identityBlock(stage, block int, inputs tf.Output, filterDepths []int64, kernelSize int) (x tf.Output) {
	return
}
