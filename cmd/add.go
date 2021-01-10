package cmd

import (
	"fmt"
	"strings"

	"github.com/cgardner/dsnote/util"
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Run: func(cmd *cobra.Command, args []string) {
		options := util.ParseArgs(args)

		fmt.Printf("Adding %s with %s tags in %s book\n", options.Options[0], options.Tags, options.Book)

		tags := "  - " + strings.Join(options.Tags, "\n  - ")

		template := `---
UUID: ` + GetUUID4() + `
Title: ` + options.Options[0] + `
Tags: 
` + tags + `
---
`
		newNote := string(util.EditBytes([]byte(template)))
		fmt.Println("After Editing", newNote)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func GetUUID4() string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return u.String()
}
