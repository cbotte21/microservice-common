package schema

type Schema[T any] interface {
	Database() string   // MongoDB field
	Collection() string // MongoDB field
	Key() string        // Redis field
}
