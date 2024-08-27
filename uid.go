package myuid

import (
	"math/rand"
	"time"
)

func RandomIdGenerator() int64 {
	nanscnd := time.Now().UnixNano() % 10000000
	time.Sleep(1 * time.Millisecond)
	nanscnd2 := time.Now().UnixNano() % 10000000
	uid := nanscnd2 * nanscnd

	uid += rand.Int63n(1000000000000)
	number := lenLoop10(uid)
	if number < 14 {
		uid = uid * int64((14-number)*10)
	}
	if number > 15 {
		uid = uid / int64((number-14)*10)
	}
	return uid

}
func lenLoop10(i int64) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for int64(x) <= i {
		x *= 10
		count++
	}
	return count
}
