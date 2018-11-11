package model

// AppSetting launch and configure appsetting on app start
type AppSetting struct {
	Port       string  `json:"port"`
	Env        string  `json:"env"`
	Connection string  `json:"dbConnection"`
	DriverName string `json:"driverName"`
	EnableCors bool    `json:"enablecors"`
	Logging    Logging `json:"logging"`
}

type Logging struct {
	Filename string `json:"filepath"`
	MaxSize  int    `json:"maxSize"`
}
