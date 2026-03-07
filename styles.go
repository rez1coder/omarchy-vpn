package main

import "charm.land/lipgloss/v2"

// Catppuccin Mocha
var (
	green    = lipgloss.Color("#a6e3a1")
	red      = lipgloss.Color("#f38ba8")
	blue     = lipgloss.Color("#89b4fa")
	yellow   = lipgloss.Color("#f9e2af")
	lavender = lipgloss.Color("#b4befe")
	mauve    = lipgloss.Color("#cba6f7")
	teal     = lipgloss.Color("#94e2d5")
	peach    = lipgloss.Color("#fab387")
	textCol  = lipgloss.Color("#cdd6f4")
	subtext  = lipgloss.Color("#a6adc8")
	overlay0 = lipgloss.Color("#6c7086")
	surface0 = lipgloss.Color("#313244")
	surface1 = lipgloss.Color("#45475a")
	base     = lipgloss.Color("#1e1e2e")

	// Title bar
	titleStyle = lipgloss.NewStyle().
			Foreground(blue).
			Bold(true)

	titleAccentStyle = lipgloss.NewStyle().
				Foreground(lavender)

	// Panel borders
	activeBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(blue)

	connectedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(green)

	inactiveBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(surface1)

	// Panel titles (rendered inside top border)
	panelTitleStyle = lipgloss.NewStyle().
			Foreground(blue).
			Bold(true).
			Background(base).
			Padding(0, 1)

	panelTitleConnectedStyle = lipgloss.NewStyle().
					Foreground(green).
					Bold(true).
					Background(base).
					Padding(0, 1)

	panelTitleDimStyle = lipgloss.NewStyle().
				Foreground(overlay0).
				Background(base).
				Padding(0, 1)

	// List items
	itemStyle = lipgloss.NewStyle().
			Foreground(subtext)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(green).
				Bold(true)

	activeMarkerStyle = lipgloss.NewStyle().
				Foreground(green)

	// Status field icons and labels
	labelStyle = lipgloss.NewStyle().
			Foreground(overlay0).
			Width(14)

	valueStyle = lipgloss.NewStyle().
			Foreground(textCol)

	// Feedback
	connectedStyle = lipgloss.NewStyle().Foreground(green)
	errorStyle     = lipgloss.NewStyle().Foreground(red)
	warnStyle      = lipgloss.NewStyle().Foreground(yellow)
	dimStyle       = lipgloss.NewStyle().Foreground(overlay0)

	// Bottom bar
	shortcutKeyStyle = lipgloss.NewStyle().
				Foreground(lavender).
				Bold(true)


	// Help overlay
	helpOverlayStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(mauve).
				Padding(1, 3)

	helpTitleStyle = lipgloss.NewStyle().
			Foreground(mauve).
			Bold(true)

	// Inline input
	inputPromptStyle = lipgloss.NewStyle().
				Foreground(blue)

	// Spinner
	spinnerStyle = lipgloss.NewStyle().
			Foreground(blue)

	// Connection status indicator
	connectedIndicator   = lipgloss.NewStyle().Foreground(green).Bold(true).Render("●")
	disconnectedIndicator = lipgloss.NewStyle().Foreground(overlay0).Render("○")
)

