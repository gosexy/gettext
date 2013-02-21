# gosexy/gettext

Go bindings for [GNU gettext][1], an internationalization and localization
library for writing multilingual systems.

## Requeriments

The GNU C library. If you're using GNU/Linux, FreeBSD or OSX you should already
have it.

## Installation

Use `go get` to download and install the binding:

```sh
go get github.com/gosexy/gettext
```

## Usage

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

You can use `os.Setenv` to set the `LANGUAGE` environment variable or set it
on a terminal:

```sh
export LANGUAGE="es_MX.utf8"
./gettext-program
```

Note that `xgettext` does not officially support Go syntax yet, however, you
can generate a valid `.pot` file by forcing `xgettest` to use the C++
syntax:

```sh
xgettext -d example -s gettext_test.go -o example.pot -L c++ -i \
--keyword=NGettext:1,2 --keyword=Gettext
```

This will generate a `example.pot` file.

After translating the `.pot` file, you must generate `.po` and `.mo` files and
remember to set the UTF-8 charset.

```sh
msginit -l es_MX -o example.po -i example.pot
msgfmt -c -v -o example.mo example.po
```

Finally, move the `.mo` file to an appropriate location.

```sh
mv example.mo examples/es_MX.utf8/LC_MESSAGES/example.mo
```

## Documentation

You can read `gosexy/gettext` documentation from a terminal

```sh
go doc github.com/gosexy/gettext
```

Or you can [browse it](http://godoc.org/github.com/gosexy/gettext) online.

The original gettext documentation could be very useful as well:

```sh
man 3 gettext
```

Here's another [good tutorial][2] on using gettext.

[1]: http://www.gnu.org/software/gettext/
[2]: http://oriya.sarovar.org/docs/gettext_single.html
