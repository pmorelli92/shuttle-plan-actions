package main

import (
	"context"

	_ "github.com/lunarway/shuttle/pkg/executors/golang/cmder"
)

func Build(ctx context.Context) error {
	println("build")
	return nil
}
