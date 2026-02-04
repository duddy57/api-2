package handlers

import (
	"encoding/json"
	"net/http"
	"olidesk-api-2/internal/handlers/spec"
	"olidesk-api-2/internal/usecase"

	"github.com/discord-gophers/goapi-gen/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handlers struct {
	validator      *validator.Validate
	logger         *zap.Logger
	usersUsecase   usecase.UserUseCase
	clientsUsecase usecase.ClientUseCase
	formsUsecase   usecase.FormsUseCase
}

func NewHandlers(logger *zap.Logger, usersUsecase usecase.UserUseCase, clientsUsecase usecase.ClientUseCase, formsUsecase usecase.FormsUseCase) Handlers {
	validator := validator.New(validator.WithRequiredStructEnabled())
	return Handlers{
		validator,
		logger,
		usersUsecase,
		clientsUsecase,
		formsUsecase,
	}
}

// Create client
// (POST /v1/clients/create)
func (api *Handlers) PostCreateClient(w http.ResponseWriter, r *http.Request) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.PostCreateClientJSON400Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	var payload spec.CriarCliente
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return spec.PostCreateClientJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	if err := api.validator.Struct(payload); err != nil {
		return spec.PostCreateClientJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	id, err := api.clientsUsecase.CreateClient(usecase.CreateClientInput{
		ClientName: payload.NomeCliente,
		ClientType: payload.TipoCliente.ToValue(),
		Contact: usecase.ContactPerson{
			Email:          string(payload.EmailContato),
			Phone:          payload.TelefoneContato,
			ResposableName: payload.NomeContato,
		},
		Address: usecase.Address{
			Neighborhood: payload.Endereco.Bairro,
			PostalCode:   payload.Endereco.Cep,
			City:         payload.Endereco.Cidade,
			Complement:   payload.Endereco.Complement,
			Latitude:     payload.Endereco.Latitude,
			Longitude:    payload.Endereco.Longitude,
			Number:       payload.Endereco.Numero,
			Country:      payload.Endereco.Pais,
			State:        payload.Endereco.Estado,
			Street:       payload.Endereco.Rua,
		},
	}, r.Context())
	if err != nil {
		return spec.PostCreateClientJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.PostCreateClientJSON200Response(spec.Resp200{
		Message: "Cliente criado com sucesso",
		ID:      id.String(),
	})

}

// Get all clients
// (GET /v1/clients/list)
func (api *Handlers) GetV1clientsList(w http.ResponseWriter, r *http.Request) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.GetV1clientsListJSON500Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	clients, err := api.clientsUsecase.ListClient(r.Context())
	if err != nil {
		return spec.GetV1clientsListJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	clientsList := make([]spec.Cliente, 0, len(clients))
	for _, client := range clients {
		clientsList = append(clientsList, spec.Cliente{
			ID:              client.ID.String(),
			EmailContato:    types.Email(client.Contact.Email),
			NomeContato:     client.Contact.ResposableName,
			TelefoneContato: client.Contact.Phone,
			CnpjOuCpf:       client.CnpjOrCpf,
			TipoCliente:     getClientType(client.ClientType),
			NomeCliente:     client.ClientName,

			Endereco: spec.Endereco{
				Bairro:     client.Address.Neighborhood,
				Cep:        client.Address.PostalCode,
				Cidade:     client.Address.City,
				Complement: client.Address.Complement,
				Latitude:   client.Address.Latitude,
				Longitude:  client.Address.Longitude,
				Numero:     client.Address.Number,
				Pais:       client.Address.Country,
				Estado:     client.Address.State,
				Rua:        client.Address.Street,
			},

			CreatedAt: client.CreatedAt.UTC(),
			UpdatedAt: client.UpdatedAt.UTC(),
		})
	}

	return spec.GetV1clientsListJSON200Response(spec.ListaClientes{
		Clientes: clientsList,
	})

}

// Delete client
// (DELETE /v1/clients/delete/{clientID})
func (api *Handlers) DeleteClient(w http.ResponseWriter, r *http.Request, clientID string) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.DeleteClientJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	id := uuid.MustParse(clientID)

	if err := api.clientsUsecase.DeleteClient(id, r.Context()); err != nil {
		return spec.DeleteClientJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.DeleteClientJSON204Response(spec.Resp204{
		Message: "Cliente deletado com sucesso",
	})
}

// Update client
// (PUT /v1/clients/update/{clientID})
func (api *Handlers) PutClient(w http.ResponseWriter, r *http.Request, clientID string) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.PutClientJSON400Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	id := uuid.MustParse(clientID)

	var payload spec.AtualizarCliente
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return spec.PutClientJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	if err := api.validator.Struct(payload); err != nil {
		return spec.PutClientJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	if err := api.clientsUsecase.UpdateClient(id, usecase.UpdateClientInput{
		ClientName: payload.NomeCliente,
		ClientType: payload.TipoCliente.ToValue(),
		Contact: usecase.ContactPerson{
			Email:          string(payload.EmailContato),
			Phone:          payload.TelefoneContato,
			ResposableName: payload.NomeContato,
		},
		Address: usecase.Address{
			Neighborhood: payload.Endereco.Bairro,
			PostalCode:   payload.Endereco.Cep,
			City:         payload.Endereco.Cidade,
			Complement:   payload.Endereco.Complement,
			Latitude:     payload.Endereco.Latitude,
			Longitude:    payload.Endereco.Longitude,
			Number:       payload.Endereco.Numero,
			Country:      payload.Endereco.Pais,
			State:        payload.Endereco.Estado,
			Street:       payload.Endereco.Rua,
		},
	}, r.Context()); err != nil {
		return spec.PutClientJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.PutClientJSON204Response(spec.Resp204{
		Message: "Cliente atualizado com sucesso",
	})

}

// Get client by ID
// (GET /v1/clients/{clientID})
func (api *Handlers) GetByIDClient(w http.ResponseWriter, r *http.Request, clientID string) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.GetByIDClientJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	id := uuid.MustParse(clientID)

	c, err := api.clientsUsecase.GetClient(id, r.Context())
	if err != nil {
		return spec.GetByIDClientJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.GetByIDClientJSON200Response(spec.BuscaCliente{
		Cliente: &spec.Cliente{
			ID:              c.ID.String(),
			CnpjOuCpf:       c.CnpjOrCpf,
			NomeCliente:     c.ClientName,
			TipoCliente:     getClientType(c.ClientType),
			EmailContato:    types.Email(c.Contact.Email),
			TelefoneContato: c.Contact.Phone,
			NomeContato:     c.Contact.ResposableName,
			Endereco: spec.Endereco{
				Bairro:     c.Address.Neighborhood,
				Cep:        c.Address.PostalCode,
				Cidade:     c.Address.City,
				Complement: c.Address.Complement,
				Latitude:   c.Address.Latitude,
				Longitude:  c.Address.Longitude,
				Numero:     c.Address.Number,
				Pais:       c.Address.Country,
				Estado:     c.Address.State,
				Rua:        c.Address.Street,
			},
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
	})
}

// Form client
// (POST /v1/forms/create)
func (api *Handlers) PostCreateForm(w http.ResponseWriter, r *http.Request) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.PostCreateFormJSON400Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	var payload spec.CriarFormulario

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return spec.PostCreateFormJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}
	if err := api.validator.Struct(payload); err != nil {
		return spec.PostCreateFormJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	tecIDs := make([]uuid.UUID, 0, len(payload.TecnicosResponsavel))
	for _, tecID := range payload.TecnicosResponsavel {
		tecIDs = append(tecIDs, uuid.MustParse(tecID))
	}

	id, err := api.formsUsecase.CreateForm(usecase.CreateFormInput{
		TecnicoResponsavelId: tecIDs,
		DataDeAbertura:       payload.DataOcorrencia,
		ClienteId:            uuid.MustParse(payload.ClienteID),
		SolicitedBy:          payload.Solicitante,
		DifficultyLevel:      payload.NivelDificuldade.ToValue(),
		DefectDescription:    payload.DescricaoDefeito,
		SolutionDescription:  payload.DescricaoSolucao,
	}, r.Context())
	if err != nil {
		return spec.PostCreateFormJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.PostCreateFormJSON201Response(spec.Resp200{
		Message: "Formulário criado com sucesso",
		ID:      id.String(),
	})
}

// Delete form
// (DELETE /v1/forms/delete/{formID})
func (api *Handlers) DeleteForm(w http.ResponseWriter, r *http.Request, formID string) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.DeleteFormJSON400Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	if err := api.formsUsecase.DeleteForm(uuid.MustParse(formID), r.Context()); err != nil {
		return spec.DeleteFormJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.DeleteFormJSON204Response(spec.Resp204{
		Message: "Formulário deletado com sucesso",
	})
}

// List forms
// (GET /v1/forms/list)
func (api *Handlers) ListForms(w http.ResponseWriter, r *http.Request) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.ListFormsJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	rawForms, err := api.formsUsecase.ListForms(r.Context())
	if err != nil {
		return spec.ListFormsJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	listForm := make([]spec.Formulario, 0, len(rawForms.Forms))
	for _, f := range rawForms.Forms {
		listTecnicos := make([]spec.Tecnico, 0, len(f.TecnicoResponsavelId))
		for _, t := range f.TecnicoResponsavelId {
			listTecnicos = append(listTecnicos, spec.Tecnico{
				ID:   t.ID.String(),
				Nome: t.Name,
			})
		}

		listForm = append(listForm, spec.Formulario{
			ID:                  f.ID.String(),
			DataOcorrencia:      f.DataDeAbertura,
			Solicitante:         f.SolicitedBy,
			NivelDificuldade:    getLevel(f.DifficultyLevel),
			DescricaoDefeito:    f.DefectDescription,
			DescricaoSolucao:    f.SolutionDescription,
			TecnicosResponsavel: listTecnicos,
			UpdatedAt:           f.UpdatedAt.UTC(),
			CreatedAt:           f.CreatedAt.UTC(),
		})
	}

	return spec.ListFormsJSON200Response(spec.ListaFormulario{
		Formularios: listForm,
	})
}

// Update form
// (PUT /v1/forms/update/{formID})
func (api *Handlers) PutForm(w http.ResponseWriter, r *http.Request, formID string) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.PutFormJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	var payload spec.AtualizarFormulario
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return spec.PostCreateFormJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}
	if err := api.validator.Struct(payload); err != nil {
		return spec.PostCreateFormJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	tecIDs := make([]uuid.UUID, 0, len(payload.TecnicosResponsavel))
	for _, tecID := range payload.TecnicosResponsavel {
		tecIDs = append(tecIDs, uuid.MustParse(tecID))
	}

	if err := api.formsUsecase.UpdateForm(
		uuid.MustParse(formID),
		usecase.UpdateFormInput{
			SolicitedBy:          payload.Solicitante,
			DifficultyLevel:      payload.NivelDificuldade.ToValue(),
			DataDeAbertura:       payload.DataOcorrencia.UTC(),
			TecnicoResponsavelId: tecIDs,
			DefectDescription:    payload.DescricaoDefeito,
			SolutionDescription:  payload.DescricaoSolucao,
		},
		r.Context(),
	); err != nil {
		return spec.PutFormJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.PutFormJSON204Response(spec.Resp204{
		Message: "Formulário atualizado com sucesso",
	})
}

// Get forms
// (GET /v1/forms/{formID})
func (api *Handlers) GetFormByID(w http.ResponseWriter, r *http.Request, formID string) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.GetFormByIDJSON400Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	f, err := api.formsUsecase.GetForm(uuid.MustParse(formID), r.Context())
	if err != nil {
		return spec.GetFormByIDJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	listTecnicos := make([]spec.Tecnico, 0, len(f.Form.TecnicoResponsavelId))
	for _, t := range f.Form.TecnicoResponsavelId {
		listTecnicos = append(listTecnicos, spec.Tecnico{
			ID:   t.ID.String(),
			Nome: t.Name,
		})
	}

	return spec.GetFormByIDJSON200Response(spec.BuscaFormulario{
		Formulario: spec.Formulario{
			ID:                  f.Form.ID.String(),
			DataOcorrencia:      f.Form.DataDeAbertura,
			Solicitante:         f.Form.SolicitedBy,
			NivelDificuldade:    getLevel(f.Form.DifficultyLevel),
			DescricaoDefeito:    f.Form.DefectDescription,
			DescricaoSolucao:    f.Form.SolutionDescription,
			TecnicosResponsavel: listTecnicos,
			UpdatedAt:           f.Form.UpdatedAt.UTC(),
			CreatedAt:           f.Form.CreatedAt.UTC(),
		},
	})

}

// Get members
// (GET /v1/members/list)
func (api *Handlers) ListMembers(w http.ResponseWriter, r *http.Request) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.ListMembersJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	members, err := api.usersUsecase.GetMembers(r.Context())
	if err != nil {
		return spec.ListMembersJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	usuarios := make([]spec.Usuario, 0, len(members))
	for _, member := range members {
		usuarios = append(usuarios, spec.Usuario{
			Cargo: member.Role,
			ID:    member.ID.String(),
			Nome:  member.Name,
			Email: types.Email(member.Email),
		})
	}

	return spec.ListMembersJSON200Response(spec.ListaUsuarios{Usuarios: usuarios})
}

// Create a new user
// (POST /v1/users/create)
func (api *Handlers) PostCreateUser(w http.ResponseWriter, r *http.Request) *spec.Response {
	_, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.GetUserAccountJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	var payload spec.CriarUsuario
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		api.logger.Error("failed to decode user", zap.Error(err))
		return spec.PostCreateUserJSON400Response(spec.ErrorResponse{
			Message: err.Error(),
		})
	}

	id, err := api.usersUsecase.CreateUser(usecase.CreateUserInput{
		Name:     payload.Nome,
		Email:    string(payload.Email),
		Password: payload.Password,
		Role:     payload.Cargo.ToValue(),
	}, r.Context())
	if err != nil {
		return spec.PostCreateUserJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.PostCreateUserJSON200Response(spec.Resp200{
		ID:      id.String(),
		Message: "Usuário criado com sucesso",
	})
}

// Delete user
// (DELETE /v1/users/delete)
func (api *Handlers) DeleteUserAccount(w http.ResponseWriter, r *http.Request) *spec.Response {
	userID, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.DeleteUserAccountJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	if err := api.usersUsecase.DeleteUser(userID, r.Context()); err != nil {
		return spec.DeleteUserAccountJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.DeleteUserAccountJSON204Response(spec.Resp204{
		Message: SuccessMessage,
	})

}

// Get user
// (GET /v1/users/details)
func (api *Handlers) GetUserAccount(w http.ResponseWriter, r *http.Request) *spec.Response {
	userID, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.GetUserAccountJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	user, err := api.usersUsecase.GetUser(userID, r.Context())
	if err != nil {
		return spec.GetUserAccountJSON500Response(spec.ErrorResponse{
			Message: ErrInternalError,
		})
	}

	return spec.GetUserAccountJSON200Response(spec.BuscaUsuario{
		Usuario: &spec.Usuario{
			ID:        user.ID.String(),
			Email:     types.Email(user.Email),
			Nome:      user.Name,
			Cargo:     user.Role,
			UpdatedAt: user.UpdatedAt.UTC(),
			CreatedAt: user.CreatedAt.UTC(),
		},
	})
}

// Login user
// (POST /v1/users/login)
func (api *Handlers) PostLoginUser(w http.ResponseWriter, r *http.Request) *spec.Response {
	var payload spec.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		api.logger.Error("failed to decode user", zap.Error(err))
		return spec.PostLoginUserJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	if err := api.validator.Struct(payload); err != nil {
		api.logger.Error("failed to validate user", zap.Error(err))
		return spec.PostLoginUserJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	token, err := api.usersUsecase.LoginUser(usecase.LoginUserInput{
		Email:    string(payload.Email),
		Password: payload.Password,
	}, r.Context())
	if err != nil {
		api.logger.Error("failed to login user", zap.Error(err))
		return spec.PostLoginUserJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	return spec.PostLoginUserJSON200Response(spec.LoginRes{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		ExpiresIn:   &token.ExpiresIn,
	})
}

// Update user
// (PUT /v1/users/update)
func (api *Handlers) PutUpdateUser(w http.ResponseWriter, r *http.Request) *spec.Response {
	userID, err := GetUserIDFromContext(r.Context())
	if err != nil {
		return spec.PutUpdateUserJSON401Response(spec.ErrorResponse{
			Message: ErrNotAuthorized,
		})
	}

	var payload spec.AtualizarUsuario
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return spec.PutUpdateUserJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	err = api.usersUsecase.UpdateUser(userID, usecase.UpdateUserInput{
		Name:  payload.Nome,
		Email: (*string)(payload.Email),
	}, r.Context())
	if err != nil {
		return spec.PutUpdateUserJSON400Response(spec.ErrorResponse{
			Message: ErrBadRequest,
		})
	}

	return spec.PutUpdateUserJSON204Response(spec.Resp204{
		Message: "Usuário atualizado com sucesso",
	})
}

func getLevel(level string) spec.FormularioNivelDificuldade {
	switch level {
	case "low":
		return spec.FormularioNivelDificuldadeHigh
	case "medium":
		return spec.FormularioNivelDificuldadeMedium
	case "high":
		return spec.FormularioNivelDificuldadeHigh
	default:
		return spec.FormularioNivelDificuldadeLow
	}
}

func getClientType(clientType string) spec.ClienteTipoCliente {
	switch clientType {
	case "avulso":
		return spec.ClienteTipoClienteAvulso
	case "contrato":
		return spec.ClienteTipoClienteContrato
	default:
		return spec.ClienteTipoClienteAvulso
	}
}
