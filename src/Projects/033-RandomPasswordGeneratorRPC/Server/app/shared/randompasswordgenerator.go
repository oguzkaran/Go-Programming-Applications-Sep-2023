package shared

type RandomPasswordGenerator interface {
	GeneratePasswords(info *PasswordInfo, result *PasswordsInfo) error
	//GeneratePassword(info *RandomRange, result *string) error
}
