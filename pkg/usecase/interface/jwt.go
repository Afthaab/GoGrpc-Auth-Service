package interfaces

type JwtUseCase interface {
	GenerateAccessToken(userid int, email string, role string) (string, error)
}
