package unifi //nolint: testpackage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// UniFi Network 11 returns a client's dev_id_override as a numeric string; older
// controllers sent a JSON number. Both must decode, and a missing field must not fail.
func TestUserDevIdOverrideUnmarshal(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input string
		want  int
	}{
		"number":         {input: `{"dev_id_override":1908}`, want: 1908},
		"numeric string": {input: `{"dev_id_override":"1908"}`, want: 1908},
		"empty string":   {input: `{"dev_id_override":""}`, want: 0},
		"absent":         {input: `{"mac":"aa:bb:cc:dd:ee:ff"}`, want: 0},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var u User
			require.NoError(t, json.Unmarshal([]byte(tc.input), &u))
			assert.Equal(t, tc.want, u.DevIdOverride)
		})
	}
}
