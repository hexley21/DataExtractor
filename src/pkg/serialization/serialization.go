package serialization

type Processor interface {
	Serializer
	Deserializer
}

type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
}

type Deserializer interface {
	Deserialize(content []byte, data interface{}) error
}
