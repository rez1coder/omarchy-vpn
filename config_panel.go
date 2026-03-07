package main

import (
	"fmt"
	"strings"
)

func (m model) renderConfigPanel(width, height int) string {
	contentHeight := height - 2

	// Panel title
	var title string
	count := len(m.configs)
	if count > 0 {
		title = panelTitleStyle.Render(fmt.Sprintf("Configs (%d)", count))
	} else {
		title = panelTitleDimStyle.Render("Configs")
	}
	_ = title

	var lines []string

	if len(m.configs) == 0 {
		lines = append(lines, "")
		lines = append(lines, dimStyle.Render("  No configs found."))
		lines = append(lines, dimStyle.Render("  Press "+shortcutKeyStyle.Render("i")+" to import."))
	} else {
		lines = append(lines, "") // top padding
		for i, cfg := range m.configs {
			line := m.renderConfigItem(cfg, i, width-4)
			lines = append(lines, line)
		}
	}

	// Pad to fill height
	for len(lines) < contentHeight {
		lines = append(lines, "")
	}
	if len(lines) > contentHeight {
		lines = lines[:contentHeight]
	}

	content := strings.Join(lines, "\n")

	// Choose border style based on connection state
	border := activeBorderStyle
	if m.activeVPN != "" {
		border = connectedBorderStyle
	}

	return border.
		Width(width - 2).
		Height(contentHeight).
		Render(content)
}

func (m model) renderConfigItem(name string, index, maxWidth int) string {
	isActive := name == m.activeVPN
	isSelected := index == m.cursor

	// Handle inline modals
	if isSelected && m.modal == modalRenaming {
		return inputPromptStyle.Render("  ▸ ") + m.renameInput.View()
	}

	if isSelected && m.modal == modalDeleting {
		prompt := fmt.Sprintf("  ▸ Delete %s? ", name)
		return errorStyle.Render(prompt) + dimStyle.Render("[y/n]")
	}

	if isSelected && m.modal == modalConnecting && m.connectName == name {
		return "  " + selectedItemStyle.Render("▸ ") +
			m.spinner.View() + " " +
			selectedItemStyle.Render(name)
	}

	// Normal rendering
	var parts strings.Builder

	if isSelected {
		parts.WriteString("  ")
		parts.WriteString(selectedItemStyle.Render("▸ "))
		parts.WriteString(selectedItemStyle.Render(name))
	} else {
		parts.WriteString("    ")
		parts.WriteString(itemStyle.Render(name))
	}

	if isActive {
		parts.WriteString(" ")
		parts.WriteString(connectedIndicator)
	}

	return parts.String()
}
