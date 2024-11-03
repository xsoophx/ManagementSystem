package persist

import (
	"ManagementSystem/persist/models"
	sm "ManagementSystem/services/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{db: GetDB()}
}

func (dao *UserDao) CreateUser(user *sm.User) error {
	return dao.db.Create(dao.toUserDbo(user)).Error
}

func (dao *UserDao) GetUser(id uuid.UUID) (*sm.User, error) {
	var user models.UserDbo
	if err := dao.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return dao.toUser(&user), nil
}

func (dao *UserDao) GetAllUsers() ([]sm.User, error) {
	var userDbos []models.UserDbo
	if err := dao.db.Find(&userDbos).Error; err != nil {
		return nil, err
	}

	var users []sm.User
	for _, userDbo := range userDbos {
		users = append(users, *dao.toUser(&userDbo))
	}
	return users, nil
}

func (dao *UserDao) UpdateUser(user *sm.User) error {
	return dao.db.Save(dao.toUserDbo(user)).Error
}

func (dao *UserDao) DeleteUser(id uuid.UUID) error {
	return dao.db.Delete(&models.UserDbo{}, id).Error
}

func (dao *UserDao) toUserDbo(user *sm.User) *models.UserDbo {
	return &models.UserDbo{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (dao *UserDao) toUser(user *models.UserDbo) *sm.User {
	return &sm.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
