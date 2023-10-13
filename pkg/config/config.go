package config

type (
	app struct {
		Debug bool `default:"true"`
		Key   string
	}
	mail struct {
		SMTP struct {
			Host     string
			Port     int
			Username string
			Password string
		}
		From struct {
			Name    string
			Address string
		}
	}
	postgres struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	}
	redis struct {
		Type     string
		Host     string
		Password string
	}
)

var (
	App = &app{}

	Postgres = &postgres{}

	Redis = &redis{}

	Mail = &mail{}
)
