package model

const Version = "1.8.9"

var UserAgent = "Pipe/" + Version + "; +https://github.com/b3log/pipe"

var Conf *Configuration

var Models = []interface{}{
	&User{},&Article{},Article{},&Category{},&Comment{},
	&Correlation{},&Navigation{},&Setting{},&Tag{},
}

type Configuration struct {
	Server                string // server name host and port
	StaticServer          string // static resource
	StaticResourceVersion string
	LogLevel              string
	ShowSQL               bool
	SessionSecret         string
	SessionMaxAge         int
	SQLite                string
	MYSQL                 string
	Port                  string
	AxiosBaseURL          string
	MockServer            string
}

func LoadConf() {
	Conf = &Configuration{}
	Conf.SQLite = "pipe.db"
	Conf.Port = "8080"
}
