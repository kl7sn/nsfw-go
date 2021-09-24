package service

import (
	"testing"

	tensorflow "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TestNsfwService_Build(t *testing.T) {
	type fields struct {
		Weights     []float64
		Predictions tensorflow.Output
	}
	type args struct {
		weightsPath string
		inputType   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test-1",
			fields:  fields{},
			args:    args{
				weightsPath: "/Users/duminxiang/cosmos/go/src/github.com/kl7sn/nsfw-go/pkg/data/open_nsfw-weights.npy",
				inputType:   TENSOR,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NsfwService{
				Weights:     make([]float64, 0),
				Predictions: tt.fields.Predictions,
			}
			if err := n.Build(tt.args.weightsPath, tt.args.inputType); (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
