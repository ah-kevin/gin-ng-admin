package config

// LogHook 日志钩子
type LogHook string

type Log struct {
	Level         int      `json:"level,omitempty" mapstructure:"level"`
	Format        string   `json:"format,omitempty" mapstructure:"format"`
	Output        string   `json:"output,omitempty" mapstructure:"output"`
	OutputFile    string   `json:"output-file,omitempty" mapstructure:"output-file"`
	EnableHook    bool     `json:"enable-hook,omitempty" mapstructure:"enable-hook"`
	HookLevels    []string `json:"hook-levels,omitempty" mapstructure:"hook-levels"`
	Hook          LogHook  `json:"hook,omitempty" mapstructure:"hook"`
	HookMaxThread int      `json:"hook-max-thread,omitempty" mapstructure:"hook-max-thread"`
	HookMaxBuffer int      `json:"hook-max-buffer,omitempty" mapstructure:"hook-max-buffer"`
}
