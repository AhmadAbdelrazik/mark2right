package main

import (
	"AhmadAbdelrazik/mark2right/internal/app"
)

func main() {
	// renders := Renderers()
	// checker := app.NewChecker()
}

func Renderers() []app.IRender {
	var renders []app.IRender

	hr := app.NewHeaderRenderer()
	renders = append(renders, hr)
	fr := app.NewFontRenderer()
	renders = append(renders, fr)
	lr := app.NewListRenderer()
	renders = append(renders, lr)
	cr := app.NewCodeRenderer()
	renders = append(renders, cr)
	llr := app.NewLinkRenderer()
	renders = append(renders, llr)

	return renders
}
