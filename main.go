package main

import (
	_ "rme/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"rme/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
