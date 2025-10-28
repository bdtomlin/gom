package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Icon(name string, width, height, class string) Node {
	switch name {
	case "info":
		return InfoIcon(width, height, class)
	}
	return nil
}

// <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-info-icon lucide-info"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>
func InfoIcon(width, height, class string) Node {
	return SVG(
		Attr("xmlns", "http://www.w3.org/2000/svg"),
		Width(width),
		Height(height),
		Attr("viewBox", "0 0 24 24"),
		Attr("fill", "none"),
		Attr("stroke", "currentColor"),
		Attr("stroke-width", "2"),
		Attr("stroke-linecap", "round"),
		Attr("stroke-linejoin", "round"),
		Class(class),
		El(
			"circle",
			Attr("cx", "12"),
			Attr("cy", "12"),
			Attr("r", "10"),
		),
		El(
			"path",
			Attr("d", "M12 16v-4"),
		),
		El(
			"path",
			Attr("d", "M12 16v-4"),
			Attr("d", "M12 8h.01"),
		),
	)
}
