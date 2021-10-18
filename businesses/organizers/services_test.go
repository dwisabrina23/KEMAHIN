package organizers_test

import (
	"kemahin/app/middlewares"
	"kemahin/businesses"
	"kemahin/businesses/organizers"
	_orgsMock "kemahin/businesses/organizers/mocks"
	"kemahin/helpers/encrypt"
	_encryptMock "kemahin/helpers/encrypt/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUserRepo   _orgsMock.Repository
	mockEncrypt    _encryptMock.Helper
	orgService     organizers.Service
	domainTest     organizers.Domain
	hashedPassword string
)

func TestMain(m *testing.M) {
	orgService = organizers.NewService(&mockUserRepo, &middlewares.ConfigJWT{})
	hashedPassword = encrypt.HashAndSalt([]byte("password"))
	domainTest = organizers.Domain{
		Username: "ukmtest",
		Pasword:  hashedPassword,
		Name:     "Dwi Sabrina",
		Phone:    "085123123",
	}

	m.Run()
}

func TestUserRegister(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.Anything).Return(organizers.Domain{}, nil).Once()
		mockUserRepo.On("Register", mock.Anything).Return(domainTest, nil).Once()

		inputUser := organizers.Domain{
			Username: "ukmtest",
			Pasword:  hashedPassword,
			Name:     "Dwi Sabrina",
			Phone:    "085123123",
		}

		resp, err := orgService.Register(&inputUser)

		assert.Nil(t, err)
		assert.Equal(t, domainTest, resp)
	})
	t.Run("Invalid Test | Unregistered org", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.Anything).Return(organizers.Domain{}, assert.AnError).Once()

		inputUser := organizers.Domain{
			Username: "ukmtest",
			Pasword:  hashedPassword,
			Name:     "Dwi Sabrina",
			Phone:    "085123123",
		}

		resp, err := orgService.Register(&inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, resp)
	})
	t.Run("Invalid Test | Duplicate Data", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.Anything).Return(domainTest, nil).Once()
		// mockUserRepo.On("Register", mock.Anything).Return(organizers.Domain{}, businesses.ErrDuplicateData).Once()
		inputUser := organizers.Domain{
			Username: "ukmtest",
			Pasword:  "password",
			Name:     "Dwi Sabrina",
			Phone:    "085123123",
		}

		resp, err := orgService.Register(&inputUser)

		assert.Equal(t, organizers.Domain{}, resp)
		assert.Equal(t, err, businesses.ErrDuplicateData)
	})
	t.Run("Invalid Test | Invalid Input Data", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(organizers.Domain{}, nil).Once()
		mockUserRepo.On("Register", mock.Anything).Return(organizers.Domain{}, assert.AnError).Once()
		inputUser := organizers.Domain{
			Username: "",
			Pasword:  "password",
			Name:     "test",
			Phone:    "",
		}

		resp, err := orgService.Register(&inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(domainTest, nil).Once()

		inputUser := organizers.Domain{
			Username: "ukmtest",
			Pasword:  "password",
		}

		resp, err := orgService.Login(inputUser.Username, inputUser.Pasword)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("Invalid Test | empty username and password", func(t *testing.T) {
		// mockUserRepo.On("GetByNIM", mock.AnythingOfType("string")).Return(organizers.Domain, nil).Once()
		inputUser := organizers.Domain{
			Username: "",
			Pasword:  "",
		}

		resp, err := orgService.Login(inputUser.Username, inputUser.Pasword)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})

	t.Run("Invalid Test | wrong username", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(organizers.Domain{}, assert.AnError).Once()

		inputUser := organizers.Domain{
			Username: "190030552",
			Pasword:  "password",
		}

		resp, err := orgService.Login(inputUser.Username, inputUser.Pasword)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})

	t.Run("Invalid Test | wrong password", func(t *testing.T) {
		mockUserRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(domainTest, nil).Once()
		mockEncrypt.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		inputUser := organizers.Domain{
			Username: "ukmtest",
			Pasword:  "passwordsalah",
		}

		resp, err := orgService.Login(inputUser.Username, inputUser.Pasword)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.AnythingOfType("int")).Return(domainTest, nil).Once()

		resp, err := orgService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, domainTest, resp)
	})
	t.Run("Invalid Test | org not found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.AnythingOfType("int")).Return(organizers.Domain{}, assert.AnError).Once()

		resp, err := orgService.GetByID(1)

		assert.NotNil(t, err)
		assert.Equal(t, organizers.Domain{}, resp)
	})
}
