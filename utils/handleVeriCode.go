package utils

import "github.com/garyburd/redigo/redis"

// 发送验证码之前检查，十分钟之内是否已经发送过了
func AlreadyVeri(email string) bool {
	conn := veriCodePool.Get()
	is_email_exit, _ := redis.Bool(conn.Do("EXISTS", email))
	conn.Close()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	return is_email_exit
}

// 验证码十分钟过期
func SetVeriCode(email, varicode string) error {
	//通过连接池获得连接
	conn := veriCodePool.Get()
	//使用连接操作数据
	_, err := conn.Do("set", email, varicode, "EX", "600")
	conn.Close()
	return err
}

// 假如正常获得验证码，说明发送了验证码而且没有过期
// 假如没有获得验证码，说明根本没发送验证码或者验证码过期了
func GetVeriCode(email string) string {
	conn := veriCodePool.Get()
	varicode, _ := redis.String(conn.Do("Get", email))
	conn.Close()
	return varicode
}

func DelVeriCode(email string) error {
	conn := veriCodePool.Get()
	_, err := conn.Do("DEL", email)
	conn.Close()
	return err
}
