package shared

type RandomRange struct {
	Origin, Bound int
}

type PasswordInfo struct {
	Range RandomRange
	Count int
}

type PasswordsInfo struct {
	Passwords []string
}
