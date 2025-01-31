package main

import (
	"os/signal"
	"syscall"

	"github.com/nerdwave-nick/nerdlocke/cmd"
	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	cmd.Execute(ctx)
}
