package main

import (
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab3/internal/controller"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab3/internal/model"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab3/internal/view"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	w := app.NewWindow(
		app.Title("Chat Client"),
		app.Size(unit.Dp(400), unit.Dp(600)),
	)

	th := material.NewTheme()

	md := model.NewModel()
	vw := view.NewView(th)
	contr := controller.NewController(w, md, vw)

	go func() {
		if err := contr.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
