# gosexy/gettext

Go bindings for [GNU gettext][1], an internationalization and localization
library for writing multilingual systems.

## Requeriments

The GNU C library. If you're using GNU/Linux, FreeBSD or OSX you probably
already have it.

## Installation

Use `go get` to download and install the binding:

```sh
go get github.com/gosexy/gettext
```

## Usage

This is an example program showing the `BindTextdomain`, `Textdomain` and
`SetLocale` bindings:

```go
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
}
```

Set the `LANGUAGE` env to the name of the language you want to use in your
program:

```sh
export LANGUAGE="es_MX.utf8"
./myapp
```

You can use the `xgettext` command to extract strings to be translated from a
Go program:

```
go get github.com/gosexy/gettext/go-xgettext

xgettext -o outfile.pot --keyword=Gettext --keyword-plural=NGettext -i infile.go
```

This will generate a `example.pot` file.

After actually translating the `.pot` file, you'll have to generate `.po` and
`.mo` files with `msginit` and `msgfmt`:

```sh
msginit -l es_MX -o example.po -i example.pot
msgfmt -c -v -o example.mo example.po
```

Finally, move the `.mo` file to an appropriate location.

```sh
mv example.mo examples/es_MX.utf8/LC_MESSAGES/example.mo
```

## Documentation

Check out the API documentation [godoc.org/github.com/gosexy/gettext)](http://godoc.org/github.com/gosexy/gettext).

The original gettext documentation:

```sh
man 3 gettext
```

And here's a [good tutorial][2] on using gettext.

[1]: http://www.gnu.org/software/gettext/
[2]: http://oriya.sarovar.org/docs/gettext_single.html
