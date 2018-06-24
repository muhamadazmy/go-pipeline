[![GoDoc](https://godoc.org/github.com/muhamadazmy/go-pipeline?status.svg)](https://godoc.org/github.com/muhamadazmy/go-pipeline)

# Introduction
Pipeline is a very thin wrapper around "golang.org/x/sync/errgroup" which allows clean building
of pipelines where each stage feed in data to the next stage.