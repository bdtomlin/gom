package components

import (
	"app/forms"
	"strings"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

type InputOpts struct {
	Prefix  string
	Postfix string
	Label   string
}

func InputInt(field *forms.IntField, opts InputOpts, nodes ...Node) Node {
	return InputBase("text", field, opts, nodes...)
}

func InputString(field *forms.StringField, opts InputOpts, nodes ...Node) Node {
	return InputBase("text", field, opts, nodes...)
}

func InputBase(itype string, field forms.FieldInterface, opts InputOpts, nodes ...Node) Node {
	return Div(
		If(opts.Label != "",
			Label(
				For(field.Name()),
				Class("block text-sm/6 font-medium text-gray-900"),
				Text(opts.Label),
			),
		),
		Div(
			Classes{"flex": true, "mt-2": opts.Label != ""},
			If(opts.Prefix != "",
				Div(
					Class("flex shrink-0 items-center rounded-l-md bg-white px-3 text-base text-gray-500 outline-1 -outline-offset-1 outline-gray-300 sm:text-sm/6"), Text(opts.Prefix)),
			),

			Input(
				Type(itype),
				Name(field.Name()),
				Value(field.FormVal()),
				ID(field.Name()),
				Classes{
					"block w-full bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-blue-600 sm:text-sm/6": true,
					"outline-gray-300": !field.HasErrors(),
					"outline-red-300":  field.HasErrors(),
					"rounded-l":        opts.Prefix == "",
					"-ml-px":           opts.Prefix != "",
					"rounded-r":        opts.Postfix == "",
					"-mr-px":           opts.Postfix != "",
				},
				Group(nodes),
			),
			If(opts.Postfix != "",
				Div(
					Class("flex shrink-0 items-center rounded-l-md bg-white px-3 text-base text-gray-500 outline-1 -outline-offset-1 outline-gray-300 sm:text-sm/6"), Text("$")),
			),
		),
		If(field.HasErrors(),
			Div(
				Class("text-red-500"),
				Text(strings.Join(field.Errors(), ",")),
			),
		),
	)
}
