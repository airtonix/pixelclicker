package components

import (
	"github.com/airtonix/pixelclicker/app/ui/assets"
	"github.com/ebitenui/ebitenui/widget"
)

func NewTabBook(
	defaultTab *widget.TabBookTab,
	otherTabs ...*widget.TabBookTab,
) *widget.TabBook {
	face10, _ := assets.LoadFont(10)
	tabs := []*widget.TabBookTab{}
	tabs = append(tabs, defaultTab)
	for _, tab := range otherTabs {
		tabs = append(tabs, tab)
	}

	TabBook := widget.NewTabBook(
		widget.TabBookOpts.Tabs(
			tabs...,
		),

		widget.TabBookOpts.InitialTab(defaultTab),
		widget.TabBookOpts.TabButtonImage(
			&widget.ButtonImage{
				Idle:    assets.Images.PurpleDarker,
				Pressed: assets.Images.PurpleLight,
			},
		),
		widget.TabBookOpts.TabButtonText(face10, &widget.ButtonTextColor{
			Idle:     assets.Colors.PurpleLightest,
			Disabled: assets.Colors.Purple,
		}),
		widget.TabBookOpts.TabButtonSpacing(0),
		widget.TabBookOpts.TabButtonOpts(
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.MinSize(98, 0)),
		),
		widget.TabBookOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(
					widget.AnchorLayoutData{
						StretchHorizontal:  true,
						StretchVertical:    true,
						HorizontalPosition: widget.AnchorLayoutPositionEnd,
						VerticalPosition:   widget.AnchorLayoutPositionCenter,
					},
				),
			),
		),
	)

	return TabBook
}
