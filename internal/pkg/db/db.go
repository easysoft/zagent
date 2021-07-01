package _db

import (
	"fmt"
	"github.com/easysoft/zagent/internal/pkg/lib/file"
	"github.com/easysoft/zagent/internal/pkg/lib/log"
	serverConf "github.com/easysoft/zagent/internal/server/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"path/filepath"
	"strings"
	"time"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	inst *Instance

	FlagVarDBFile string
)

func GetInst() *Instance {
	return inst
}

func InitDB(model string) {
	var dialector gorm.Dialector

	if model == "agent" || serverConf.Config.DB.Adapter == "sqlite3" {
		conn := DBFile(model)
		dialector = sqlite.Open(conn)

	} else if serverConf.Config.DB.Adapter == "mysql" {
		conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
			serverConf.Config.DB.User, serverConf.Config.DB.Password, serverConf.Config.DB.Host, serverConf.Config.DB.Port, serverConf.Config.DB.Name)
		dialector = mysql.Open(conn)

	} else if serverConf.Config.DB.Adapter == "postgres" {
		conn := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
			serverConf.Config.DB.User, serverConf.Config.DB.Password, serverConf.Config.DB.Host, serverConf.Config.DB.Name)
		dialector = postgres.Open(conn)

	} else {
		_logUtils.Info("not supported database adapter")
	}

	DB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   serverConf.Config.DB.Prefix,
			SingularTable: false,
		},
	})

	if err != nil {
		_logUtils.Info(err.Error())
	}

	_ = DB.Use(
		dbresolver.Register(
			dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	DB.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})

	inst = &Instance{}
	inst.db = DB
	inst.config = &serverConf.Config.DB
}

func (*Instance) DB() *gorm.DB {
	return inst.db
}

type Instance struct {
	config *serverConf.DBConfig
	db     *gorm.DB
}

func (i *Instance) Close() error {
	if i.db != nil {
		sqlDB, _ := i.db.DB()
		return sqlDB.Close()
	}
	return nil
}

func DBFile(mode string) string {
	if mode != "agent" {
		mode = serverConf.Config.DB.Name
	}

	path := filepath.Join(_fileUtils.GetExeDir(), strings.ToLower(mode+".db"))
	return path
}
