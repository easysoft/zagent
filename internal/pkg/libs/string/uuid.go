package _stringUtils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	uid := uuid.NewV4().String()
	return strings.Replace(uid, "-", "", -1)
}
