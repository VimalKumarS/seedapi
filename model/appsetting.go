package model

// AppSetting launch and configure appsetting on app start
type AppSetting struct {
	Port       string `json:"port"`
	Env        string `json:"env"`
	Connection string `json:"dbConnection"`
}
