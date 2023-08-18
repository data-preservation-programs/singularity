package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/cockroachdb/errors"
	"github.com/ipfs/go-cid"
)

var ErrInvalidCIDEntry = errors.New("invalid CID entry in the database")
var ErrInvalidStringSliceEntry = errors.New("invalid string slice entry in the database")
var ErrInvalidStringMapEntry = errors.New("invalid string map entry in the database")

type StringSlice []string

type StringMap map[string]string

type CID cid.Cid

func (c CID) MarshalBinary() ([]byte, error) {
	return cid.Cid(c).MarshalBinary()
}

func (c *CID) UnmarshalBinary(b []byte) error {
	var c2 cid.Cid
	err := c2.UnmarshalBinary(b)
	if err != nil {
		return errors.WithStack(err)
	}
	*c = CID(c2)
	return nil
}

func (c CID) MarshalJSON() ([]byte, error) {
	if cid.Cid(c) == cid.Undef {
		return json.Marshal("")
	}

	return json.Marshal(cid.Cid(c).String())
}

func (c CID) String() string {
	if cid.Cid(c) == cid.Undef {
		return ""
	}
	return cid.Cid(c).String()
}

func (c *CID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return errors.WithStack(err)
	}

	if s == "" {
		*c = CID(cid.Undef)
	} else {
		cid, err := cid.Decode(s)
		if err != nil {
			return errors.WithStack(err)
		}
		*c = CID(cid)
	}

	return nil
}

func (c CID) Value() (driver.Value, error) {
	if cid.Cid(c) == cid.Undef {
		return []byte(nil), nil
	}
	return cid.Cid(c).Bytes(), nil
}

func (c *CID) Scan(src any) error {
	if src == nil {
		*c = CID(cid.Undef)
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return ErrInvalidCIDEntry
	}

	if len(source) == 0 {
		*c = CID(cid.Undef)
		return nil
	}

	cid, err := cid.Cast(source)
	if err != nil {
		return errors.Wrap(err, "failed to cast CID")
	}

	*c = CID(cid)
	return nil
}

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}
func (m StringMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (ss *StringSlice) Scan(src any) error {
	if src == nil {
		*ss = nil
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return ErrInvalidStringSliceEntry
	}

	return json.Unmarshal(source, ss)
}

func (m *StringMap) Scan(src any) error {
	if src == nil {
		*m = nil
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return ErrInvalidStringMapEntry
	}

	return json.Unmarshal(source, m)
}

type WorkerType string

type JobState string

type JobType string

const (
	DealTracker   WorkerType = "deal_tracker"
	DealPusher    WorkerType = "deal_pusher"
	DatasetWorker WorkerType = "dataset_worker"
)

const (
	Scan   JobType = "scan"
	Pack   JobType = "pack"
	DagGen JobType = "daggen"
)

var JobTypes = []JobType{
	Scan,
	Pack,
	DagGen,
}

var JobTypeStrings = []string{
	string(Scan),
	string(Pack),
	string(DagGen),
}

var JobStates = []JobState{
	Created,
	Ready,
	Paused,
	Processing,
	Complete,
	Error,
}

var JobStateStrings = []string{
	string(Created),
	string(Ready),
	string(Paused),
	string(Processing),
	string(Complete),
	string(Error),
}

const (
	// Created means the job has been created is not ready for processing.
	Created JobState = "created"
	// Ready means the job is ready for processing.
	Ready JobState = "ready"
	// Paused means the job is ready but has been paused and should not be picked up for processing.
	Paused JobState = "paused"
	// Processing means the job is currently being processed.
	Processing JobState = "processing"
	// Complete means the job is complete.
	Complete JobState = "complete"
	// Error means the job has some error.
	Error JobState = "error"
)

var ErrInvalidJobState = errors.New("invalid job state")

func (js *JobState) Set(value string) error {
	for _, state := range JobStates {
		if state == JobState(value) {
			*js = state
			return nil
		}
	}
	return ErrInvalidJobState
}

func (js *JobState) String() string {
	return string(*js)
}
