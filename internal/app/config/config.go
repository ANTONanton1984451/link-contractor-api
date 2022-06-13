package config

type (
	GlobalConfig struct {
		Dev *SubConfig `yaml:"dev"`
	}

	SubConfig struct {
		DbDsn                  string `yaml:"db_dsn"`
		RedirectHost           string `yaml:"redirect_host"`
		RetriesLinkCreateCount int64  `yaml:"retries_link_create_count"`
		WorkPort               string `yaml:"work_port"`
		MaxDbConn              int64  `yaml:"max_conn"`
	}
)
