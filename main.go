package main

import (
	"fmt"
	"log"
	"os"

	"github.com/airtonix/pixelclicker/app/scanning"
	"github.com/airtonix/pixelclicker/app/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputMode int

const (
	ModeNone InputMode = iota
	ModeKey
	ModePattern
	ModeScan
)

type App struct {
	Window     *ui.Window
	Scanner    *scanning.Scanner
	KeyToPress string
	InputMode  InputMode
}

func NewApp() *App {
	return &App{
		Window:  ui.NewWindow(),
		Scanner: scanning.NewScanner(),
	}
}

var pressedKeys []ebiten.Key

func (a *App) Update() error {
	a.Window.Update()

	pressedKeys = inpututil.AppendJustPressedKeys(pressedKeys[:0])

	if a.InputMode == ModeKey && len(pressedKeys) > 0 {
		a.InputMode = ModeNone
		a.KeyToPress = fmt.Sprintf("%v", pressedKeys[0])
		return nil
	}

	if a.InputMode == ModePattern {
		a.Scanner.CapturePattern()
		a.InputMode = ModeNone
		return nil

	}

	if a.InputMode == ModeScan && ebiten.IsKeyPressed(ebiten.KeyE) {
		a.InputMode = ModeNone
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		// records the pattern
		a.InputMode = ModePattern
		return nil

	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		// records the key to press
		a.InputMode = ModeKey
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		a.InputMode = ModeScan
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) {
		os.Exit(0)
		return nil
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Window.Draw(screen)
}

func (a *App) DrawHelp(screen *ebiten.Image) {

	helpArray := []string{
		fmt.Sprintf("[Mode: %d]", a.InputMode),
		"---",
	}

	if a.Scanner.HasPattern() {
		helpArray = append(helpArray,
			fmt.Sprintf("Pattern: %s", a.Scanner.GetPatternSize()),
			"[P] set new pattern",
		)
	} else {
		helpArray = append(helpArray,
			"[P] capture pattern",
		)
	}

	if a.InputMode == ModeScan || a.InputMode == ModeNone {
		helpArray = append(helpArray,
			"[X] exit",
		)
	}

	if a.InputMode == ModeScan {
		helpArray = append(helpArray,
			"[E] end scan",
		)
	} else {
		helpArray = append(helpArray,
			"[S] start scan",
		)
	}

	if len(a.KeyToPress) > 0 {
		helpArray = append(helpArray,
			"---",
			fmt.Sprintf("[PRESS]: %s", a.KeyToPress),
			"[K] change key",
		)
	} else {
		helpArray = append(helpArray,
			"---",
			"[K] set key to press",
		)
	}

	if a.InputMode == ModeScan {
		for _, line := range a.Scanner.GetHistory().Items() {
			helpArray = append(helpArray,
				"---",
				fmt.Sprintf("History: %v", line.Confidence),
			)
		}
	}

	// join helpArray with newlines
	help := ""
	for _, line := range helpArray {
		help += line + "\n"
	}

	ebitenutil.DebugPrint(screen, help)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowDecorated(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	app := NewApp()

	err := ebiten.RunGameWithOptions(app, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}
