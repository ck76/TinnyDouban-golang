package setting

import "github.com/spf13/viper"

type Settring struct {
	vp *viper.Viper
}

func NewSetting() (*Settring, error) {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Settring{vp}, nil

}
