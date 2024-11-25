package repository

func Init() {
	SetMemberRepository(NewMemberRepositoryMySQL())
}
