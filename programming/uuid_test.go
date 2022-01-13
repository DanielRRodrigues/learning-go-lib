package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pf ProgrammingFunctions = ProgrammingFunctions{}

func TestNewUUIDWithHyphen(t *testing.T) {
	uuidWithHyphen := pf.NewUUID(false)

	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")
}

func TestNewUUIDWithoutHyphen(t *testing.T) {
	uuidWithoutHyphen := pf.NewUUID(true)

	assert.Len(t, uuidWithoutHyphen, 32)
	assert.NotContains(t, uuidWithoutHyphen, "-")
}
