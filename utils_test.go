package matrixapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripUsername(t *testing.T) {
	assert.Equal(t, "test", stripUsername("@test:host.domain"), "should be equal")
	assert.Equal(t, "test", stripUsername("@test"), "should be equal")
}
