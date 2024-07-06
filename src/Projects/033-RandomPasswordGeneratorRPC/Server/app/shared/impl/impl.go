package impl

import (
	"Server/app/shared"
	"Server/csd/util/str"
	"errors"
	"math/rand/v2"
)

type RandomPasswordGenerator int

func (rg *RandomPasswordGenerator) GeneratePasswords(info *shared.PasswordInfo, result *shared.PasswordsInfo) error {
	origin := info.Range.Origin
	bound := info.Range.Bound
	count := info.Count
	if count <= 0 || origin >= bound {
		return errors.New("invalid values for generating passwords")
	}

	for i := 0; i < info.Count; i++ {
		result.Passwords = append(result.Passwords, str.GenerateRandomTextEN(rand.IntN(bound-origin)+origin))
	}

	return nil
}
