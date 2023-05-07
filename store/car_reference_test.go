package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestCarReferenceBlockStore(t *testing.T) {
	// Create a temporary test file
	f, err := ioutil.TempFile("", "test-car-file")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	// Write some test data to the file
	_, err = f.Write([]byte("hello world"))
	require.NoError(t, err)

	// Close the file
	err = f.Close()
	require.NoError(t, err)

	// Initialize a new in-memory SQLite database for testing
	db := database.OpenInMemory()
	defer database.DropAll(db)

	// Create a new instance of the CarReferenceBlockStore
	store := CarReferenceBlockStore{DB: db}

	// Create a new Car record in the database referencing the test file
	car := model.Car{
		FilePath: f.Name(),
	}
	err = db.Create(&car).Error
	require.NoError(t, err)

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

	// Create a new CarBlock record in the database referencing the CID of the test data
	cb := model.CarBlock{
		CarID:  car.ID,
		CID:    c.String(),
		Offset: 5,
		Length: 5,
	}
	err = db.Create(&cb).Error
	require.NoError(t, err)

	// Test GetSize method with an existing CID
	size, err := store.GetSize(context.Background(), c)
	assert.NoError(t, err)
	assert.Equal(t, 5, size)

	// Test Get method with an existing CID
	block, err := store.Get(context.Background(), c)
	assert.NoError(t, err)
	assert.Equal(t, " worl", string(block.RawData()))
}
