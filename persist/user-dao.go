package persist

import (
	"ManagementSystem/persist/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{db: GetDB()}
}

func (dao *UserDao) CreateUser(user *models.UserDbo) error {
	return dao.db.Create(user).Error
}

func (dao *UserDao) GetUser(id uuid.UUID) (*models.UserDbo, error) {
	var user models.UserDbo
	if err := dao.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDao) GetAllUsers() ([]models.UserDbo, error) {
	var users []models.UserDbo
	if err := dao.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *UserDao) UpdateUser(user *models.UserDbo) error {
	return dao.db.Save(user).Error
}

func (dao *UserDao) DeleteUser(id uuid.UUID) error {
	return dao.db.Delete(&models.UserDbo{}, id).Error
}
