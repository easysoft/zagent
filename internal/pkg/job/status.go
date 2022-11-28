package job

import (
	"sync"
	"time"
)

var (
	TaskStatus sync.Map
)

func GetSpeed(startTime time.Time, sumKByte int64) (ret float64) {
	sec := time.Now().Unix() - startTime.Unix()
	if sec == 0 {
		sec = 1
	}

	ret = float64(sumKByte) / float64(sec)

	return
}

func SaveTaskStatus(cache *sync.Map, id uint, rate, speed float64) {
	valObj, ok := cache.Load(id)

	if !ok {
		valObj = map[string]float64{
			"rate":  rate,
			"speed": speed,
		}
		cache.Store(id, valObj)

		return
	}

	valMap := valObj.(map[string]float64)

	if rate > valMap["rate"] {
		valMap["rate"] = rate
	}
	if speed > 0 {
		valMap["speed"] = speed
	}

	cache.Store(id, valMap)

	return
}

func GetTaskStatus(cache sync.Map, id uint) (rate, speed float64) {
	valObj, ok := cache.Load(id)

	if !ok {
		return
	}

	valMap := valObj.(map[string]float64)

	rate = valMap["rate"]
	speed = valMap["speed"]

	return
}
