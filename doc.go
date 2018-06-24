/*
Pipeline is a very thin wrapper around "golang.org/x/sync/errgroup" which allows clean building
of pipelines where each stage feed in data to the next stage.

Working with piplines

	pipe := pipeline.New()

	stream := stage1(pipe.Stage(), args...)
	stream = stage2(pipe.Stage(), stream)
	//add more stages if needed

	result := stageLast(stream)
	err := pipe.Wait()

Another approach is to wait in a go routine, and then do the consumption for the stream in the same function context
	go func() {
		pipe.Wait()
	}

	for obj := range stream {
		//process object
	}

	err := pipe.Wait()
	//inspect if we had an error
*/
package pipeline
