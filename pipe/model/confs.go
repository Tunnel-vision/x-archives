package model


var Conf *Configuration

var Models = []interface{}{
	&User{},
}


type Configuration struct {
	Server string // server name host and port
	StaticServer string // static resource
	StaticResourceVersion string
	LogLevel string
	ShowSQL bool
	SessionSecret string
	SessionMaxAge int
	SQLite string
	MYSQL string
	Port string
	AxiosBaseURL string
	MockServer string
}

func LoadConf()  {
	Conf =  &Configuration{}
	Conf.SQLite = "pipe.db"
}