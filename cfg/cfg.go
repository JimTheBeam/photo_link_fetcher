package cfg

import "time"

//  Config - configuration from file config.yaml
type Config struct {
	App CfgApp `yaml:"app"`
}

// CfgApp - app configuration
type CfgApp struct {
	Timeout   time.Duration `yaml:"timeout"`
	UrlScheme string        `yaml:"url_scheme"`
}
