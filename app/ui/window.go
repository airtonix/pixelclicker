package ui

import (
	"fmt"
	"image/color"

	"github.com/airtonix/pixelclicker/app/core"
	"github.com/airtonix/pixelclicker/app/core/terminalcolors"
	"github.com/airtonix/pixelclicker/app/scanning"
	"github.com/airtonix/pixelclicker/app/ui/assets"
	"github.com/airtonix/pixelclicker/app/ui/components"
	"github.com/ebitenui/ebitenui"
	"github.com/muesli/gamut"

	ebitenuiImage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X float32
	Y float32
}

type Dimensions struct {
	Width  float32
	Height float32
}

type Visibility struct {
	Visible bool
}

type ScanLens struct {
	Size     float32
	Position Position
	Center   Position
	Dimensions
	Visibility
}

type Tabs map[string]*widget.TabBookTab
type TabBook struct {
	tabs   Tabs
	active string
}

// a Window is a rectangular area on the screen
// that can be moved and resized.
// It is used to define the area of the screen
// that the scanner will scan.
type Window struct {
	Position
	Dimensions
	lens     ScanLens
	tabbook  TabBook
	ui       *ebitenui.UI
	snaphots *core.Stack[*ebiten.Image]
}

func NewWindow() *Window {

	x, y := float32(0), float32(0)
	// width/height comes from the ebiten window
	width, height := ebiten.WindowSize()

	window := &Window{
		Position: Position{
			X: x,
			Y: y,
		},
		Dimensions: Dimensions{
			Width:  float32(width),
			Height: float32(height),
		},
		snaphots: core.NewStack[*ebiten.Image](2),
		lens: ScanLens{
			Visibility: Visibility{
				Visible: true,
			},
			Size: 10,
		},
	}

	SetingsTab := window.CreateSettingsTab()
	ScannerTab := window.CreateScannerTab()

	window.tabbook = TabBook{
		tabs: Tabs{
			"Scanner":  ScannerTab,
			"Settings": SetingsTab,
		},
	}

	TabBook := components.NewTabBook(
		ScannerTab,
		SetingsTab,
	)

	PageContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(
					widget.NewInsetsSimple(10),
				),
			),
		),
	)

	PageContainer.AddChild(TabBook)

	window.ui = &ebitenui.UI{
		Container: PageContainer,
	}

	return window
}

func (w *Window) Update() {
	w.ui.Update()

	width, height := ebiten.WindowSize()

	w.Width = float32(width)
	w.Height = float32(height)

	w.lens.Width = float32(w.lens.Size)
	w.lens.Height = float32(w.lens.Size)

	bounds := w.tabbook.tabs["Scanner"].Children()[0].GetWidget().Rect.Bounds()

	w.lens.Center = Position{
		X: float32(bounds.Min.X + bounds.Dx()/2),
		Y: float32(bounds.Min.Y + bounds.Dy()/2),
	}

	w.lens.Position.X = w.lens.Center.X - w.lens.Width/2
	w.lens.Position.Y = w.lens.Center.Y - w.lens.Height/2
}

func (w *Window) Draw(screen *ebiten.Image) {
	components.DrawBox(screen, w.lens.Position.X, w.lens.Position.Y, w.lens.Width, w.lens.Height)
	w.ui.Draw(screen)
}

func (w *Window) CreateScannerTab() *widget.TabBookTab {
	TabScanner := widget.NewTabBookTab("Scanner",
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				//Define number of columns in the grid
				widget.GridLayoutOpts.Columns(1),

				//Define how much padding to inset the child content
				// widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(0)),

				//Define how far apart the rows and columns should be
				widget.GridLayoutOpts.Spacing(5, 5),

				//Define how to stretch the rows and columns. Note it is required to
				//specify the Stretch for each row and column.
				widget.GridLayoutOpts.Stretch(
					// columns
					[]bool{
						true,
					},
					// rows
					[]bool{
						true,
						false,
					},
				),
			),
		),
	)

	// the lens spacer
	lensContainer := widget.NewContainer(

		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),

		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
	)

	lensContainer.AddChild(
		widget.NewContainer(),
	)

	TabScanner.AddChild(lensContainer)

	controlsContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ebitenuiImage.NewNineSliceColor(color.NRGBA{R: 80, G: 80, B: 140, A: 255})),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(

				//Define number of columns in the grid
				widget.GridLayoutOpts.Columns(6),

				//Define how much padding to inset the child content
				widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(5)),

				//Define how far apart the rows and columns should be
				widget.GridLayoutOpts.Spacing(5, 5),

				//Define how to stretch the rows and columns. Note it is required to
				//specify the Stretch for each row and column.
				widget.GridLayoutOpts.Stretch(
					// columns
					[]bool{
						false,
						false,
						false,
						false,
						false,
						false,
					},
					// rows
					[]bool{
						true,
					},
				),
			),
		),
	)

	controlsContainer.AddChild(
		components.NewPrimaryButton(
			"Capture",
			func(args *widget.ButtonClickedEventArgs) {
				w.TakeScreenshot()
			},
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		),
	)

	controlsContainer.AddChild(
		components.NewSlider(func(args *widget.SliderChangedEventArgs) {
			value := args.Current
			if value == 0 {
				value = 1
			}

			w.lens.Size = float32(value)
		}),
	)

	controlsContainer.AddChild(
		components.NewPrimaryButton(
			"Start",
			func(args *widget.ButtonClickedEventArgs) {
				windowX, windowY := ebiten.WindowPosition()

				x, y, w, h :=
					int(w.lens.Position.X),
					int(w.lens.Position.Y),
					int(w.lens.Width),
					int(w.lens.Height)

				fmt.Println("Screenshot", x+windowX, y+windowY, w, h)

			},
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		),
	)
	controlsContainer.AddChild(
		components.NewPrimaryButton(
			"Stop",
			func(args *widget.ButtonClickedEventArgs) {
				windowX, windowY := ebiten.WindowPosition()

				x, y, w, h :=
					int(w.lens.Position.X),
					int(w.lens.Position.Y),
					int(w.lens.Width),
					int(w.lens.Height)

				fmt.Println("Screenshot", x+windowX, y+windowY, w, h)

			},
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			},
		),
	)

	TabScanner.AddChild(controlsContainer)

	return TabScanner
}

func (w *Window) CreateSettingsTab() *widget.TabBookTab {
	face10, _ := assets.LoadFont(10)

	SettingsTab := widget.NewTabBookTab("Settings",
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(1),

			//Define how much padding to inset the child content
			// widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(0)),

			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(5, 5),

			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch(
				// columns
				[]bool{
					true,
				},
				// rows
				[]bool{
					false,
					true,
				},
			),
		)),
	)

	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(ebitenuiImage.NewNineSliceColor(color.NRGBA{R: 80, G: 80, B: 140, A: 255})),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(widget.Insets{
					Top:    5,
					Left:   5,
					Right:  5,
					Bottom: 5,
				}),
			),
		),
	)

	SettingsTab.AddChild(sliderContainer)
	SettingsTab.AddChild(widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(400, 200),
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position:  widget.RowLayoutPositionCenter,
				MaxHeight: 100,
				Stretch:   true,
			}),
		),
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     ebitenuiImage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			Disabled: ebitenuiImage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
		}),
		widget.TextInputOpts.Face(face10),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          color.NRGBA{254, 255, 255, 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{254, 255, 255, 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		}),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(face10, 2),
		),
		widget.TextInputOpts.Placeholder("Enter keystrokes to replay"),

		//This method is called whenever there is a text change.
		//It allows the developer to allow or deny a change.
		//In this case we are forcing the string to be all caps.
		//The first return parameter is whether or not to accept the text as is.
		//The second return parameter is what to replace the text with if it is not accepted (optional)
		widget.TextInputOpts.Validation(func(newInputText string) (bool, *string) {
			// newInputText = strings.ToUpper(newInputText)
			return false, &newInputText
		}),
		widget.TextInputOpts.AllowDuplicateSubmit(true),
		widget.TextInputOpts.IgnoreEmptySubmit(true),
		widget.TextInputOpts.SubmitHandler(func(args *widget.TextInputChangedEventArgs) {
			fmt.Println("Text Submitted: ", args.InputText)
			args.TextInput.SetText(
				fmt.Sprintf("%s\n", args.InputText),
			)
		}),
		widget.TextInputOpts.ChangedHandler(func(args *widget.TextInputChangedEventArgs) {
			fmt.Println("Text Changed: ", args.InputText)
		}),
	))

	return SettingsTab
}

func (w *Window) TakeScreenshot() {
	w.lens.Visible = false
	windowX, windowY := ebiten.WindowPosition()
	screen := ebiten.Monitor().Name()

	x := int(w.lens.Position.X) + windowX
	y := int(w.lens.Position.Y) + windowY
	width := int(w.lens.Width)
	height := int(w.lens.Height)

	pixels := scanning.GetPixels(x, y, width, height)

	fmt.Printf("Screenshot: %s %d pixels at %dx%d\n", screen, width*height, x, y)
	fmt.Printf("pixels: \n")
	// each row
	for i := 0; i < len(pixels); i++ {
		// each column
		for j := 0; j < len(pixels[i]); j++ {
			cell := pixels[i][j]
			// convert hex to rgb
			color := gamut.Hex(fmt.Sprintf("#%x", cell))
			r, g, b, _ := color.RGBA()
			terminalcolors.Background(uint8(r), uint8(g), uint8(b)).Print(" ")
		}
		fmt.Printf("\n")
	}
	w.lens.Visible = true

}
