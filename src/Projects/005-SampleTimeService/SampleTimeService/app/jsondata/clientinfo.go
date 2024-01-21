package jsondata

type ClientInfo struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	DateTime string `json:"dateTime"`
}

func NewClientInfo(host, name, dateTime string) *ClientInfo {
	return &ClientInfo{host, name, dateTime}
}
