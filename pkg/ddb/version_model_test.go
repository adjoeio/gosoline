package ddb_test

import (
	"github.com/applike/gosoline/pkg/ddb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel_GetVersion(t *testing.T) {
	tests := []struct {
		name    string
		Version uint
		want    uint
	}{
		{"test", 10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ddb.VersionModel{
				Version: tt.Version,
			}
			if got := vm.GetVersion(); got != tt.want {
				t.Errorf("Model.GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModel_IncreaseVersion(t *testing.T) {
	tests := []struct {
		name    string
		Version uint
		want    uint
	}{
		{"test", 10, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &ddb.VersionModel{
				Version: tt.Version,
			}
			vm.IncreaseVersion()

			assert.Equal(t, tt.want, vm.GetVersion())
		})
	}
}
