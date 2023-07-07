//go:build exclude

package store

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestItemReferenceBlockStore(t *testing.T) {
	// Create a temporary test file
	f, err := ioutil.TempFile("", "test-file")
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
	store := ItemReferenceBlockStore{DB: db, HandlerResolver: datasource.NewDefaultHandlerResolver()}

	source := model.Source{
		Type: model.Local,
	}
	require.NoError(t, db.Create(&source).Error)

	item := model.Item{
		SourceID: source.ID,
		Type:     model.File,
		Path:     f.Name(),
		Size:     11,
		Offset:   0,
		Length:   11,
	}
	require.NoError(t, db.Create(&item).Error)

	c := cid.Decode("bafy2bzaceajbdbdel5jjeehkborcsrub5cyrq4om3ee5umomqd323ynkpyjh4")
	require.NoError(t, err)
	// Test Has method with a non-existent CID
	has, err := store.Has(context.Background(), c)
	require.NoError(t, err)
	require.False(t, has)

	// Test Get method with a non-existent CID
	_, err = store.Get(context.Background(), c)
	require.Error(t, err)
	require.ErrorContains(t, err, "could not find")

	// Test GetSize method with a non-existent CID
	_, err = store.GetSize(context.Background(), c)
	require.Error(t, err)
	require.ErrorContains(t, err, "could not find")

	// Create a new CarBlock record in the database referencing the CID of the test data
	cb := model.ItemBlock{
		ItemID: item.ID,
		CID:    c.String(),
		Offset: 5,
		Length: 5,
	}
	err = db.Create(&cb).Error
	require.NoError(t, err)

	// Test GetSize method with an existing CID
	size, err := store.GetSize(context.Background(), c)
	require.NoError(t, err)
	require.Equal(t, 5, size)

	// Test Get method with an existing CID
	block, err := store.Get(context.Background(), c)
	require.NoError(t, err)
	require.Equal(t, " worl", string(block.RawData()))
}
