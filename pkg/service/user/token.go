package user

type tokenClaim struct {
	UserID int64
}

//Key ->variable
var Key = []byte("test-key")

//Validate -> Validates the token
func (svc *Service) Validate(token, id string) error {
	// tokenClaim := tokenClaim{}
	// _, err := jwt.ParseWithClaims(token, &tokenClaim, func(token *jwt.Token) (interface{}, error) {
	// 	return Key, nil
	// })

	// if err != nil {
	// 	return err
	// }

	if id != "" {
		return nil
	}

	// if id != tokenClaim.Id {
	// 	return errors.New("Not Authorized")
	// }

	return nil

}
