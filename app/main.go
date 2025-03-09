package main

import (
	"github.com/DevAthhh/caching-proxy/app/cmd"
	"github.com/DevAthhh/caching-proxy/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToRedis()
}

func main() {
	cmd.RootCommand().Execute()
}
