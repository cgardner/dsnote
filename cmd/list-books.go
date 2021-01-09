package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listBooks)
}

type BookInfo struct {
	Name      string `json:"name"`
	NoteCount int    `json:"noteCount"`
}

var listBooks = &cobra.Command{
	Use:   "list-books",
	Short: "List the existing note books",
	Run: func(cmd *cobra.Command, args []string) {
		bookDir := filepath.Join(rootDir, "books")

		bookInfo, err := ioutil.ReadDir(bookDir)
		if err != nil {
			fmt.Printf("Error reading Book Directory")
			os.Exit(3)
		}

		books := []BookInfo{}

		for _, book := range bookInfo {
			if book.IsDir() == false {
				continue
			}

			bookPath := filepath.Join(bookDir, book.Name())
			bookNoteInfo, err := ioutil.ReadDir(bookPath)

			if err != nil {
				fmt.Println("Error reading Book Directory", bookPath)
				os.Exit(3)
			}

			books = append(books, BookInfo{
				Name:      book.Name(),
				NoteCount: len(bookNoteInfo),
			})
		}

		if isatty.IsTerminal(os.Stdout.Fd()) {
			for _, book := range books {
				fmt.Println(book.Name, " - ", book.NoteCount)
			}
		} else {
			data, err := json.MarshalIndent(books, "", "  ")
			if err != nil {
				fmt.Println("Error Marshalling JSON:", err)
				os.Exit(4)
			}
			io.Copy(os.Stdout, bytes.NewBuffer(data))

		}

	},
}
