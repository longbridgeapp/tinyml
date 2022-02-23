package tinyml

import (
	"testing"

	"github.com/longbridgeapp/assert"
)

func Test_parseSecurityTag(t *testing.T) {
	counterID, name, err := parseSecurityTag("ST/US/BABA")
	assert.Error(t, err)

	counterID, name, err = parseSecurityTag("ST/US/BABA#阿里巴巴.US")
	assert.NoError(t, err)
	assert.Equal(t, "ST/US/BABA", counterID)
	assert.Equal(t, "阿里巴巴.US", name)

	counterID, name, err = parseSecurityTag(" st/us/baba # 阿里巴巴.US ")
	assert.NoError(t, err)
	assert.Equal(t, "ST/US/BABA", counterID)
	assert.Equal(t, "阿里巴巴.US", name)
}
