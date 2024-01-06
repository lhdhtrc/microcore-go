package micro

type ConfigEntity struct {
	Address   string `json:"address" bson:"address" yaml:"address" mapstructure:"address"`         // current service deploy address outside
	Namespace string `json:"namespace" bson:"namespace" yaml:"namespace" mapstructure:"namespace"` // current service namespace format service/group/xxxx
	MaxRetry  uint32 `json:"max_retry" bson:"max_retry" yaml:"max_retry" mapstructure:"max_retry"` // service global lease unexpected disconnection retry max frequency
	Deploy    bool   `json:"deploy" bson:"deploy" yaml:"deploy" mapstructure:"deploy"`             // current service deploy mode: local area network or not
	TTL       uint32 `json:"ttl" bson:"ttl" yaml:"ttl" mapstructure:"ttl"`                         // current service lease time
}

type ValueEntity struct {
	Name      string `json:"name"`
	Endpoints string `json:"endpoints"`
}
