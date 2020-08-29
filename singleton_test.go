package prognoslog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonLog(t *testing.T) {
	// NOTE: this is a cheesy test... just a sanity check

	log1 := SingletonLog()
	log2 := SingletonLog()
	log3 := SingletonLog()
	log4 := SingletonLog()

	assert.Same(t, log1, log2)
	assert.Same(t, log1, log3)
	assert.Same(t, log1, log4)
}
