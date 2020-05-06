package model

type Server struct {
	JWT     JwtSignStruct `mapstructure:"jwt" json:"jwt"`
	Mysql   Mysql         `mapstructure:"mysql" json:"mysql"`
	Captcha Captcha       `mapstructure:"captcha" json:"captcha"`
}

type JwtSignStruct struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey"`
}

// var JwtSignStruct = "GGBENG"
type Mysql struct {
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Path         string `mapstructure:"path" json:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight"`
}
