package configs

import "github.com/spf13/viper"

type Enviroment struct {
	TemplateName string `mapstructure:"TEMPLATE_NAME"`
	FileOutput   string `mapstructure:"FILE_OUTPUT"`
}

func LoadEnviroment(path string) (*Enviroment, error) {
	var cfg *Enviroment
	viper.SetConfigName("certificated")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
