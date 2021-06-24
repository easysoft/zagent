package jwt

import (
	bizConst "github.com/easysoft/zagent/internal/server/biz/const"
	"github.com/easysoft/zagent/internal/server/cfg"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	SessionID = "Utl_SessionID"
	CredKey   = "Utl_Credentials"
	session   = sessions.New(sessions.Config{Cookie: SessionID})
)

func Get(ctx iris.Context, key string) (obj interface{}) {
	sess := session.Start(ctx)
	obj = sess.Get(key)
	return
}
func Set(ctx iris.Context, key string, obj interface{}) {
	sess := session.Start(ctx)
	sess.Set(key, obj)
}

func GetCredentials(ctx iris.Context) (cred *bizConst.UserCredentials) {
	if serverConf.Config.Redis.Enable {
		credObj := ctx.Values().Get("sess")
		if credObj == nil {
			return
		}
		cred = credObj.(*bizConst.UserCredentials)
	} else {
		sess := session.Start(ctx)
		credObj := sess.Get(CredKey)
		if credObj == nil {
			return
		}

		cred = credObj.(*bizConst.UserCredentials)
	}

	return
}

func SaveCredentials(ctx iris.Context, cred *bizConst.UserCredentials) {
	sess := session.Start(ctx)
	sess.Set(CredKey, cred)
}
func RemoveCredentials(ctx iris.Context) {
	sess := session.Start(ctx)
	sess.Delete(CredKey)
}
