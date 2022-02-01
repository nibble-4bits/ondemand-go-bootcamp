package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomProperties_ParseInt(t *testing.T) {
	tests := []struct {
		strVal     string
		defaultVal int
		want       intProperty
	}{
		{strVal: "1", defaultVal: 0, want: intProperty(1)},
		{strVal: "-1", defaultVal: 0, want: intProperty(-1)},
		{strVal: "1.2", defaultVal: 0, want: intProperty(0)},
		{strVal: "1000000000000000000000000000000", defaultVal: 0, want: intProperty(0)},
		{strVal: "not parseable", defaultVal: 0, want: intProperty(0)},
		{strVal: "", defaultVal: 0, want: intProperty(0)},
	}

	for _, test := range tests {
		var i intProperty

		i.ParseInt(test.strVal, test.defaultVal)

		assert.Equal(t, i, test.want)
	}
}

func TestCustomProperties_ParseBool(t *testing.T) {
	tests := []struct {
		strVal     string
		defaultVal bool
		want       boolProperty
	}{
		{strVal: "true", defaultVal: true, want: boolProperty(true)},
		{strVal: "false", defaultVal: false, want: boolProperty(false)},
		{strVal: "not parseable", defaultVal: false, want: boolProperty(false)},
		{strVal: "", defaultVal: false, want: boolProperty(false)},
	}

	for _, test := range tests {
		var i boolProperty

		i.ParseBool(test.strVal, test.defaultVal)

		assert.Equal(t, i, test.want)
	}
}
