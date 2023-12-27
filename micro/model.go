package micro

type ConfigEntity struct {
	Address   string `json:"address" yaml:"address" mapstructure:"address"`       // current service deploy address outside
	Namespace string `json:"namespace" yaml:"namespace" mapstructure:"namespace"` // current service namespace format service/group/xxxx
	MaxRetry  uint32 `json:"max_retry" yaml:"maxRetry" mapstructure:"maxRetry"`   // service global lease unexpected disconnection retry max frequency
	Deploy    bool   `json:"deploy" yaml:"maxRetry" mapstructure:"maxRetry"`      // current service deploy mode: local area network or not
}

type ValueEntity struct {
	Name      string `json:"name"`
	Endpoints string `json:"endpoints"`
}

type DiscoverEntity struct {
	Gateway   string `json:"gateway"`
	Namespace string `json:"namespace"`
	Service   string `json:"service"`
}
