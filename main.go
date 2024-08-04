package main

import (
	"selector"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("selector", func(w *unison.Window) {
		w.Content().AddChild(selector.New().Layout())
	})
}
