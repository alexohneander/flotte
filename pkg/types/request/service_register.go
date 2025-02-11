package request

type ServiceRegister struct {
	Name     string `json:"name"`
	NodeType string `json:"nodeType"`
	Address  string `json:"address"`
	Port     string `json:"port"`
}
