package some_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/user1.test/task-0/internal/some"
)

func TestSome(t *testing.T) {
	require.Equal(t, "some", some.Some())
}
