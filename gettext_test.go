package gettext

import (
	"fmt"
	"os"
	"testing"
)

/*
	NOTE:

	xgettext does not officially support Go syntax, however, you can generate a valid .pot file by forcing
	xgettest to use the C++ syntax:

	% xgettext -d example -s gettext_test.go -o example.pot -L c++ -i --keyword=NGettext:1,2 --keyword=Gettext

	This will generate a example.pot file.

	After translating the .pot file, you must generate .po and .mo files.

	Remember to set the UTF-8 charset.

	% msginit -l es_MX -o example.po -i example.pot
	% msgfmt -c -v -o example.mo example.po

	And finally, move the .mo file to an appropriate location.

	% mv example.mo examples/es_MX.utf8/LC_MESSAGES/example.mo

*/

func TestSpanishMexico(t *testing.T) {

	os.Setenv("LANGUAGE", "es_MX.utf8")

	SetLocale(LC_ALL, "")
	BindTextdomain("example", "./examples/")
	Textdomain("example")

	t1 := Gettext("Hello, world!")

	fmt.Println(t1)

	if t1 != "¡Hola mundo!" {
		t.Errorf("Failed translation.")
	}

	t2 := Sprintf(NGettext("An apple", "%d apples", 1), 1, "garbage")

	fmt.Println(t2)

	if t2 != "Una manzana" {
		t.Errorf("Failed translation.")
	}

	t3 := Sprintf(NGettext("An apple", "%d apples", 3), 3)

	fmt.Println(t3)

	if t3 != "3 manzanas" {
		t.Errorf("Failed translation.")
	}

	t4 := Gettext("Good morning")

	fmt.Println(t4)

	if t4 != "Buenos días" {
		t.Errorf("Failed translation.")
	}

	t5 := Gettext("Good bye!")

	fmt.Println(t5)

	if t5 != "¡Hasta luego!" {
		t.Errorf("Failed translation.")
	}

}

func TestGermanDeutschland(t *testing.T) {

	os.Setenv("LANGUAGE", "de_DE.utf8")

	SetLocale(LC_ALL, "")
	BindTextdomain("example", "./examples/")
	Textdomain("example")

	t1 := Gettext("Hello, world!")

	fmt.Println(t1)

	if t1 != "Hallo, Welt!" {
		t.Errorf("Failed translation.")
	}

	t2 := Sprintf(NGettext("An apple", "%d apples", 1), 1, "garbage")

	fmt.Println(t2)

	if t2 != "Ein Apfel" {
		t.Errorf("Failed translation.")
	}

	t3 := Sprintf(NGettext("An apple", "%d apples", 3), 3)

	fmt.Println(t3)

	if t3 != "3 Äpfel" {
		t.Errorf("Failed translation.")
	}

	t4 := Gettext("Good morning")

	fmt.Println(t4)

	if t4 != "Guten morgen" {
		t.Errorf("Failed translation.")
	}

	t5 := Gettext("Good bye!")

	fmt.Println(t5)

	if t5 != "Aufwiedersehen!" {
		t.Errorf("Failed translation.")
	}

}
