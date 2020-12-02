package day02_2020

type Password struct {
	Text   string
	Policy PasswordPolicy
}

type PasswordPolicy struct {
	Char rune
	Int1 int
	Int2 int
}
