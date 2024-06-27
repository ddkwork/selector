package main

import (
	"selector"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("selector", func(w *unison.Window) {
		w.Content().AddChild(selector.New().Layout())
	})
}
