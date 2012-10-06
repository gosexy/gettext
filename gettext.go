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
*/
import "C"

var (
	// For all of the locale.
	LC_ALL = uint(C.LC_ALL)

	// For regular expression matching (it determines the meaning of range expressions and equivalence classes) and string collation.
	LC_COLATE = uint(C.LC_ALL)

	// For regular expression matching, character classification, conversion, case-sensitive comparison, and wide character functions.
	LC_CTYPE = uint(C.LC_CTYPE)

	// For localizable natural-language messages.
	LC_MESSAGES = uint(C.LC_MESSAGES)

	// For monetary formatting.
	LC_MONETARY = uint(C.LC_MONETARY)

	// For number formatting (such as the decimal point and the thousands separator).
	LC_NUMERIC = uint(C.LC_NUMERIC)

	// For time and date formatting.
	LC_TIME = uint(C.LC_TIME)
)

// Sets or queries the program's current locale.
func SetLocale(category uint, locale string) string {
	return C.GoString(C.setlocale(C.int(category), C.CString(locale)))
}

// Sets directory containing message catalogs.
func BindTextdomain(domainname string, dirname string) string {
	return C.GoString(C.bindtextdomain(C.CString(domainname), C.CString(dirname)))
}

// Sets the output codeset for message catalogs for domain domainname.
func BindTextdomainCodeset(domainname string, codeset string) string {
	return C.GoString(C.bind_textdomain_codeset(C.CString(domainname), C.CString(codeset)))
}

// Sets or retrieves the current message domain.
func Textdomain(domainname string) string {
	return C.GoString(C.textdomain(C.CString(domainname)))
}

// Attempt to translate a text string into the user's native language, by looking up the translation in a message
// catalog.
func Gettext(msgid string) string {
	return C.GoString(C.gettext(C.CString(msgid)))
}

// Like Gettext(), but looking up the message in the specified domain.
func DGettext(domain string, msgid string) string {
	return C.GoString(C.dgettext(C.CString(domain), C.CString(msgid)))
}

// Like Gettext(), but looking up the message in the specified domain and category.
func DCGettext(domain string, msgid string, category uint) string {
	return C.GoString(C.dcgettext(C.CString(domain), C.CString(msgid), C.int(category)))
}

// Attempt to translate a text string into the user's native language, by looking up the appropriate plural form
// of the translation in a message catalog.
func NGettext(msgid string, msgid_plural string, n uint64) string {
	return C.GoString(C.ngettext(C.CString(msgid), C.CString(msgid_plural), C.ulong(n)))
}

// Like NGettext(), but looking up the message in the specified domain.
func DNGettext(domainname string, msgid string, msgid_plural string, n uint64) string {
	return C.GoString(C.dngettext(C.CString(domainname), C.CString(msgid), C.CString(msgid_plural), C.ulong(n)))
}

// Like NGettext(), but looking up the message in the specified domain and category.
func DCNGettext(domainname string, msgid string, msgid_plural string, n uint64, category uint) string {
	return C.GoString(C.dcngettext(C.CString(domainname), C.CString(msgid), C.CString(msgid_plural), C.ulong(n), C.int(category)))
}
