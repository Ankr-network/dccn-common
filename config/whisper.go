package config

type Whisper interface {
	GetSecret(saRole, keyPath string) (*Secret, error)
}
