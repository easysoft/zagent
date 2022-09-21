package repo

import (
	"github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/pkg/domain"
	_intUtils "github.com/easysoft/zv/pkg/lib/int"
)

type CommonRepo struct {
}

func (r CommonRepo) FindAssetByOs(osCategory consts.OsCategory, osType consts.OsType, osLang consts.OsLang,
	items []domain.VmAssert, inIds []uint) (assertIds []uint, found bool) {

	mp := map[int][]uint{}
	similarity := 0
	for _, item := range items {
		itemCategory := item.OsCategory
		itemType := item.OsType
		itemLang := item.OsLang

		if similarity < 3 && (itemCategory == osCategory && itemType == osType && itemLang == osLang) {
			if inIds == nil || _intUtils.FindUintInArr(item.ID, inIds) {
				similarity = 3
				mp[similarity] = append(mp[similarity], item.ID)
			}

		} else if similarity < 2 && (itemCategory == osCategory && itemType == osType) {
			if inIds == nil || _intUtils.FindUintInArr(item.ID, inIds) {
				similarity = 2
				mp[similarity] = append(mp[similarity], item.ID)
			}

		} else if similarity < 1 && (itemCategory == osCategory) {
			if inIds == nil || _intUtils.FindUintInArr(item.ID, inIds) {
				similarity = 1
				mp[similarity] = append(mp[similarity], item.ID)
			}

		}
	}

	for i := 3; i > 0; i-- { // order by similarity desc
		if mp[i] == nil {
			continue
		}

		for _, id := range mp[i] {
			assertIds = append(assertIds, id)
		}
	}
	if len(assertIds) > 0 {
		found = true
	}

	return
}
