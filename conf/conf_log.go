package conf

type Logger struct {
	Level    string `yaml:" level"`
	Director string `yaml:"director"`
	//行号和目的
	ShowLine     bool   `yaml:"show-line"`
	LogInConsole bool   `yaml:"log-in-console"`
	Prefix       string `yaml:"prefix"`
}
