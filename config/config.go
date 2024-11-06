package config

import "github.com/hibiken/asynq"

// RedisConfig trả về cấu hình Redis cho client và server worker
func RedisConfig() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}
}
