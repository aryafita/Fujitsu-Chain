var ContractConfig = &PluginConfig{
	Name:                  "fujitsu_contract",
	Id:                    1,
	Version:               1,
	SupportedTransactions: []string{"send", "reward"},
	TransactionTypeUrls: []string{
		"type.googleapis.com/types.MessageSend",
		"type.googleapis.com/types.MessageReward",
	},
	EventTypeUrls: nil,
}
