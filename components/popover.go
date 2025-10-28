package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func PopoverComponent(id string, buttonContent Node, popoverContent ...Node) Node {
	return Span(
		El("el-popover",
			Attr("popover", ""),
			Attr("anchor", "top"),
			ID(id),
			Class("p-2 rounded-sm relative bg-black text-white overflow-visible w-48 text-center -mt-2"),
			Div(
				Class("text-sm"),
				Group(popoverContent),
			),
			Div(
				Class("absolute bottom-0 right-22 transform translate-y-1/2 w-4 h-4 bg-black rotate-45 shadow-lg z-10"),
			),
		),
		Button(
			Attr("popovertarget", id),
			Class("inline ml-2"),
			buttonContent,
		),
	)
}
