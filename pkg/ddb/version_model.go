package ddb

// VersionedModel ...
type VersionedModel interface {
	GetVersion() uint
	IncreaseVersion()
}

// Model ...
type VersionModel struct {
	Version uint `dynamo:"version"`
}

// GetVersion ...
func (vm *VersionModel) GetVersion() uint {
	return vm.Version
}

// IncreaseVersion ...
func (vm *VersionModel) IncreaseVersion() {
	vm.Version++
}
