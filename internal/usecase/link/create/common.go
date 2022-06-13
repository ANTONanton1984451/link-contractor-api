package create

type ValidateErr struct {
	ValidateRule string
}

func (ve ValidateErr) Error() string {
	return "not valid"
}
