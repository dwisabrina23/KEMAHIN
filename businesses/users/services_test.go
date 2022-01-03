package users_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kemahin/app/middlewares"
	"kemahin/businesses"
	"kemahin/businesses/users"
	_usersMock "kemahin/businesses/users/mocks"
	"kemahin/helpers/encrypt"
	_encryptMock "kemahin/helpers/encrypt/mocks"
	"testing"
	"time"
)

var (
	mockUserRepo     _usersMock.Repository
	mockEncrypt      _encryptMock.Helper
	userService      users.Service
	domainTest       users.Domain
	hashedPassword   string
	updateDomainTest users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewService(&mockUserRepo, &middlewares.ConfigJWT{})
	hashedPassword = encrypt.HashAndSalt([]byte("password"))
	domainTest = users.Domain{
		NIM:     "190030551",
		Pasword: hashedPassword,
		Name:    "Dwi Sabrina",
		Prodi:   "Sistem Informasi",
		Phone:   "085123123",
		Email:   "dwi@gmail.com",
		Role:    1,
	}
	updateDomainTest = users.Domain{
		Id:        1,
		NIM:       "190030550",
		Pasword:   hashedPassword,
		Name:      "test",
		Prodi:     "Sistem Informasi",
		Phone:     "085123123",
		Email:     "dwi@gmail.com",
		Role:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.Run()
}

func TestUserRegister(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.Anything).Return(users.Domain{}, nil).Once()
		mockUserRepo.On("Register", mock.Anything).Return(domainTest, nil).Once()

		inputUser := users.Domain{
			NIM:     "190030551",
			Pasword: "password",
			Name:    "Dwi Sabrina",
			Prodi:   "Sistem Informasi",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    1,
		}

		resp, err := userService.Register(&inputUser)

		assert.Nil(t, err)
		assert.Equal(t, domainTest, resp)
	})
	t.Run("Invalid Test | Unregistered user", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			NIM:     "190030551",
			Pasword: "password",
			Name:    "Dwi Sabrina",
			Prodi:   "Sistem Informasi",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    1,
		}

		resp, err := userService.Register(&inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, resp)
	})
	t.Run("Invalid Test | Duplicate Data", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.Anything).Return(domainTest, nil).Once()
		inputUser := users.Domain{
			NIM:     "190030551",
			Pasword: "password",
			Name:    "Dwi Sabrina",
			Prodi:   "Sistem Informasi",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    1,
		}

		resp, err := userService.Register(&inputUser)

		assert.Equal(t, users.Domain{}, resp)
		assert.Equal(t, err, businesses.ErrDuplicateData)
	})
	t.Run("Invalid Test | Invalid Input  Data", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.Anything).Return(users.Domain{}, nil).Once()
		mockUserRepo.On("Register", mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			NIM:     "",
			Pasword: "",
			Name:    "",
			Prodi:   "",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    -1,
		}

		resp, err := userService.Register(&inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, resp)
	})

}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.AnythingOfType("string")).Return(domainTest, nil).Once()

		inputUser := users.Domain{
			NIM:     "190030551",
			Pasword: "password",
		}

		resp, err := userService.Login(inputUser.NIM, inputUser.Pasword)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("Invalid Test | empty username and password", func(t *testing.T) {
		// mockUserRepo.On("GetByNIM", mock.AnythingOfType("string")).Return(users.Domain, nil).Once()
		inputUser := users.Domain{
			NIM:     "",
			Pasword: "",
		}

		resp, err := userService.Login(inputUser.NIM, inputUser.Pasword)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})

	t.Run("Invalid Test | wrong nim", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.AnythingOfType("string")).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			NIM:     "190030552",
			Pasword: "password",
		}

		resp, err := userService.Login(inputUser.NIM, inputUser.Pasword)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})

	t.Run("Invalid Test | wrong password", func(t *testing.T) {
		mockUserRepo.On("GetByNIM", mock.AnythingOfType("string")).Return(domainTest, nil).Once()
		mockEncrypt.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		inputUser := users.Domain{
			NIM:     "190030551",
			Pasword: "passwordsalah",
		}

		resp, err := userService.Login(inputUser.NIM, inputUser.Pasword)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.AnythingOfType("int")).Return(domainTest, nil).Once()

		resp, err := userService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, domainTest, resp)
	})
	t.Run("Invalid Test | user not found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.AnythingOfType("int")).Return(users.Domain{}, assert.AnError).Once()

		resp, err := userService.GetByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, users.Domain{}, resp)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything).Return(users.Domain{}, nil).Once()
		mockUserRepo.On("Update", mock.Anything).Return(updateDomainTest, nil).Once()

		inputUser := users.Domain{
			Pasword: "password",
			Prodi:   "Sistem Informasi",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    1,
		}
		resp, err := userService.Update(inputUser)

		assert.Nil(t, err)
		assert.Equal(t, updateDomainTest, resp)
	})
	t.Run("Invalid Test | invalid id", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.AnythingOfType("int")).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			Pasword: "password",
			Prodi:   "Sistem Informasi",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    1,
		}

		resp, err := userService.Update(inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, resp)
	})

	t.Run("Invalid Test | Invalid Input  Data", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything).Return(users.Domain{}, nil).Once()
		mockUserRepo.On("Update", mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		inputUser := users.Domain{
			NIM:     "",
			Pasword: "",
			Name:    "",
			Prodi:   "",
			Phone:   "085123123",
			Email:   "dwi@gmail.com",
			Role:    -1,
		}

		resp, err := userService.Update(inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, resp)
	})
}
