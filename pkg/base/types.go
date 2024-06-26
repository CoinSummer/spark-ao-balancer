package base

type ArgusConfig struct {
	Name              string `mapstructure:"name"`
	BotPk             string `mapstructure:"bot_pk"`
	BotAddress        string `mapstructure:"bot_address"`
	Argus             string `mapstructure:"argus"`
	Safe              string `mapstructure:"safe"`
	ArweaveAddressHex string `mapstructure:"arweave_address_hex"`
	Rpc               string `mapstructure:"rpc"`
	Ws                string `mapstructure:"ws"`
}

type SlackConfig struct {
	Token   string `mapstructure:"token"`
	Channel string `mapstructure:"channel"`
}
