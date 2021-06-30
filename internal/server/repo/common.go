package repo

import (
	commConst "github.com/easysoft/zagent/internal/comm/const"
	commDomain "github.com/easysoft/zagent/internal/comm/domain"
	_intUtils "github.com/easysoft/zagent/internal/pkg/lib/int"
)

type CommonRepo struct {
}

func (r CommonRepo) FindAssetByOs(osCategory commConst.OsCategory, osType commConst.OsType, osLang commConst.OsLang,
	items []commDomain.VmAssert, backingIds []int) (assertIds []int, found bool) {

	mp := map[int][]int{}
	similarity := 0
	for _, item := range items {
		itemCategory := item.OsCategory
		itemType := item.OsType
		itemLang := item.OsLang

		if similarity < 3 && (itemCategory == osCategory && itemType == osType && itemLang == osLang) &&
			_intUtils.FindInArr(int(item.ID), backingIds) {

			similarity = 3
			mp[similarity] = append(mp[similarity], int(item.ID))
		} else if similarity < 2 && (itemCategory == osCategory && itemType == osType) &&
			_intUtils.FindInArr(int(item.ID), backingIds) {

			similarity = 2
			mp[similarity] = append(mp[similarity], int(item.ID))
		} else if similarity < 1 && (itemCategory == osCategory) &&
			_intUtils.FindInArr(int(item.ID), backingIds) {

			similarity = 1
			mp[similarity] = append(mp[similarity], int(item.ID))
		}
	}

	for i := 3; i > 0; i-- { // order by similarity desc
		if mp[i] == nil {
			continue
		}

		for _, id := range assertIds {
			assertIds = append(assertIds, id)
		}
	}
	if len(assertIds) > 0 {
		found = true
	}

	return
}
