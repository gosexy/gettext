/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package gettext

/*

#include <libintl.h>
#include <locale.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

var (
	// For all of the locale.
	LC_ALL = uint(C.LC_ALL)

	// For regular expression matching (it determines the meaning of range
	// expressions and equivalence classes) and string collation.
	LC_COLATE = uint(C.LC_ALL)

	// For regular expression matching, character classification, conversion,
	// case-sensitive comparison, and wide character functions.
	LC_CTYPE = uint(C.LC_CTYPE)

	// For localizable natural-language messages.
	LC_MESSAGES = uint(C.LC_MESSAGES)

	// For monetary formatting.
	LC_MONETARY = uint(C.LC_MONETARY)

	// For number formatting (such as the decimal point and the thousands
	// separator).
	LC_NUMERIC = uint(C.LC_NUMERIC)

	// For time and date formatting.
	LC_TIME = uint(C.LC_TIME)
)

// Sets or queries the program's current locale.
func SetLocale(category uint, locale string) string {
	clocale := C.CString(locale)

	res := C.GoString(C.setlocale(C.int(category), clocale))

	C.free(unsafe.Pointer(clocale))
	return res
}

// Sets directory containing message catalogs.
func BindTextdomain(domainname string, dirname string) string {
	cdirname := C.CString(dirname)
	cdomainname := C.CString(domainname)

	res := C.GoString(C.bindtextdomain(cdomainname, cdirname))

	C.free(unsafe.Pointer(cdirname))
	C.free(unsafe.Pointer(cdomainname))
	return res
}

// Sets the output codeset for message catalogs for domain domainname.
func BindTextdomainCodeset(domainname string, codeset string) string {
	cdomainname := C.CString(domainname)
	ccodeset := C.CString(codeset)

	res := C.GoString(C.bind_textdomain_codeset(cdomainname, ccodeset))

	C.free(unsafe.Pointer(cdomainname))
	C.free(unsafe.Pointer(ccodeset))
	return res
}

// Sets or retrieves the current message domain.
func Textdomain(domainname string) string {
	cdomainname := C.CString(domainname)

	res := C.GoString(C.textdomain(cdomainname))

	C.free(unsafe.Pointer(cdomainname))
	return res
}

// Attempt to translate a text string into the user's native language, by
// looking up the translation in a message catalog.
func Gettext(msgid string) string {
	cmsgid := C.CString(msgid)

	res := C.GoString(C.gettext(cmsgid))

	C.free(unsafe.Pointer(cmsgid))
	return res
}

// Like Gettext(), but looking up the message in the specified domain.
func DGettext(domain string, msgid string) string {
	cdomain := C.CString(domain)
	cmsgid := C.CString(msgid)

	res := C.GoString(C.dgettext(cdomain, cmsgid))

	C.free(unsafe.Pointer(cdomain))
	C.free(unsafe.Pointer(cmsgid))
	return res
}

// Like Gettext(), but looking up the message in the specified domain and
// category.
func DCGettext(domain string, msgid string, category uint) string {
	cdomain := C.CString(domain)
	cmsgid := C.CString(msgid)

	res := C.GoString(C.dcgettext(cdomain, cmsgid, C.int(category)))

	C.free(unsafe.Pointer(cdomain))
	C.free(unsafe.Pointer(cmsgid))
	return res
}

// Attempt to translate a text string into the user's native language, by
// looking up the appropriate plural form of the translation in a message
// catalog.
func NGettext(msgid string, msgid_plural string, n uint64) string {
	cmsgid := C.CString(msgid)
	cmsgid_plural := C.CString(msgid_plural)

	res := C.GoString(C.ngettext(cmsgid, cmsgid_plural, C.ulong(n)))

	C.free(unsafe.Pointer(cmsgid))
	C.free(unsafe.Pointer(cmsgid_plural))

	return res
}

// Like fmt.Sprintf() but without %!(EXTRA) errors.
func Sprintf(format string, a ...interface{}) string {
	expects := strings.Count(format, "%") - strings.Count(format, "%%")

	if expects > 0 {
		arguments := make([]interface{}, expects)
		for i := 0; i < expects; i++ {
			if len(a) > i {
				arguments[i] = a[i]
			}
		}
		return fmt.Sprintf(format, arguments...)
	}

	return format
}

// Like NGettext(), but looking up the message in the specified domain.
func DNGettext(domainname string, msgid string, msgid_plural string, n uint64) string {
	cdomainname := C.CString(domainname)
	cmsgid := C.CString(msgid)
	cmsgid_plural := C.CString(msgid_plural)

	res := C.GoString(C.dngettext(cdomainname, cmsgid, cmsgid_plural, C.ulong(n)))

	C.free(unsafe.Pointer(cdomainname))
	C.free(unsafe.Pointer(cmsgid))
	C.free(unsafe.Pointer(cmsgid_plural))

	return res
}

// Like NGettext(), but looking up the message in the specified domain and
// category.
func DCNGettext(domainname string, msgid string, msgid_plural string, n uint64, category uint) string {
	cdomainname := C.CString(domainname)
	cmsgid := C.CString(msgid)
	cmsgid_plural := C.CString(msgid_plural)

	res := C.GoString(C.dcngettext(cdomainname, cmsgid, cmsgid_plural, C.ulong(n), C.int(category)))

	C.free(unsafe.Pointer(cdomainname))
	C.free(unsafe.Pointer(cmsgid))
	C.free(unsafe.Pointer(cmsgid_plural))

	return res
}
