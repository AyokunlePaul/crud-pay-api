package password_service

type Service interface {
	Generate(string) (string, error)
	Compare(string, string) error
}
