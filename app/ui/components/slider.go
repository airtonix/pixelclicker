package components

import (
	"image/color"

	ebitenuiImage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func NewSlider(
	onChange func(args *widget.SliderChangedEventArgs),
) *widget.Slider {

	// construct a slider
	slider := widget.NewSlider(
		// Set the slider orientation - n/s vs e/w
		widget.SliderOpts.Direction(widget.DirectionHorizontal),

		// Set the minimum and maximum value for the slider
		widget.SliderOpts.MinMax(0, 100),

		widget.SliderOpts.WidgetOpts(
			// Set the Widget to layout in the center on the screen
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
			// Set the widget's dimensions
			widget.WidgetOpts.MinSize(200, 20),
		),

		widget.SliderOpts.Images(
			// Set the track images
			&widget.SliderTrackImage{
				Idle:  ebitenuiImage.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: ebitenuiImage.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			// Set the handle images
			&widget.ButtonImage{
				Idle:    ebitenuiImage.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Hover:   ebitenuiImage.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
				Pressed: ebitenuiImage.NewNineSliceColor(color.NRGBA{255, 100, 100, 255}),
			},
		),

		// Set the size of the handle
		widget.SliderOpts.FixedHandleSize(6),

		// Set the offset to display the track
		widget.SliderOpts.TrackOffset(0),

		// Set the size to move the handle
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),

		// Set the callback to call when the slider value is changed
		widget.SliderOpts.ChangedHandler(onChange),
	)
	// Set the current value of the slider
	slider.Current = 5
	return slider
}
