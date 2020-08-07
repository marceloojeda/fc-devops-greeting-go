package main

import "testing"

func TestGreeting(t *testing.T) {
	greeting := greeting("Code.education Rocks!")
	if greeting != "<b>Code.education Rocks!</b>" {
		t.Errorf("Resposta incorreta. Chegou: %s, deveria: %s.", greeting, "<b>Code.education Rocks!</b>")
	}
}
