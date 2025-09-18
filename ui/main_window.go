package ui

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"unicode/utf8"

	"github.com/AllenDang/giu"
	g "github.com/AllenDang/giu"
	"github.com/WinPooh32/dsman/ui/themes"
)

type MainWindowSettings struct {
	Title  string
	Width  int
	Height int
	Theme  string
}

func (s *MainWindowSettings) Validate() error {
	var errs []error

	const maxTitleChars = 64
	if runeCount := utf8.RuneCountInString(s.Title); s.Title == "" || runeCount > maxTitleChars {
		errs = append(errs, fmt.Errorf("title length %d is out of allowed range (%d; %d]", runeCount, 0, maxTitleChars))
	}

	const minWindowSize = 1

	if s.Width < 0 {
		errs = append(errs, fmt.Errorf("window width %d is lower than %d", s.Width, minWindowSize))
	}

	if s.Height < 0 {
		errs = append(errs, fmt.Errorf("window height %d is lower than %d", s.Height, minWindowSize))
	}

	switch s.Theme {
	case themes.DarkName:
	case themes.LightName:
	default:
		errs = append(errs, fmt.Errorf("theme %s is unknown", s.Theme))
	}

	return errors.Join(errs...)
}

type MainWindow struct {
	wnd      *g.MasterWindow
	settings MainWindowSettings

	closeC    chan struct{}
	closeOnce sync.Once
}

func NewMainWindow(settings MainWindowSettings) *MainWindow {
	mw := &MainWindow{
		wnd:       nil,
		settings:  settings,
		closeC:    make(chan struct{}),
		closeOnce: sync.Once{},
	}

	mw.wnd = g.NewMasterWindow("dsman", settings.Width, settings.Height, 0)

	switch settings.Theme {
	case themes.LightName:
		mw.wnd.SetStyle(themes.Light())
	case themes.DarkName:
	}

	mw.wnd.SetCloseCallback(func() bool {
		mw.Close()
		return true
	})

	return mw
}

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func (mw *MainWindow) Run(ctx context.Context) {
	mw.wnd.Run(func() {
		exit := false

		select {
		case <-mw.closeC:
			exit = true
		case <-ctx.Done():
			exit = true
		default:
		}

		if exit {
			mw.wnd.SetShouldClose(true)
			giu.Update()

			return
		}

		g.SingleWindow().Layout(
			g.Label("Hello world from giu"),
			g.Row(
				g.Button("Click Me").OnClick(onClickMe),
				g.Button("I'm so cute").OnClick(onImSoCute),
			),
		)
	})
}

func (mw *MainWindow) Close() {
	mw.closeOnce.Do(func() {
		close(mw.closeC)
	})
}
