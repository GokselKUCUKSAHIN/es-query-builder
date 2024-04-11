package mode_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"

	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
)

func Test_ModeString(t *testing.T) {
	tests := []struct {
		mode   Mode.Mode
		result string
	}{
		{Mode.Min, "min"},
		{Mode.Max, "max"},
		{Mode.Sum, "sum"},
		{Mode.Avg, "avg"},
		{Mode.Median, "median"},
		{Mode.Default, "_default"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.mode.String())
		})
	}
}
