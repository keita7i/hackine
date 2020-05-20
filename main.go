package main

import (
	"github.com/keitam913/hackine/di"
)

func main() {
	dc := di.Container{}
	dc.Router().Run(":80")
}
