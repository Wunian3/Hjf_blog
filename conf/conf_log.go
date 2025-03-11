package conf

type Logger struct {
	Level    string `yaml:"level"`
	Director string `yaml:"director"`
	//行号和目的
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
	Prefix       string `yaml:"prefix"`
}
