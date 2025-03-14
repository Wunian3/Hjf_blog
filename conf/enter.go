package conf

type Config struct {
	Mysql   Mysql   `yaml:"mysql"`
	Logger  Logger  `yaml:"logger"`
	SiteInf SiteInf `yaml:"site_inf"`
	System  System  `yaml:"system"`
	QQ      QQ      `yaml:"qq"`
	QiNiu   QiNiu   `yaml:"qi_niu"`
	Email   Email   `yaml:"email"`
	Jwt     Jwt     `yaml:"jwt"`
}
