package user_domain

func (ths User) VerifyPassword(plainPassword string) bool {
	return ths.passwordHash.Verify(plainPassword)
}
