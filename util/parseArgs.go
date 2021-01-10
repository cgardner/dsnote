package util

import (
	"fmt"
)

type CommandOptions struct {
	Options []string
	Tags    []string
	Book    string
}

type OptionsModifier func(opts *CommandOptions)

func WithTags(tags ...string) OptionsModifier {
	return func(opts *CommandOptions) {
		opts.Tags = tags
	}
}

func WithOptions(options ...string) OptionsModifier {
	return func(opts *CommandOptions) {
		opts.Options = options
	}
}

func WithBook(book string) OptionsModifier {
	return func(opts *CommandOptions) {
		opts.Book = book
	}
}

func NewCommandOptions(modifiers ...OptionsModifier) *CommandOptions {
	opts := &CommandOptions{}
	for _, modifier := range modifiers {
		modifier(opts)
	}
	return opts
}

func ParseArgs(args []string) *CommandOptions {
	options := NewCommandOptions()
	for _, arg := range args {
		if arg[0] == '+' {
			tag := arg[1:]
			fmt.Println("added a tag", tag)
			options.Tags = append(options.Tags, tag)
			continue
		}

		if arg[0:4] == "book" {
			options.Book = arg[5:]
			continue
		}

		options.Options = append(options.Options, arg)
	}

	return options
}
