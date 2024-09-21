package tblwriter

import "github.com/fatih/color"

func (t *Table) colorize(s string, c Color) string {
	if t.forceNoColor || c == Plain {
		return s
	}

	libColor := c.toLibColor()
	return libColor.SprintFunc()(s)
}

func (t *Table) ToggleColor(enabled bool) {
	color.NoColor = !enabled
}

func (t *Table) SetHeaderColors(colors ...Color) {
	if t.headerMods == nil || len(t.headerMods) == 0 {
		t.headerMods = make([]HeaderMod, len(colors))
	}

	for i, c := range colors {
		t.headerMods[i].color = c
	}
}

func (t *Table) SetHeaderMods(mods ...HeaderMod) {
	t.headerMods = mods
}

func (t *Table) SetColumnMods(mods ...ColumnMod) {
	t.columnMods = mods
}
