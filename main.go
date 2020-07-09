package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/howeyc/gopass"
	"github.com/writeas/go-writeas"
)

func main() {
	u := flag.String("u", "", "Write.as username")
	coll := flag.String("c", "", "Write.as collection to import to")
	font := flag.String(
		"font",
		"norm",
		"The font for the post (norm, sans, wrap, mono, code)",
	)

	flag.Parse()
	args := flag.Args()

	if *u == "" || len(args) == 0 {
		fmt.Fprintf(
			os.Stderr,
			"usage: baudelaire -u username [-c blog] file1 [file2|file3...]\n",
		)
		os.Exit(1)
	}

	fmt.Print("Password: ")
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading pass: %v\n", err)
		os.Exit(1)
	}

	if len(pass) == 0 {
		fmt.Fprintf(os.Stderr, "Please enter your password.\n")
		os.Exit(1)
	}

	c := writeas.NewClient()

	fmt.Print("Logging in...")
	au, err := c.LogIn(*u, string(pass))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("OK\n")
	c.SetToken(au.AccessToken)

	for _, fn := range args {
		// Read file contents
		fmt.Print("Reading file...")
		content, err := ioutil.ReadFile(fn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %s: %v\n", fn, err)
			continue
		}
		fmt.Print("OK\n")
		fmt.Print("Publishing...")
		p, err := c.CreatePost(&writeas.PostParams{
			Title:      "",
			Content:    string(content),
			Font:       *font,
			Collection: *coll,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error publishing %s: %v\n", fn, err)
			continue
		}
		if *coll != "" {
			fmt.Printf("Created post %s from %s\n", p.Slug, fn)
		} else {
			fmt.Printf("Created post %s from %s\n", p.ID, fn)
		}
	}
}
