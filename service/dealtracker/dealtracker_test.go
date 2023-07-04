package dealtracker

import (
	"context"
	"encoding/json"
	"github.com/bcicen/jstream"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Closer interface {
	Close()
}

func setupTestServer(t *testing.T) (string, Closer) {
	return setupTestServerWithBody(t, `{"0":{"Proposal":{"PieceCID":{"/":"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq"},"PieceSize":34359738368,"VerifiedDeal":true,"Client":"t0100","Provider":"t01000","Label":"bagboea4b5abcatlxechwbp7kjpjguna6r6q7ejrhe6mdp3lf34pmswn27pkkiekz","StartEpoch":0,"EndEpoch":999999999,"StoragePricePerEpoch":"0","ProviderCollateral":"0","ClientCollateral":"0"},"State":{"SectorStartEpoch":0,"LastUpdatedEpoch":691200,"SlashEpoch":-1,"VerifiedClaim":0}}}`)
}

func setupTestServerWithBody(t *testing.T, b string) (string, Closer) {
	body := []byte(b)
	encoder, err := zstd.NewWriter(nil)
	assert.NoError(t, err)
	compressed := encoder.EncodeAll(body, make([]byte, 0, len(body)))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(compressed)
	}))
	encoder.Close()
	return server.URL, server
}

func TestDealStateStreamFromHttpRequest_Compressed(t *testing.T) {
	url, server := setupTestServer(t)
	defer server.Close()
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	depth := 1
	stream, closer, err := DealStateStreamFromHTTPRequest(req, depth, true)
	assert.NoError(t, err)
	defer closer.Close()
	var kvs []jstream.KV
	for s := range stream {
		pair, ok := s.Value.(jstream.KV)
		assert.True(t, ok)
		kvs = append(kvs, pair)
	}
	assert.Len(t, kvs, 1)
	assert.Equal(t, "0", kvs[0].Key)
	assert.Equal(t, "bagboea4b5abcatlxechwbp7kjpjguna6r6q7ejrhe6mdp3lf34pmswn27pkkiekz",
		kvs[0].Value.(map[string]interface{})["Proposal"].(map[string]interface{})["Label"].(string))
}

func TestDealStateStreamFromHttpRequest_UnCompressed(t *testing.T) {
	body := []byte(`{"jsonrpc":"2.0","result":{"0":{"Proposal":{"PieceCID":{"/":"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq"},"PieceSize":34359738368,"VerifiedDeal":true,"Client":"t0100","Provider":"t01000","Label":"bagboea4b5abcatlxechwbp7kjpjguna6r6q7ejrhe6mdp3lf34pmswn27pkkiekz","StartEpoch":0,"EndEpoch":1552977,"StoragePricePerEpoch":"0","ProviderCollateral":"0","ClientCollateral":"0"},"State":{"SectorStartEpoch":0,"LastUpdatedEpoch":691200,"SlashEpoch":-1,"VerifiedClaim":0}}}}`)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer server.Close()
	req, err := http.NewRequest("GET", server.URL, nil)
	assert.NoError(t, err)
	depth := 2
	stream, closer, err := DealStateStreamFromHTTPRequest(req, depth, false)
	assert.NoError(t, err)
	defer closer.Close()
	var kvs []jstream.KV
	for s := range stream {
		pair, ok := s.Value.(jstream.KV)
		assert.True(t, ok)
		kvs = append(kvs, pair)
	}
	assert.Len(t, kvs, 1)
	assert.Equal(t, "0", kvs[0].Key)
	assert.Equal(t, "bagboea4b5abcatlxechwbp7kjpjguna6r6q7ejrhe6mdp3lf34pmswn27pkkiekz",
		kvs[0].Value.(map[string]interface{})["Proposal"].(map[string]interface{})["Label"].(string))
}

func TestTrackDeal(t *testing.T) {
	url, server := setupTestServer(t)
	defer server.Close()
	tracker := NewDealTracker(nil, 0, url, "", "")
	var deals []Deal
	callback := func(dealID uint64, deal Deal) error {
		deals = append(deals, deal)
		return nil
	}
	err := tracker.trackDeal(context.Background(), callback)
	assert.NoError(t, err)
	assert.Len(t, deals, 1)
}

func TestShouldTrackDeal(t *testing.T) {
	db := database.OpenInMemory()
	tracker := NewDealTracker(db, time.Second, "", "", "")
	should, err := tracker.shouldTrackDeal(context.Background())
	assert.NoError(t, err)
	assert.True(t, should)
	should, err = tracker.shouldTrackDeal(context.Background())
	assert.NoError(t, err)
	assert.False(t, should)
	time.Sleep(time.Second)
	should, err = tracker.shouldTrackDeal(context.Background())
	assert.NoError(t, err)
	assert.True(t, should)
	should, err = tracker.shouldTrackDeal(context.Background())
	assert.NoError(t, err)
	assert.False(t, should)
}

func TestRunOnce(t *testing.T) {
	db := database.OpenInMemory()
	err := db.Create(&model.Wallet{
		ID:      "t0100",
		Address: "t3xxx",
	}).Error
	assert.NoError(t, err)
	d1 := uint64(1)
	d2 := uint64(2)
	d4 := uint64(4)
	d6 := uint64(6)
	err = db.Create([]model.Deal{
		{
			DealID:           &d1,
			State:            model.DealActive,
			ClientID:         "t0100",
			Provider:         "sp1",
			ProposalID:       "proposal1",
			Label:            "label1",
			PieceCID:         "cid1",
			PieceSize:        100,
			StartEpoch:       100,
			EndEpoch:         999999999,
			SectorStartEpoch: 0,
			Verified:         true,
		},
		{
			DealID:           &d2,
			State:            model.DealPublished,
			ClientID:         "t0100",
			Provider:         "sp1",
			ProposalID:       "proposal2",
			Label:            "label2",
			PieceCID:         "cid2",
			PieceSize:        100,
			StartEpoch:       100,
			EndEpoch:         999999999,
			SectorStartEpoch: 0,
			Verified:         true,
		},
		{
			State:            model.DealProposed,
			ClientID:         "t0100",
			Provider:         "sp1",
			ProposalID:       "proposal3",
			Label:            "label3",
			PieceCID:         "cid3",
			PieceSize:        100,
			StartEpoch:       999999998,
			EndEpoch:         999999999,
			SectorStartEpoch: 0,
			Verified:         true,
		},
		{
			DealID:           &d4,
			State:            model.DealActive,
			ClientID:         "t0100",
			Provider:         "sp1",
			ProposalID:       "proposal4",
			Label:            "label4",
			PieceCID:         "cid4",
			PieceSize:        100,
			StartEpoch:       100,
			EndEpoch:         200,
			SectorStartEpoch: 100,
			Verified:         true,
		},
		{
			State:            model.DealProposed,
			ClientID:         "t0100",
			Provider:         "sp1",
			ProposalID:       "proposal5",
			Label:            "label5",
			PieceCID:         "cid5",
			PieceSize:        100,
			StartEpoch:       100,
			EndEpoch:         200,
			SectorStartEpoch: 0,
			Verified:         true,
		},
		{
			DealID:           &d6,
			State:            model.DealPublished,
			ClientID:         "t0100",
			Provider:         "sp1",
			ProposalID:       "proposal6",
			Label:            "label6",
			PieceCID:         "cid6",
			PieceSize:        100,
			StartEpoch:       100,
			EndEpoch:         200,
			SectorStartEpoch: 0,
			Verified:         true,
		},
	}).Error
	assert.NoError(t, err)

	// Deal 1 : Active -> Slashed
	// Deal 2 : Published -> Active
	// Deal 3 : Proposed -> Published
	// Deal 4 : Active -> Expired
	// Deal 5 : Proposed -> Expired
	// Deal 6 : Published -> Expired
	deals := map[string]Deal{
		"1": {
			Proposal: DealProposal{
				PieceCID:             Cid{Root: "cid1"},
				PieceSize:            100,
				VerifiedDeal:         true,
				Client:               "t0100",
				Provider:             "sp1",
				StartEpoch:           100,
				EndEpoch:             999999999,
				StoragePricePerEpoch: "0",
				Label:                "label1",
			},
			State: DealState{
				SectorStartEpoch: 0,
				LastUpdatedEpoch: 0,
				SlashEpoch:       100,
			},
		},
		"2": {
			Proposal: DealProposal{
				PieceCID:             Cid{Root: "cid2"},
				PieceSize:            100,
				VerifiedDeal:         true,
				Client:               "t0100",
				Provider:             "sp1",
				StartEpoch:           100,
				EndEpoch:             999999999,
				StoragePricePerEpoch: "0",
				Label:                "label2",
			},
			State: DealState{
				SectorStartEpoch: 200,
				LastUpdatedEpoch: -1,
				SlashEpoch:       -1,
			},
		},
		"3": {
			Proposal: DealProposal{
				PieceCID:             Cid{Root: "cid3"},
				PieceSize:            100,
				VerifiedDeal:         true,
				Client:               "t0100",
				Provider:             "sp1",
				StartEpoch:           999999998,
				EndEpoch:             999999999,
				StoragePricePerEpoch: "0",
				Label:                "label3",
			},
			State: DealState{
				SectorStartEpoch: -1,
				LastUpdatedEpoch: -1,
				SlashEpoch:       -1,
			},
		},
		"7": {
			Proposal: DealProposal{
				PieceCID:             Cid{Root: "cid7"},
				PieceSize:            100,
				VerifiedDeal:         true,
				Client:               "t0100",
				Provider:             "sp1",
				StartEpoch:           100,
				EndEpoch:             999999999,
				StoragePricePerEpoch: "0",
				Label:                "label7",
			},
			State: DealState{
				SectorStartEpoch: 100,
				LastUpdatedEpoch: -1,
				SlashEpoch:       -1,
			},
		},
	}
	body, err := json.Marshal(deals)
	url, server := setupTestServerWithBody(t, string(body))
	defer server.Close()
	assert.NoError(t, err)
	tracker := NewDealTracker(db, time.Minute, url, "", "")
	err = tracker.runOnce(context.Background())
	assert.NoError(t, err)
	var allDeals []model.Deal
	err = db.Find(&allDeals).Error
	assert.NoError(t, err)
	assert.Len(t, allDeals, 7)
	assert.Equal(t, model.DealSlashed, allDeals[0].State)
	assert.Equal(t, model.DealActive, allDeals[1].State)
	assert.Equal(t, model.DealPublished, allDeals[2].State)
	assert.Equal(t, model.DealExpired, allDeals[3].State)
	assert.Equal(t, model.DealExpired, allDeals[4].State)
	assert.Equal(t, model.DealExpired, allDeals[5].State)
	assert.Equal(t, model.DealActive, allDeals[6].State)
}
