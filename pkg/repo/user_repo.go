package repo

import (
	"errors"

	"github.com/sekthor/singing-city/pkg/model"
	"gorm.io/gorm"
)

var (
	ErrorEmailExists = errors.New("email already exists")
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (r *UserRepo) Save(user model.User) (model.User, error) {
	result := r.db.Save(&user)
	return user, result.Error
}

func (r *UserRepo) Create(user model.User) (model.User, error) {
	result := r.db.Create(&user)

	// TODO: actually check the nature of the error
	// for now we just assume this is why it fails
	if result.Error != nil {
		return user, ErrorEmailExists
	}

	return user, result.Error
}

func (r *UserRepo) FetchById(id int) (model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *UserRepo) FetchByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *UserRepo) FetchAll() []model.User {
	var users []model.User
	_ = r.db.Find(&users)
	return users
}

func (r *UserRepo) DeleteById(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *UserRepo) Exists(id int) bool {
	var user model.User
	result := r.db.First(&user, id)
	return result.Error == nil && int(user.ID) == id
}

func (r *UserRepo) SaveOmitPassword(user model.User) (model.User, error) {
	result := r.db.Omit("Password").Save(&user)
	return user, result.Error
}

func (r *UserRepo) CreateInvite(i model.Invite) (model.Invite, error) {
	result := r.db.Create(&i)
	return i, result.Error
}

func (r *UserRepo) DeleteInviteById(id string) error {
	result := r.db.Where("invite = ?", id).Delete(&model.Invite{})
	return result.Error
}

func (r *UserRepo) FetchAllInvites() []model.Invite {
	var invs []model.Invite
	_ = r.db.Find(&invs)
	return invs
}

func (r *UserRepo) FetchInviteById(id string) (model.Invite, error) {
	var inv model.Invite
	result := r.db.Where("invite = ?", id).First(&inv)
	return inv, result.Error
}

func (r *UserRepo) CreatePasswordResetRequest(pr model.PasswordReset) (model.PasswordReset, error) {
	result := r.db.Create(&pr)
	return pr, result.Error
}

func (r *UserRepo) FetchPasswordResetRequestByCode(code string) (model.PasswordReset, error) {
	var req model.PasswordReset
	result := r.db.Where("code = ?", code).First(&req)
	return req, result.Error
}

func (r *UserRepo) DeletePasswordResetRequestsByUserID(userId uint) error {
	result := r.db.Where("user_id = ?", userId).Delete(&model.PasswordReset{})
	return result.Error
}

func (r *UserRepo) SetNewPasswordForUser(userId uint, passwordHash string) error {
	return r.db.Model(&model.User{}).
		Where("id = ?", userId).
		Update("password", passwordHash).Error
}
