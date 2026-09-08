package main

import (
	"bytes"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestAccessibleForm(t *testing.T) {
	var out bytes.Buffer
	var name string
	f := newForm(&name).WithAccessible(true).WithInput(strings.NewReader("Ada\n")).WithOutput(&out)
	if err := f.Run(); err != nil {
		t.Fatal(err)
	}
	if name != "Ada" {
		t.Fatal("value not retained")
	}
	if !strings.Contains(out.String(), "Name") {
		t.Fatal("prompt missing")
	}
}
func TestEmbeddedInitializationAndCancel(t *testing.T) {
	var name string
	m := model{form: newForm(&name)}
	if m.Init() == nil {
		t.Fatal("form init command lost")
	}
	n, cmd := m.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})
	if !n.(model).canceled || cmd == nil {
		t.Fatal("cancel lost")
	}
	if n.(model).View().Content != "Canceled" {
		t.Fatal("partial form treated as completed")
	}
}

func TestAccessibleEOFIsNotSubmission(t *testing.T) {
	var name string
	var out bytes.Buffer
	f := newForm(&name).WithAccessible(true).WithInput(strings.NewReader("")).WithOutput(&out)
	_ = f.Run()
	if validateName(name) == nil {
		t.Fatal("EOF accepted as a completed name")
	}
}
