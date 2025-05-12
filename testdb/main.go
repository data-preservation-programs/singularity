package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"gorm.io/gorm"
)

//nolint:gosec
var r = rand.New(rand.NewSource(0))

func main() {
	gofakeit.Seed(1)
	err := run()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.TODO()
	db, closer, err := database.OpenWithLogger("sqlite:test.db")
	if err != nil {
		return errors.WithStack(err)
	}
	defer closer.Close()
	err = model.DropAll(db)
	if err != nil {
		return errors.WithStack(err)
	}

	err = model.Migrator(db).Migrate()
	if err != nil {
		return errors.WithStack(err)
	}

	err = createPreparation(ctx, db)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func createPreparation(ctx context.Context, db *gorm.DB) error {
	db = db.WithContext(ctx)
	// Setup source storage
	source := model.Storage{
		Name:   gofakeit.Noun(),
		Type:   "s3",
		Path:   toLowerSnakeCase(gofakeit.BookTitle()),
		Config: map[string]string{"provider": "AWS"},
	}
	// Setup output storage
	output := model.Storage{
		Name: gofakeit.Noun(),
		Type: "local",
		Path: urlToPath(gofakeit.URL()),
	}
	// Setup wallet
	wallet := model.Wallet{
		ActorID: fmt.Sprintf("f0%d", r.Intn(10000)),
		Address: "f1" + randomLetterString(39),
	}

	// Setup preparation
	preparation := model.Preparation{
		Name:           gofakeit.AppName(),
		MaxSize:        30 << 30,
		PieceSize:      1 << 35,
		Wallets:        []model.Wallet{wallet},
		SourceStorages: []model.Storage{source},
	}

	err := db.Create(&preparation).Error
	if err != nil {
		return errors.WithStack(err)
	}
	err = db.Model(&preparation).Association("OutputStorages").Append(&output)
	if err != nil {
		return errors.WithStack(err)
	}

	var sourceAttachment model.SourceAttachment
	err = db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Root Directory
	root := model.Directory{
		CID:          randomCID(),
		Name:         "",
		AttachmentID: sourceAttachment.ID,
	}
	err = db.Create(&root).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Setup a folder with lots of files
	lots := model.Directory{
		CID:          randomCID(),
		Name:         "lots_of_files",
		AttachmentID: sourceAttachment.ID,
		ParentID:     ptr.Of(root.ID),
	}
	err = db.Create(&lots).Error
	if err != nil {
		return errors.WithStack(err)
	}

	job := model.Job{
		Type:         model.Pack,
		State:        model.Complete,
		AttachmentID: sourceAttachment.ID,
	}
	err = db.Create(&job).Error
	if err != nil {
		return errors.WithStack(err)
	}

	var files []model.File
	for i := range r.Intn(10_000) {
		size := r.Int63n(1 << 20)
		rCID := randomCID()
		files = append(files, model.File{
			CID:              rCID,
			Path:             fmt.Sprintf("lots_of_files/%d-%s.txt", i, gofakeit.Noun()),
			Hash:             randomLetterString(6),
			Size:             size,
			LastModifiedNano: randomLastModifiedNano(),
			AttachmentID:     sourceAttachment.ID,
			DirectoryID:      ptr.Of(lots.ID),
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: size,
				CID:    rCID,
				JobID:  ptr.Of(job.ID),
			}},
		})
	}
	err = db.Create(&files).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Setup a folder with a single large file
	large := model.Directory{
		CID:          randomCID(),
		Name:         "large_files",
		AttachmentID: sourceAttachment.ID,
		ParentID:     ptr.Of(root.ID),
	}
	err = db.Create(&large).Error
	if err != nil {
		return errors.WithStack(err)
	}

	largeFile := model.File{
		CID:              randomCID(),
		Path:             "large_files/large.txt",
		Hash:             randomLetterString(6),
		Size:             100 << 34,
		LastModifiedNano: randomLastModifiedNano(),
		AttachmentID:     sourceAttachment.ID,
		DirectoryID:      ptr.Of(large.ID),
		FileRanges:       nil,
	}

	for i := range 100 {
		largeFile.FileRanges = append(largeFile.FileRanges, model.FileRange{
			Offset: int64(i << 34),
			Length: 1 << 34,
			CID:    randomCID(),
			Job: &model.Job{
				Type:         model.Pack,
				State:        model.Complete,
				AttachmentID: sourceAttachment.ID,
			},
		})
	}
	err = db.Create(&largeFile).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Setup a file with multiple versions
	for range 10 {
		size := r.Int63n(1 << 20)
		rCID := randomCID()
		err = db.Create(&model.File{
			CID:              rCID,
			Path:             "multiple_versions.txt",
			Hash:             randomLetterString(6),
			Size:             r.Int63n(1 << 20),
			LastModifiedNano: randomLastModifiedNano(),
			AttachmentID:     sourceAttachment.ID,
			DirectoryID:      ptr.Of(root.ID),
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: size,
				CID:    rCID,
				Job: &model.Job{
					Type:         model.Pack,
					State:        model.Complete,
					AttachmentID: sourceAttachment.ID,
				},
			}},
		}).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Setup CAR for each job
	var jobs []model.Job
	err = db.Where("type = ?", model.Pack).Find(&jobs).Error
	if err != nil {
		return errors.WithStack(err)
	}
	for _, job := range jobs {
		pieceCID, err := randomPieceCID()
		if err != nil {
			return errors.WithStack(err)
		}
		err = db.Create(&model.Car{
			PieceCID:      pieceCID,
			PieceSize:     1 << 35,
			RootCID:       randomCID(),
			FileSize:      1 << 34,
			StorageID:     ptr.Of(output.ID),
			StoragePath:   pieceCID.String() + ".car",
			PreparationID: preparation.ID,
			AttachmentID:  ptr.Of(sourceAttachment.ID),
			JobID:         ptr.Of(job.ID),
		}).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Some Car files without association with the preparation
	for range 5 {
		pieceCID, err := randomPieceCID()
		if err != nil {
			return errors.WithStack(err)
		}
		err = db.Create(&model.Car{
			PieceCID:      pieceCID,
			PieceSize:     1 << 35,
			RootCID:       randomCID(),
			FileSize:      1 << 34,
			StorageID:     ptr.Of(output.ID),
			StoragePath:   pieceCID.String() + ".car",
			PreparationID: preparation.ID,
		}).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Create a one-off schedule
	oneoff := model.Schedule{
		Provider:       fmt.Sprintf("f0%d", r.Intn(10000)),
		Verified:       true,
		KeepUnsealed:   true,
		AnnounceToIPNI: true,
		StartDelay:     24 * time.Hour,
		Duration:       365 * 24 * time.Hour,
		State:          model.ScheduleCompleted,
		Notes:          "this is a one-off schedule",
		PreparationID:  preparation.ID,
	}
	err = db.Create(&oneoff).Error
	if err != nil {
		return errors.WithStack(err)
	}

	cron := model.Schedule{
		Provider:           fmt.Sprintf("f0%d", r.Intn(10000)),
		Verified:           true,
		KeepUnsealed:       true,
		AnnounceToIPNI:     true,
		StartDelay:         24 * time.Hour,
		Duration:           365 * 24 * time.Hour,
		ScheduleCron:       "@hourly",
		ScheduleDealNumber: 100,
		State:              model.ScheduleCompleted,
		Notes:              "this is a cron schedule",
		PreparationID:      preparation.ID,
	}
	err = db.Create(&cron).Error
	if err != nil {
		return errors.WithStack(err)
	}

	var cars []model.Car
	err = db.Where("preparation_id = ?", preparation.ID).Find(&cars).Error
	if err != nil {
		return errors.WithStack(err)
	}
	// Create some deal with oneoff and cron
	for _, schedule := range []model.Schedule{oneoff, cron} {
		for _, car := range cars {
			states := []model.DealState{
				model.DealProposed,
				model.DealPublished,
				model.DealSlashed,
				model.DealActive}
			state := states[r.Intn(len(states))]
			deal := model.Deal{
				State:      state,
				Provider:   schedule.Provider,
				ProposalID: uuid.NewString(),
				Label:      car.RootCID.String(),
				PieceCID:   car.PieceCID,
				PieceSize:  car.PieceSize,
				DealID:     nil,
				StartEpoch: int32(10000 + r.Intn(10000)),
				EndEpoch:   int32(20000 + r.Intn(10000)),
				Price:      "0",
				Verified:   true,
				ScheduleID: ptr.Of(schedule.ID),
				ClientID:   wallet.ActorID,
			}
			if state == model.DealActive {
				deal.SectorStartEpoch = int32(10000 + r.Intn(10000))
			}
			if state == model.DealProposed || state == model.DealPublished {
				deal.DealID = ptr.Of(uint64(r.Intn(10000)))
			}
			err = db.Create(&deal).Error
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func toLowerSnakeCase(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "_")
}

func urlToPath(s string) string {
	u, _ := url.Parse(s)
	return u.Path
}

func randomLetterString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz123456789"

	b := make([]byte, length)
	for i := range b {
		//nolint:gosec
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func randomCID() model.CID {
	value := cid.NewCidV1(cid.Raw, util.Hash([]byte(randomLetterString(6))))
	return model.CID(value)
}

func randomPieceCID() (model.CID, error) {
	calc := &commp.Calc{}
	_, err := bytes.NewBufferString(randomLetterString(1000)).WriteTo(calc)
	if err != nil {
		return model.CID{}, errors.WithStack(err)
	}
	c, _, err := pack.GetCommp(calc, 1<<30)
	if err != nil {
		return model.CID{}, errors.WithStack(err)
	}
	return model.CID(c), nil
}

func randomLastModifiedNano() int64 {
	return time.Unix(r.Int63n(1<<20), 0).UnixNano()
}
