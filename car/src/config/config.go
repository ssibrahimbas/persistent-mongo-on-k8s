package config

type App struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
	Db   struct {
		Host string `env:"DB_HOST"`
		Port string `env:"DB_PORT"`
		Name string `env:"DB_NAME"`
		Pw   string `env:"DB_PASSWORD"`
		Usr  string `env:"DB_USERNAME"`
	}
	I18n struct {
		LocaleDir string   `env:"I18N_LOCALE_DIR"`
		Fallback  string   `env:"I18N_FALLBACK"`
		Locales   []string `env:"I18N_LOCALES"`
	}
}
