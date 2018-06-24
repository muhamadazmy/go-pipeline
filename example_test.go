package pipeline_test

import (
	"fmt"
	"math"

	"github.com/muhamadazmy/go-pipeline"
)

func ExamplePipeline() {
	generate := func(s pipeline.Stage, f, t int) <-chan int {
		out := make(chan int)
		s.Go(func() error {
			defer close(out)

			for i := f; i < t; i++ {
				select {
				case out <- i:
				case <-s.Done():
					return nil
				}
			}
			return nil
		})

		return out
	}

	square := func(s pipeline.Stage, in <-chan int) <-chan int {
		out := make(chan int)
		s.Go(func() error {
			defer close(out)
			for x := range in {
				select {
				case out <- int(math.Pow(float64(x), 2)):
				case <-s.Done():
					return nil
				}
			}
			return nil
		})

		return out
	}

	pipe := pipeline.New()

	stream := generate(pipe.Stage(), 1, 5)
	stream = square(pipe.Stage(), stream)

	go pipe.Wait()

	for result := range stream {
		fmt.Println(result)
	}

	if err := pipe.Wait(); err != nil {
		panic(err)
	}

	//Output:
	//1
	//4
	//9
	//16
}
