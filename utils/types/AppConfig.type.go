package types

type AppConfig struct {
	Env               string // dev|deploy
	Log_enabled       string // true|false
	App_port          string // 4020
	App_domain        string
	App_api_prefix_v1 string
	Database_host     string
	Database_port     string
	Database_name     string
	Database_username string
	Database_password string
}
