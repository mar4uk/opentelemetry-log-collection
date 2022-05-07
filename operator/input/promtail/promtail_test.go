package promtail

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-log-collection/testutil"
)

func TestBuild(t *testing.T) {
	basicConfig := func() *PromtailInputConfig {
		cfg := NewPromtailInputConfig("testfile")
		return cfg
	}

	cfg := basicConfig()
	op, err := cfg.Build(testutil.Logger(t))
	require.NoError(t, err)
	fmt.Println(op)
}
