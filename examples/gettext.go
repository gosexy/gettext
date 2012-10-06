package main

import (
	"github.com/gosexy/gettext"
	"fmt"
	"os"
)

func main() {
	gettext.BindTextdomain("example", ".")
	gettext.Textdomain("example")

	os.Setenv("LANGUAGE", "es_MX.utf8")

	gettext.SetLocale(gettext.LC_ALL, "")

	fmt.Println(gettext.Gettext("Hello, world!"))

	os.Setenv("LANGUAGE", "de_DE.utf8")

	gettext.SetLocale(gettext.LC_ALL, "")

	fmt.Println(gettext.Gettext("Hello, world!"))

	os.Setenv("LANGUAGE", "en_US.utf8")

	gettext.SetLocale(gettext.LC_ALL, "")

	fmt.Println(gettext.Gettext("Hello, world!"))
}

