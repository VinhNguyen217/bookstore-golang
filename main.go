package main

import (
	"book-store/cmd"
	"book-store/log"
	"context"
	"fmt"
	"os"
)

func init() {
}

func main() {
	pid := os.Getpid()
	log.Infow(context.Background(), fmt.Sprintf("Process ID: %v", pid))
	cmd.Execute()
}
