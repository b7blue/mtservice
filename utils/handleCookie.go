package utils

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 根据uid生成mtstempid，并存入redis，过期时间为6h（暂定）
// 返回mtstempid，假如为空就说明出错
func NewCookie(uid int) string {
	lastlogin := time.Now()
	salt := lastlogin.String()
	mtstempid := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d", uid)+salt)))

	// 存入radis
	conn := cookiePool.Get()
	_, err := conn.Do("set", mtstempid, uid, "EX", "21600")
	conn.Close()
	if err != nil {
		return ""
	}
	return mtstempid
}

// 后端根据用户cookie中的mtstempid查redis获取uid
// 假如radis中不存在key（已超时或者是非法的mtstempid），
func TryCookie(mtstempid string) (uid int) {
	conn := cookiePool.Get()
	uid, _ = redis.Int(conn.Do("Get", mtstempid))
	conn.Close()

	return uid
}
