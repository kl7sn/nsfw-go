package main

import (
	"fmt"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
	"github.com/kl7sn/nsfw-go/internal/apiv1"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {

	checkEnv()

	if err := ego.New().Serve(func() *egin.Component {
		server := egin.Load("server.http").Build()
		server.GET("/hello", apiv1.Nsfw)
		return server
	}()).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}


func checkEnv(){
	// Construct a graph with an operation that produces a string constant.
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version " + tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output[0].Value())
	return
}