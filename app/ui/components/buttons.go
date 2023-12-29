package components

import (
	"github.com/airtonix/pixelclicker/app/ui/assets"
	"github.com/ebitenui/ebitenui/widget"
)

var primaryButtonTxtColors = &widget.ButtonTextColor{
	Idle:     assets.Colors.PurpleLight,
	Disabled: assets.Colors.PurpleDark,
}

var primaryButtonBgColors = &widget.ButtonImage{
	Idle:         assets.Images.PurpleDark,
	Hover:        assets.Images.PurpleLight,
	Pressed:      assets.Images.Purple,
	Disabled:     assets.Images.PurpleDark,
	PressedHover: assets.Images.PurpleLight,
}

var primaryButtonPadding = widget.Insets{
	Left:   10,
	Right:  10,
	Top:    5,
	Bottom: 5,
}

func NewPrimaryButton(
	// button label
	label string,
	onClick widget.ButtonClickedHandlerFunc,
	layout interface{},
) *widget.Button {
	face, _ := assets.LoadFont(12)

	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(layout),
		),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text(label, face, primaryButtonTxtColors),

		// speficy the button's background image and its colors
		widget.ButtonOpts.Image(primaryButtonBgColors),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(primaryButtonPadding),

		// specify the button's clicked handler
		widget.ButtonOpts.ClickedHandler(onClick),
	)

	return button
}
