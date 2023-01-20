package main

import (
	"demo02/internal/cmd"
	_ "demo02/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "demo02/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
