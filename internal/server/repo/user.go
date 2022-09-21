package repo

import (
	"errors"
	"fmt"
	commDomain "github.com/easysoft/zv/internal/pkg/domain"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) NewUser() *model.User {
	return &model.User{}
}

func (r *UserRepo) Get(id uint) (po model.User, err error) {
	err = r.DB.Where("id = ?", id).First(&po).Error

	return
}

// GetUser get user
func (r *UserRepo) GetUser(search *commDomain.Search) (*model.User, error) {
	t := r.NewUser()
	err := r.Found(search).First(t).Error
	if !r.IsNotFound(err) {
		return t, err
	}
	return t, nil
}

// DeleteUser del user . if user's username is username ,can't del it.
func (r *UserRepo) DeleteUser(id uint) error {
	s := &commDomain.Search{
		Fields: []*commDomain.Filed{
			{
				Key:       "id",
				Condition: "=",
				Value:     id,
			},
		},
	}
	u, err := r.GetUser(s)
	if err != nil {
		return err
	}
	if u.Username == "username" {
		return errors.New(fmt.Sprintf("不能删除管理员 : %s \n ", u.Username))
	}

	if err := r.DB.Delete(u, id).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteUserByIdErr:%s \n ", err))
		return err
	}
	return nil
}

// GetAllUsers get all users
func (r *UserRepo) GetAllUsers(s *commDomain.Search) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64
	q := r.GetAll(&model.User{}, s)
	if err := q.Count(&count).Error; err != nil {
		return nil, count, err
	}
	q = q.Scopes(r.Paginate(s.Offset, s.Limit), r.Relation(s.Relations))
	if err := q.Find(&users).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllUserErr:%s \n ", err))
		return nil, count, err
	}
	return users, count, nil
}

func (r *UserRepo) UpdateRefreshToken(id uint, token string) {
	r.DB.Model(&model.User{}).
		Where("id=?", id).
		Updates(map[string]interface{}{"token": token, "token_updated_time": time.Now()})
}

func (r *UserRepo) GetByToken(token string) (model.User, error) {
	user := model.User{}
	err := r.DB.Model(&user).Where("token", token).First(&user).Error
	return user, err
}

func (r *UserRepo) UpdateUserDefaultProject(userId string, projectId int) {
	r.DB.Model(&model.User{}).
		Where("id=?", userId).
		Updates(map[string]interface{}{"project_id": projectId})
}
