package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRawBlockStore(t *testing.T) {
	// Initialize a new in-memory SQLite database for testing
	db := database.OpenInMemory()
	defer database.DropAll(db)

	// Create a new instance of the RawBlockStore
	store := RawBlockStore{DB: db}

	c := cid.MustParse("bafy2bzaceajbdbdel5jjeehkborcsrub5cyrq4om3ee5umomqd323ynkpyjh4")
	// Test Has method with a non-existent CID
	has, err := store.Has(context.Background(), c)
	assert.NoError(t, err)
	assert.False(t, has)

	// Test Get method with a non-existent CID
	_, err = store.Get(context.Background(), c)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "could not find")

	// Test GetSize method with a non-existent CID
	_, err = store.GetSize(context.Background(), c)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "could not find")

	// Test Put method
	err = db.Create(&model.RawBlock{
		CID:     c.String(),
		Block:   []byte("hello world"),
		Length:  11,
	}).Error
	assert.NoError(t, err)

	// Test Has method with an existing CID
	has, err = store.Has(context.Background(), c)
	assert.NoError(t, err)
	assert.True(t, has)

	// Test Get method with an existing CID
	getBlock, err := store.Get(context.Background(), c)
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), getBlock.RawData())

	// Test GetSize method with an existing CID
	size, err := store.GetSize(context.Background(), c)
	assert.NoError(t, err)
	assert.Equal(t, 11, size)

}
