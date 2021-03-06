package env

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type filesystem struct {
	t *testing.T
}

func newFilesystem(t *testing.T) *filesystem {
	return &filesystem{
		t: t,
	}
}

func (f *filesystem) ReadString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		err = fmt.Errorf("can not read test data from file %s: %w", filename, err)
		assert.FailNow(f.t, err.Error())
		return ""
	}

	return string(bytes)
}
