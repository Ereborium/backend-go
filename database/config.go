package database

import "time"

type Config struct {
	Address     string
	Username    string
	Password    string
	Database    string
	DialTimeout time.Duration
}
