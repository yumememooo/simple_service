package config

type ConfigurationStruct struct {
	Service  ServiceInfo
	Logger   LoggerInfo
	Database DatabaseInfo
	Services map[string]ServicesInfo
}

type ServiceInfo struct {
	Port       int
	StartupMsg string
}

type LoggerInfo struct {
	Level string
	File  string
}

type DatabaseInfo struct {
	Host     string
	Port     int
	Username string
	Password string
}

// func (d DatabaseInfo) Url() string {
// 	url := fmt.Sprintf("mongodb://%s:%d", d.Host, d.Port)
// 	return url
// }

type ServicesInfo struct {
	Host     string
	Port     int
	Protocol string
}

// func (c ClientInfo) Url(apiRoute string) string {
// 	url := fmt.Sprintf("%s://%s:%d%s", c.Protocol, c.Host, c.Port, apiRoute)
// 	return url
// }
