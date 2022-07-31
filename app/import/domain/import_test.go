package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImport(t *testing.T) {
	t.Parallel()
	t.Run("aaa", func(t *testing.T) {
		filename := "2847-1815492-sync.csv"
		i, err := NewImport(filename)
		require.NoError(t, err)

		assert.Equal(t, i.importId.Value(), "1815492")
		assert.Equal(t, i.shopId.Value(), "2847")
		assert.Equal(t, i.shopFile.Value(), filename)

	})
}
