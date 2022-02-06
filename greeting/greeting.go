package greeting

import "errors"

func Greeting(country string) (string, error) {
	switch country {
	case "brazil":
		return "Olá meu amigo", nil
	case "france":
		return "Bonjour, mon ami", nil
	default:
		return "", errors.New("country past and invalid")
	}
}
