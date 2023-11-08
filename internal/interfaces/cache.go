package interfaces

type Cache interface {
	GetSchema(schemaName string) (Schema, error)
}
