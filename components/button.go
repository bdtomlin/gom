package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type BtnOpts struct {
	Type    string
	Content Node
}

func Btn(opts BtnOpts, content Node) Node {
	if opts.Type == "" {
		opts.Type = "button"
	}

	// <button type="button" class="">Button text</button>
	return Button(
		Type(opts.Type),
		Class("rounded-md bg-blue-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-xs hover:bg-blue-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600"),
		content,
	)
}
