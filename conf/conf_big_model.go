package conf

type ModelOption struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Disable bool   `yaml:"disable" json:"disable"`
}

type Setting struct {
	Name      string `yaml:"name" json:"name"`
	Enable    bool   `yaml:"enable" json:"enable"`
	Order     int    `json:"order" yaml:"order"`
	ApiKey    string `yaml:"api-key" json:"api-key"`
	ApiSecret string `yaml:"api-secret" json:"api-secret"`
	Title     string `yaml:"title" json:"title"`
	Prompt    string `yaml:"prompt" json:"prompt"` //加强提示词
	Slogan    string `yaml:"slogan" json:"slogan"`
}

type SessionSetting struct {
	ChatScope    int `yaml:"chat-scope" json:"chat-scope"`       //对话积分消耗
	SessionScope int `yaml:"session-scope" json:"session-scope"` //会话消耗
	DayScope     int `yaml:"day-scope" json:"day-scope"`
}
type BigModel struct {
	Setting        Setting        `yaml:"setting"`
	ModelList      []ModelOption  `yaml:"model-list"`
	SessionSetting SessionSetting `yaml:"session-setting"`
}
