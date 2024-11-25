package repository

func init() {
	SetMemberRepository(NewMemberRepositoryMySQL())
}
