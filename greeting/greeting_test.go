package greeting

import "testing"

func TestGreeting(t *testing.T) {
	greeting := []struct {
		country string
		expect  string
	}{
		{"brazil", "Olá meu amigo"},
		{"france", "Bonjour, mon ami"},
	}

	for _, tt := range greeting {
		g, err := Greeting(tt.country)
		if err != nil {
			t.Errorf("Expected error was nil but it arrived %v", err)
		}
		if tt.expect != g {
			t.Errorf("Expected greeting was %s but %s arrived", tt.expect, g)
		}
	}
}
func TestGreetingError(t *testing.T) {
	g, err := Greeting("")
	expect := "country past and invalid"
	if g != "" {
		t.Errorf("An empty string was expected but it arrived %v", g)
	}
	if err.Error() != expect {
		t.Errorf("an error '%v' was expected but '%v' arrived", expect, err.Error())
	}
}
