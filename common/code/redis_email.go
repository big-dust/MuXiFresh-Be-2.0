package code

func SetEmailCode(prefix string, key string, value string) error {
	return redisClient.Setex(prefix+key, value, EmailCodeExpired*60)
}

func VerifyEmailCode(prefix string, key string, value string) bool {
	val, err := redisClient.Get(prefix + key)
	if err != nil {
		return false
	}
	return value != "" && val == value
}
