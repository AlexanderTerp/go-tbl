package tblwriter

import (
	"github.com/fatih/color"
	"regexp"
	"strings"
)

var COLOR_ALL = regexp.MustCompile("[\\s\\S]*")

func (t *Table) colorize(s string, c Color) string {
	return t.colorizeRegex(s, c, COLOR_ALL)
}

func (t *Table) colorizeRegex(s string, c Color, regexp *regexp.Regexp) string {
	if t.forceNoColor || c == Plain || regexp == nil {
		return s
	}

	colorFunc := c.toLibColor().SprintFunc()

	var result strings.Builder
	lastIndex := 0
	matches := regexp.FindAllStringIndex(s, -1)

	for _, match := range matches {
		start, end := match[0], match[1]
		result.WriteString(s[lastIndex:start])
		result.WriteString(colorFunc(s[start:end]))
		lastIndex = end
	}

	result.WriteString(s[lastIndex:])
	return result.String()
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

func (t *Table) SetColumnMods(mods map[int]ColumnMod) {
	t.columnModsByIdx = mods
}
