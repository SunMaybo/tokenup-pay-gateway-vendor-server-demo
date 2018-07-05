package config

type SystemConfig struct {
	Url                    string `yaml:"url"`
	AppId                  string `yaml:"app_id"`
	AppKey                 string `yaml:"app_key"`
	NotifyUrl              string `yaml:"notify_url"`
	PrivateKey             string `yaml:"private_key"`
	CallBackPartyPublicKey string `yaml:"call_back_party_public_key"`
}
