package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func EditBytes(data []byte) []byte {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	tmpfile, err := ioutil.TempFile("", "dsnote.*.md")
	if err != nil {
		fmt.Println("Failed to create temporary file")
		os.Exit(1)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write(data)
	if err != nil {
		fmt.Println("Failed to write to the temporary file")
		os.Exit(1)
	}

	err = RunCommand(editor, tmpfile.Name())
	if err != nil {
		fmt.Println("Failed to edit the temporary file")
		os.Exit(1)
	}

	data, err = ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		fmt.Println("Failed to read temporary file after editing")
		os.Exit(1)
	}
	return data
}

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
