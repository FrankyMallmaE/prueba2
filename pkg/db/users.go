package db

import (
	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
	"gorm.io/gorm"
)

type UsersRepo struct {
	DB                 *gorm.DB
	CreateUserObserver map[string]chan *models.User
	UpdateUserObserver map[string]chan *models.User
	DeleteUserObserver map[string]chan *models.User
}

func (m *UsersRepo) GetUsers() ([]*models.User, error) {

	var users []*models.User

	query := m.DB.Model(&models.User{}).
		Order("id ASC").
		Find(&users)
	if err := query.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UsersRepo) GetUser(id string) (*models.User, error) {

	var user *models.User

	query := m.DB.Model(&models.User{}).
		Where("id = ?", id).
		Find(&user)
	if err := query.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UsersRepo) CreateUser(user *models.User) (*models.User, error) {

	query := m.DB.Create(user)
	if err := query.Error; err != nil {
		return nil, err
	}

	for _, observer := range m.CreateUserObserver {
		observer <- user
	}

	return user, nil
}

func (m *UsersRepo) UpdateUser(user *models.User) (*models.User, error) {

	query := m.DB.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(user)
	if err := query.Error; err != nil {
		return nil, err
	}

	for _, observer := range m.UpdateUserObserver {
		observer <- user
	}

	return user, nil
}

func (m *UsersRepo) DeleteUser(user *models.User) (*models.User, error) {

	query := m.DB.Unscoped().
		Delete(&user)
	if err := query.Error; err != nil {
		return nil, err
	}

	for _, observer := range m.DeleteUserObserver {
		observer <- user
	}

	return user, nil
}
