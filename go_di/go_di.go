package go_di

import (
	"github.com/squeedee/go_di/widget"
)

func GoDi() string {
	widgetInstance := widget.NewWidget("aceness")
	return widgetInstance.Run()
}
