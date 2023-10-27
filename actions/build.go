package main

import (
	"context"

	_ "github.com/lunarway/shuttle"
)

func Build(ctx context.Context) error {
	println("build")
	return nil
}
