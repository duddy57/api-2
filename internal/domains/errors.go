package domains

import "errors"

var (
	ErrInvalidUserData       = errors.New("invalid user data")
	ErrInvalidUserName       = errors.New("name is required")
	ErrInvalidUserNameLength = errors.New("name must be between 2 and 100 characters")
	ErrInvalidUserEmail      = errors.New("email is required")
	ErrInvalidUserRole       = errors.New("role is required")
	ErrInvalidUserPassword   = errors.New("password is required")

	ErrDuplicatedEmailOrUsername = errors.New("duplicated email or username")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrUserNotFound              = errors.New("user not found")

	ErrInvalidDefectDescription    = errors.New("defect invalid")
	ErrInvalidDifficultyLevel      = errors.New("invalid difficulty level")
	ErrInvalidSolicitedBy          = errors.New("invalid solicited by")
	ErrInvalidClienteId            = errors.New("invalid client ID")
	ErrInvalidTecnicoResponsavelId = errors.New("invalid technician responsible ID")
	ErrInvalidDataDeAbertura       = errors.New("invalid open date")
	ErrInvalidSolutionDescription  = errors.New("solution description invalid")

	// Client validation errors
	ErrInvalidClientName    = errors.New("client name is required")
	ErrInvalidClientType    = errors.New("client type is required")
	ErrInvalidCnpjOrCpf     = errors.New("cnpj or cpf is required")
	ErrInvalidContactPerson = errors.New("contact person is required")
	ErrInvalidContactEmail  = errors.New("contact email is required")
	ErrInvalidContactPhone  = errors.New("contact phone is required")
	ErrInvalidContactName   = errors.New("contact name is required")
	ErrInvalidAddress       = errors.New("address is required")
	ErrInvalidPostalCode    = errors.New("postal code is required")
	ErrInvalidCountry       = errors.New("country is required")
	ErrInvalidState         = errors.New("state is required")
	ErrInvalidCity          = errors.New("city is required")
	ErrInvalidStreet        = errors.New("street is required")
	ErrInvalidNumber        = errors.New("number is required")

	ErrNoContent = errors.New("no content")
)

type ErrorResponse struct {
	ErrorResponse map[string]string `json:"error"`
}
