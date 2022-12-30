package tab

import (
	"github.com/TypicalAM/goread/internal/style"
	"github.com/charmbracelet/lipgloss"
)

var (
	statusBarCell = lipgloss.NewStyle().
			Bold(true).
			Padding(0, 1).
			Foreground(style.BasicColorscheme.BgDark)

	iconColors = map[TabType]lipgloss.Color{
		Welcome:  style.BasicColorscheme.Color4,
		Category: style.BasicColorscheme.Color5,
		Feed:     style.BasicColorscheme.Color3,
	}

	icons = map[TabType]string{
		Welcome:  "﫢",
		Category: "﫜",
		Feed:     "",
	}

	texts = map[TabType]string{
		Welcome:  "MAIN",
		Category: "CATEGORY",
		Feed:     "FEED",
	}
)

// Style the text depending on the type of the tab
func AttachIconToTab(text string, tabType TabType, isActive bool) string {
	var iconStyle, textStyle lipgloss.Style
	if isActive {
		iconStyle = style.ActiveTabIcon
		textStyle = style.ActiveTab
	} else {
		iconStyle = style.TabIcon
		textStyle = style.TabStyle
	}

	// Cut the text if the tab length is too much to handle
	if len(text) > 12 {
		text = text[:12] + ""
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		iconStyle.Copy().Foreground(iconColors[tabType]).Render(icons[tabType]),
		textStyle.Render(text),
	)
}

// Style the status bar cell depending on the the of the current tab
func StyleStatusBarCell(tabType TabType) string {
	return statusBarCell.Copy().
		Background(iconColors[tabType]).
		Render(texts[tabType])
}
