package hust_wireless

type User struct {
	name     string
	password string
	index    string
}

type Requests struct {
	UserId      string `json:"userId"`
	Password    string `json:"password"`
	QueryString string `json:"queryString"`
	Service     string `json:"service"`
	OperatorPwd string `json:"operatorPwd"`
	Validcode   string `json:"validcode"`
}

type Responses struct {
	UserIndex         string `json:"userIndex"`
	Result            string `json:"result"`
	Message           string `json:"message"`
	KeepaliveInterval int    `json:"keepaliveInterval"`
	ValidCodeUrl      string `json:"validCodeUrl"`
}
