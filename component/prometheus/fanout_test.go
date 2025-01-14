package prometheus

import (
	"testing"

	"github.com/prometheus/prometheus/storage"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestRollback(t *testing.T) {
	fanout := NewFanout([]storage.Appendable{NewFanout(nil, "1")}, "")
	app := fanout.Appender(context.Background())
	err := app.Rollback()
	require.NoError(t, err)
}

func TestCommit(t *testing.T) {
	fanout := NewFanout([]storage.Appendable{NewFanout(nil, "1")}, "")
	app := fanout.Appender(context.Background())
	err := app.Commit()
	require.NoError(t, err)
}
