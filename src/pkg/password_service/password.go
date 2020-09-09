package password_service

import "golang.org/x/crypto/bcrypt"

type password struct{}

func New() Service {
	return &password{}
}

func (password *password) Generate(passwordString string) (string, error) {
	hashedPassword, hashingError := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if hashingError != nil {
		return "", hashingError
	} else {
		return string(hashedPassword), nil
	}
}

func (password *password) Compare(hashedPassword, passwordString string) error {
	if comparisonError := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordString));
		comparisonError != nil {
		return comparisonError
	}
	return nil
}
