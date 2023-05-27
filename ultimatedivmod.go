package piscine

func UltimateDivMod(a *int, b *int) {
	p := *a
	q := *b
	*a = p / q
	*b = p % q
}
