package error

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"utils/runtimecaller"
)

func TestRuntimeCaller(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	bufferRuntimeCallerFilter := DefaultRuntimeCallerFilter[:]
	DefaultRuntimeCallerFilter = []runtimecaller.Filter{}
	defer func() {
		DefaultRuntimeCallerFilter = bufferRuntimeCallerFilter
	}()

	callinfo, ok := RuntimeCaller(0)
	require.True(ok)
	for _, testcase := range []struct {
		expected        string
		longFile        bool
		line            bool
		replacePackages map[string]string
	}{
		{"github.com//tmsong/utils/error/RuntimeCaller_test.go:24", true, true, nil},
		{"github.com//tmsong/utils/error/RuntimeCaller_test.go", true, false, nil},
		{"RuntimeCaller_test.go:24", false, true, nil},
		{"RuntimeCaller_test.go", false, false, nil},
		{"github.com//tmsong/utils/error/RuntimeCaller_test.go:24", true, true, map[string]string{"tmsong/utils/error": "error"}},
	} {
		buffer := &bytes.Buffer{}
		if _, err := WriteCallInfo(buffer, callinfo, testcase.longFile, testcase.line, testcase.replacePackages); assert.NoError(err) {
			require.EqualValues(testcase.expected, buffer.String())
		}
	}
}
