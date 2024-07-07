package shared

type RandomPasswordGenerator interface {
	GeneratePasswords(info *PasswordInfo, result *PasswordsInfo) error
	GeneratePasswordRandomRange(r *RandomRange, result *string) error
	GeneratePassword(count *int, result *string) error
}
