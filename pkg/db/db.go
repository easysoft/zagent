package _db

import (
	agentConf "github.com/easysoft/zv/internal/pkg/conf"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
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

func InitDB(mode string) {
	var dialector gorm.Dialector

	if mode == "agent" {
		conn := DBFile(mode)
		dialector = sqlite.Open(conn)

	} else {
		_logUtils.Info("not supported database adapter")
	}

	prefix := ""
	if mode == "agent" {
		prefix = agentConf.Inst.DB.Prefix
	}

	DB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
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
}

func (*Instance) DB() *gorm.DB {
	return inst.db
}

type Instance struct {
	//config *serverConf.DBConfig
	db *gorm.DB
}

func (i *Instance) Close() error {
	if i.db != nil {
		sqlDB, _ := i.db.DB()
		return sqlDB.Close()
	}
	return nil
}

func DBFile(mode string) string {
	dbName := "agent"

	path := filepath.Join(_fileUtils.GetExeDir(), strings.ToLower(dbName+".db"))
	return path
}
