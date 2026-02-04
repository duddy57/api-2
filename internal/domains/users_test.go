package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUser_Validate tests the Validate method with various scenarios
func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid user - all fields correct",
			user: User{
				Name:     "John Doe",
				Email:    "test@example.com",
				Role:     "admin",
				Password: []byte("valid value"),
			},
			wantErr: false,
		},
		{
			name: "invalid user - empty name",
			user: User{
				Name:     "",
				Email:    "test@example.com",
				Role:     "admin",
				Password: []byte("valid value"),
			},
			wantErr: true,
			errMsg:  "name",
		},
		{
			name: "invalid user - empty email",
			user: User{
				Name:     "John Doe",
				Email:    "",
				Role:     "tecnico",
				Password: []byte("valid value"),
			},
			wantErr: true,
			errMsg:  "email",
		},
		{
			name: "invalid user - empty role",
			user: User{
				Name:     "John Doe",
				Email:    "test@example.com",
				Role:     "",
				Password: []byte("valid value"),
			},
			wantErr: true,
			errMsg:  "role",
		},
		{
			name: "invalid user - empty password",
			user: User{
				Name:     "John Doe",
				Email:    "test@example.com",
				Role:     "tecnico",
				Password: []byte(""),
			},
			wantErr: true,
			errMsg:  "password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestUser_Initialization tests entity field initialization
func TestUser_Initialization(t *testing.T) {
	user := &User{
		Name:     "John Doe",
		Email:    "test@example.com",
		Password: []byte("valid value"),
	}

	assert.Equal(t, "John Doe", user.Name, "Name should be set correctly")
	assert.Equal(t, "test@example.com", user.Email, "Email should be set correctly")
	assert.Equal(t, []byte("valid value"), user.Password, "Password should be set correctly")
}

// TestUser_Name_EdgeCases tests edge cases for Name field
func TestUser_Name_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid value", value: "Valid Name", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Name:     tt.value,
				Email:    "test@example.com",
				Role:     "admin",
				Password: []byte("valid value"),
			}

			err := user.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestUser_Email_EdgeCases tests edge cases for Email field
func TestUser_Email_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "empty string", value: "", wantErr: true},
		{name: "valid email", value: "test@example.com", wantErr: false},
		{name: "invalid email format", value: "notanemail", wantErr: true},
		{name: "email without @", value: "testexample.com", wantErr: true},
		{name: "email without domain", value: "test@", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Name:     "John Doe",
				Email:    tt.value,
				Role:     "admin",
				Password: []byte("valid value"),
			}

			err := user.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestUser_Password_EdgeCases tests edge cases for Password field
func TestUser_Password_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{name: "valid value", value: "Valid Name", wantErr: false},
		{name: "empty string", value: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Name:     "John Doe",
				Email:    "test@example.com",
				Role:     "admin",
				Password: []byte(tt.value),
			}

			err := user.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
