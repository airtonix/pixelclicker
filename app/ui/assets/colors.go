package assets

import (
	"image/color"

	ebitenuiImage "github.com/ebitenui/ebitenui/image"
	"github.com/muesli/gamut"
)

type ColorPalette struct {
	PurpleDarkest  *color.NRGBA
	PurpleDarker   *color.NRGBA
	PurpleDark     *color.NRGBA
	Purple         *color.NRGBA
	PurpleLight    *color.NRGBA
	PurpleLighter  *color.NRGBA
	PurpleLightest *color.NRGBA
	White          *color.NRGBA
	Black          *color.NRGBA
	Silver         *color.NRGBA
}

type ColorImages struct {
	PurpleDarkest  *ebitenuiImage.NineSlice
	PurpleDarker   *ebitenuiImage.NineSlice
	PurpleDark     *ebitenuiImage.NineSlice
	Purple         *ebitenuiImage.NineSlice
	PurpleLight    *ebitenuiImage.NineSlice
	PurpleLighter  *ebitenuiImage.NineSlice
	PurpleLightest *ebitenuiImage.NineSlice
}

var Colors = createColorPallete()
var Images = createColorImages()

func rgbaToNrgba(r, g, b, a uint32) *color.NRGBA {
	return &color.NRGBA{
		uint8(r),
		uint8(g),
		uint8(b),
		uint8(a),
	}
}

func createColorPallete() ColorPalette {
	purple := gamut.Hex("#6000ff")
	colors := ColorPalette{
		White:          rgbaToNrgba(255, 255, 255, 255),
		Black:          rgbaToNrgba(0, 0, 0, 255),
		Silver:         rgbaToNrgba(192, 192, 192, 255),
		PurpleLightest: rgbaToNrgba(gamut.Lighter(purple, 0.6).RGBA()),
		PurpleLighter:  rgbaToNrgba(gamut.Lighter(purple, 0.4).RGBA()),
		PurpleLight:    rgbaToNrgba(gamut.Lighter(purple, 0.2).RGBA()),
		Purple:         rgbaToNrgba(purple.RGBA()),
		PurpleDark:     rgbaToNrgba(gamut.Darker(purple, 0.2).RGBA()),
		PurpleDarker:   rgbaToNrgba(gamut.Darker(purple, 0.4).RGBA()),
		PurpleDarkest:  rgbaToNrgba(gamut.Darker(purple, 0.6).RGBA()),
	}
	return colors
}

func createColorImages() ColorImages {
	imageHash := ColorImages{
		PurpleDarkest:  ebitenuiImage.NewNineSliceColor(Colors.PurpleDarkest),
		PurpleDarker:   ebitenuiImage.NewNineSliceColor(Colors.PurpleDarker),
		PurpleDark:     ebitenuiImage.NewNineSliceColor(Colors.PurpleDark),
		Purple:         ebitenuiImage.NewNineSliceColor(Colors.Purple),
		PurpleLight:    ebitenuiImage.NewNineSliceColor(Colors.PurpleLight),
		PurpleLighter:  ebitenuiImage.NewNineSliceColor(Colors.PurpleLighter),
		PurpleLightest: ebitenuiImage.NewNineSliceColor(Colors.PurpleLightest),
	}

	return imageHash
}
