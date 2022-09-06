package initService

import (
	"fmt"
	commDomain "github.com/easysoft/zv/internal/comm/domain"
	"github.com/easysoft/zv/internal/server/conf"
	"github.com/easysoft/zv/internal/server/model"
	"github.com/easysoft/zv/internal/server/repo"
	serverService "github.com/easysoft/zv/internal/server/service"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	serverRes "github.com/easysoft/zv/res/server"
	logger "github.com/sirupsen/logrus"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/azumads/faker"
	"github.com/jinzhu/configor"
)

type SeederService struct {
	InitService *InitService               `inject:""`
	RoleService *serverService.RoleService `inject:""`
	UserService *serverService.UserService `inject:""`

	CommonRepo *repo.BaseRepo `inject:""`
	UserRepo   *repo.UserRepo `inject:""`
	RoleRepo   *repo.RoleRepo `inject:""`
	PermRepo   *repo.PermRepo `inject:""`
}

func NewSeeder() *SeederService {
	seeder := &SeederService{}
	seeder.init()
	return seeder
}

var (
	Seeds = struct {
		Perms []struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayname"`
			Description string `json:"description"`
			Act         string `json:"act"`
		}
	}{}

	Fake *faker.Faker
)

func (s *SeederService) init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	exeDir := _fileUtils.GetExeDir()
	pth := ""
	if _commonUtils.IsRelease() {
		pth = filepath.Join(exeDir, "perms.yml")
		if !_fileUtils.FileExist(pth) {
			bytes, _ := serverRes.Asset("res/server/perms.yml")
			_fileUtils.WriteFile(pth, string(bytes))
		}
	} else {
		pth = filepath.Join(exeDir, "cmd", "server", "perms.yml")
	}
	filePaths, _ := filepath.Glob(pth)

	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("数据填充YML文件路径：%+v\n", filePaths))
	}
	if err := configor.Load(&Seeds, filePaths...); err != nil {
		logger.Println(err)
	}
}

func (s *SeederService) AddPerms() {
	fmt.Println(fmt.Sprintf("开始填充权限！！"))
	s.CreatePerms()
	s.CreateAdminRole()
	s.CreateAdminUser()
	fmt.Println(fmt.Sprintf("权限填充完成！！"))
}

// CreatePerms 新建权限
func (s *SeederService) CreatePerms() {
	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("填充权限：%+v\n", Seeds))
	}

	for _, m := range Seeds.Perms {
		search := &commDomain.Search{
			Fields: []*commDomain.Filed{
				{
					Key:       "name",
					Condition: "=",
					Value:     m.Name,
				}, {
					Key:       "act",
					Condition: "=",
					Value:     m.Act,
				},
			},
		}
		perm, err := s.PermRepo.GetPermission(search)
		if err == nil {
			if perm.ID == 0 {
				perm = &model.Permission{
					Name:        m.Name,
					DisplayName: m.DisplayName,
					Description: m.Description,
					Act:         m.Act,
				}
				if err := s.PermRepo.CreatePermission(perm); err != nil {
					logger.Println(fmt.Sprintf("权限填充错误：%+v\n", err))
				}
			}
		}
	}
}

// CreateAdminRole 新建管理角色
func (s *SeederService) CreateAdminRole() {
	search := &commDomain.Search{
		Fields: []*commDomain.Filed{
			{
				Key:       "name",
				Condition: "=",
				Value:     serverConf.Inst.Admin.RoleName,
			},
		},
	}
	role, err := s.RoleRepo.GetRole(search)
	var permIds []uint
	ss := &commDomain.Search{
		Limit:  -1,
		Offset: -1,
	}
	perms, _, err := s.PermRepo.GetAllPermissions(ss)
	if serverConf.Inst.Debug {
		if err != nil {
			fmt.Println(fmt.Sprintf("权限获取失败：%+v\n", err))
		}
	}

	for _, perm := range perms {
		permIds = append(permIds, perm.ID)
	}
	role.PermIds = permIds

	if err == nil {
		if role.ID == 0 {
			role = &model.Role{
				Name:        serverConf.Inst.Admin.RoleName,
				DisplayName: serverConf.Inst.Admin.RoleDisplayName,
				Description: serverConf.Inst.Admin.RoleDisplayName,
			}
			role.PermIds = permIds
			if err := s.RoleService.CreateRole(role); err != nil {
				logger.Println(fmt.Sprintf("管理角色填充错误：%+v\n", err))
			}
		} else {
			if err := s.RoleService.UpdateRole(role.ID, role); err != nil {
				logger.Println(fmt.Sprintf("管理角色填充错误：%+v\n", err))
			}
		}
	}
	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("填充角色数据：%+v\n", role))
		fmt.Println(fmt.Sprintf("填充角色权限：%+v\n", role.PermIds))
	}

}

// CreateAdminUser 新建管理员
func (s *SeederService) CreateAdminUser() {
	search := &commDomain.Search{
		Fields: []*commDomain.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     serverConf.Inst.Admin.UserName,
			},
		},
	}
	admin, err := s.UserRepo.GetUser(search)
	if err != nil {
		fmt.Println(fmt.Sprintf("GetByIdent admin error：%+v\n", err))
	}

	var roleIds []uint
	ss := &commDomain.Search{
		Limit:  -1,
		Offset: -1,
	}
	roles, _, err := s.RoleRepo.GetAllRoles(ss)
	if serverConf.Inst.Debug {
		if err != nil {
			fmt.Println(fmt.Sprintf("角色获取失败：%+v\n", err))
		}
	}

	for _, role := range roles {
		roleIds = append(roleIds, role.ID)
	}
	admin.RoleIds = roleIds

	if admin.ID == 0 {
		admin = &model.User{
			Username: serverConf.Inst.Admin.UserName,
			Name:     serverConf.Inst.Admin.Name,
			Password: serverConf.Inst.Admin.Password,
			Avatar:   "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132",
			Intro:    "檀越",
		}

		admin.RoleIds = roleIds
		if err := s.UserService.CreateUser(admin); err != nil {
			logger.Println(fmt.Sprintf("管理员填充错误：%+v\n", err))
		}
	} else {
		admin.Password = serverConf.Inst.Admin.Password
		if err := s.UserService.UpdateUserById(admin.ID, admin); err != nil {
			logger.Println(fmt.Sprintf("管理员填充错误：%+v\n", err))
		}
	}

	if serverConf.Inst.Debug {
		fmt.Println(fmt.Sprintf("管理员密码：%s\n", serverConf.Inst.Admin.Password))
		fmt.Println(fmt.Sprintf("填充管理员数据：%+v", admin))
	}
}
