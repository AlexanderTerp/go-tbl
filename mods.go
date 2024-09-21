package tblwriter

import (
	"fmt"
	"github.com/fatih/color"
)

type Color int

const (
	Plain Color = iota

	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	HiBlack
	HiRed
	HiGreen
	HiYellow
	HiBlue
	HiMagenta
	HiCyan
	HiWhite
)

type HeaderMod struct {
	color Color
}

type RowCellMod struct {
	color Color
}

type ColumnMod struct {
	color Color
}

var (
	LibBlack     = color.New(color.FgBlack)
	LibRed       = color.New(color.FgRed)
	LibGreen     = color.New(color.FgGreen)
	LibYellow    = color.New(color.FgYellow)
	LibBlue      = color.New(color.FgBlue)
	LibMagenta   = color.New(color.FgMagenta)
	LibCyan      = color.New(color.FgCyan)
	LibWhite     = color.New(color.FgWhite)
	LibHiBlack   = color.New(color.FgHiBlack)
	LibHiRed     = color.New(color.FgHiRed)
	LibHiGreen   = color.New(color.FgHiGreen)
	LibHiYellow  = color.New(color.FgHiYellow)
	LibHiBlue    = color.New(color.FgHiBlue)
	LibHiMagenta = color.New(color.FgHiMagenta)
	LibHiCyan    = color.New(color.FgHiCyan)
	LibHiWhite   = color.New(color.FgHiWhite)
)

func (c Color) toLibColor() *color.Color {
	switch c {
	case Plain:
		panic("Plain color should not be passed here")
	case Black:
		return LibBlack
	case Red:
		return LibRed
	case Green:
		return LibGreen
	case Yellow:
		return LibYellow
	case Blue:
		return LibBlue
	case Magenta:
		return LibMagenta
	case Cyan:
		return LibCyan
	case White:
		return LibWhite
	case HiBlack:
		return LibHiBlack
	case HiRed:
		return LibHiRed
	case HiGreen:
		return LibHiGreen
	case HiYellow:
		return LibHiYellow
	case HiBlue:
		return LibHiBlue
	case HiMagenta:
		return LibHiMagenta
	case HiCyan:
		return LibHiCyan
	case HiWhite:
		return LibHiWhite
	default:
		panic(fmt.Sprintf("Unknown color: %d", c))
	}
}
