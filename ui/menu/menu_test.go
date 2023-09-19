package menu

import (
	"fmt"
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestMenu(t *testing.T) {
	menuValue := map[string][]string{
		"":       {"item", "card"},
		"item":   {"item1", "item2"},
		"card":   {"card1", "card2"},
		"card1":  {"card11", "card12"},
		"card2":  {"card21", "card22"},
		"card21": {"card211", "card212"},
	}
	menuAction := map[string]func(){
		"item1":   func() { fmt.Println("item1") },
		"item2":   func() { fmt.Println("item2") },
		"card211": func() { fmt.Println("card211") },
	}
	a := test.NewApp()
	w := a.NewWindow("test")
	m := Menu{}
	m.New(w, menuValue, menuAction)
	w.ShowAndRun()

	if len(menuValue[""]) != len(m.MainMenu.Items) {
		t.Errorf("Expected menu label %d but got %d", len(menuValue[""]), len(m.MainMenu.Items))
	}

}
