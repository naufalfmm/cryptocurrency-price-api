package password

//go:generate mockgen -package=mockPassword -destination=./mockPassword/mock.go -source=password.go
type Password interface {
	Generate(password string) (string, error)
	Check(hashed, password string) error
}
