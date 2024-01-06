package micro

type ConfigEntity struct {
	Address   string `json:"address" bson:"address" yaml:"address" mapstructure:"address"`         // current service deploy address outside
	Namespace string `json:"namespace" bson:"namespace" yaml:"namespace" mapstructure:"namespace"` // current service namespace format service/group/xxxx
	MaxRetry  uint32 `json:"max_retry" bson:"max_retry" yaml:"max_retry" mapstructure:"max_retry"` // service global lease unexpected disconnection retry max frequency
	TTL       uint32 `json:"ttl" bson:"ttl" yaml:"ttl" mapstructure:"ttl"`                         // current service lease time
}

type ValueEntity struct {
	Name      string `json:"name"`
	Endpoints string `json:"endpoints"`
}
