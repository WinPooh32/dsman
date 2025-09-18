package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/WinPooh32/dsman/ui"
)

const (
	appTitle            = "dsman"
	defaultWindowWidth  = 640
	defaultWindowHeight = 400
	defaultTheme        = "light"
)

func parseMainWindowSettings() (ui.MainWindowSettings, error) {
	w := flag.Int("w", defaultWindowWidth, "window width")
	h := flag.Int("h", defaultWindowHeight, "window height")
	theme := flag.String("theme", defaultTheme, "global window theme style, possible values: light, dark")
	flag.Parse()

	settings := ui.MainWindowSettings{
		Title:  appTitle,
		Width:  *w,
		Height: *h,
		Theme:  *theme,
	}
	if err := settings.Validate(); err != nil {
		return ui.MainWindowSettings{}, fmt.Errorf("parse settings: %s", err)
	}

	return settings, nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	settings, err := parseMainWindowSettings()
	if err != nil {
		slog.Error(err.Error())
		flag.PrintDefaults()

		os.Exit(1)
	}

	window := ui.NewMainWindow(settings)

	window.Run(ctx)
}
