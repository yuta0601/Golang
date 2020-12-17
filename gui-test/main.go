package main

import (
    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/layout"
    "fyne.io/fyne/widget"
)

func main() {
    ap := app.New()
    w := ap.NewWindow("Layout Sample")
    w.CenterOnScreen()
    lbl := widget.NewLabelWithStyle(
		"0",
		fyne.TextAlignTrailing,
		fyne.TextStyle{Bold: true},
	)
    w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1), lbl,
        fyne.NewContainerWithLayout(layout.NewGridLayout(4),
            widget.NewButton("AC", func() {}),
            widget.NewButton("+-", func() {}),
            widget.NewButton("%", func() {}),
            widget.NewButton("÷", func() {}),
        ),
        fyne.NewContainerWithLayout(layout.NewGridLayout(4),
            widget.NewButton("7", func() {}),
            widget.NewButton("8", func() {}),
            widget.NewButton("9", func() {}),
            widget.NewButton("×", func() {}),
        ),
        fyne.NewContainerWithLayout(layout.NewGridLayout(4),
            widget.NewButton("4", func() {}),
            widget.NewButton("5", func() {}),
            widget.NewButton("6", func() {}),
            widget.NewButton("-", func() {}),
        ),
        fyne.NewContainerWithLayout(layout.NewGridLayout(4),
            widget.NewButton("1", func() {}),
            widget.NewButton("2", func() {}),
            widget.NewButton("3", func() {}),
            widget.NewButton("+", func() {}),
        ),
        fyne.NewContainerWithLayout(layout.NewGridLayout(4),
            widget.NewButton("0", func() {}),
            widget.NewButton(".", func() {}),
            widget.NewButton("C", func() {}),
            widget.NewButton("=", func() {}),
        ),
        fyne.NewContainerWithLayout(layout.NewGridLayout(3),
            widget.NewButton("left", func() {}),
            layout.NewSpacer(),
            widget.NewButton("right", func() {}),
        ),
        fyne.NewContainerWithLayout(layout.NewGridLayout(1),
            widget.NewButton("close", func() {
                w.Close()
            }),
        ),
    ))
    w.ShowAndRun()
}