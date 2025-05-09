package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/cockroachdb/errors"
	"github.com/ipfs/go-cid"
)

var (
	ErrInvalidCIDEntry         = errors.New("invalid CID entry in the database")
	ErrInvalidStringSliceEntry = errors.New("invalid string slice entry in the database")
	ErrInvalidStringMapEntry   = errors.New("invalid string map entry in the database")
	ErrInvalidHTTPConfigEntry  = errors.New("invalid ClientConfig entry in the database")
)

type StringSlice []string

type ConfigMap map[string]string

type CID cid.Cid

type ClientConfig struct {
	ConnectTimeout          *time.Duration    `cbor:"1,keyasint,omitempty"  json:"connectTimeout,omitempty"          swaggertype:"primitive,integer"` // HTTP Client Connect timeout
	Timeout                 *time.Duration    `cbor:"2,keyasint,omitempty"  json:"timeout,omitempty"                 swaggertype:"primitive,integer"` // IO idle timeout
	ExpectContinueTimeout   *time.Duration    `cbor:"3,keyasint,omitempty"  json:"expectContinueTimeout,omitempty"   swaggertype:"primitive,integer"` // Timeout when using expect / 100-continue in HTTP
	InsecureSkipVerify      *bool             `cbor:"4,keyasint,omitempty"  json:"insecureSkipVerify,omitempty"`                                      // Do not verify the server SSL certificate (insecure)
	NoGzip                  *bool             `cbor:"5,keyasint,omitempty"  json:"noGzip,omitempty"`                                                  // Don't set Accept-Encoding: gzip
	UserAgent               *string           `cbor:"6,keyasint,omitempty"  json:"userAgent,omitempty"`                                               // Set the user-agent to a specified string
	CaCert                  []string          `cbor:"7,keyasint,omitempty"  json:"caCert,omitempty"`                                                  // Paths to CA certificate used to verify servers
	ClientCert              *string           `cbor:"8,keyasint,omitempty"  json:"clientCert,omitempty"`                                              // Path to Client SSL certificate (PEM) for mutual TLS auth
	ClientKey               *string           `cbor:"9,keyasint,omitempty"  json:"clientKey,omitempty"`                                               // Path to Client SSL private key (PEM) for mutual TLS auth
	Headers                 map[string]string `cbor:"10,keyasint,omitempty" json:"headers,omitempty"`                                                 // Set HTTP header for all transactions
	DisableHTTP2            *bool             `cbor:"11,keyasint,omitempty" json:"disableHttp2,omitempty"`                                            // Disable HTTP/2 in the transport
	DisableHTTPKeepAlives   *bool             `cbor:"12,keyasint,omitempty" json:"disableHttpKeepAlives,omitempty"`                                   // Disable HTTP keep-alives and use each connection once.
	RetryMaxCount           *int              `cbor:"13,keyasint,omitempty" json:"retryMaxCount,omitempty"`                                           // Maximum number of retries. Default is 10 retries.
	RetryDelay              *time.Duration    `cbor:"14,keyasint,omitempty" json:"retryDelay,omitempty"              swaggertype:"primitive,integer"` // Delay between retries. Default is 1s.
	RetryBackoff            *time.Duration    `cbor:"15,keyasint,omitempty" json:"retryBackoff,omitempty"            swaggertype:"primitive,integer"` // Constant backoff between retries. Default is 1s.
	RetryBackoffExponential *float64          `cbor:"16,keyasint,omitempty" json:"retryBackoffExponential,omitempty"`                                 // Exponential backoff between retries. Default is 1.0.
	SkipInaccessibleFile    *bool             `cbor:"17,keyasint,omitempty" json:"skipInaccessibleFile,omitempty"`                                    // Skip inaccessible files. Default is false.
	UseServerModTime        *bool             `cbor:"18,keyasint,omitempty" json:"useServerModTime,omitempty"`                                        // Use server modified time instead of object metadata
	LowLevelRetries         *int              `cbor:"19,keyasint,omitempty" json:"lowlevelRetries,omitempty"`                                         // Maximum number of retries for low-level client errors. Default is 10 retries.
	ScanConcurrency         *int              `cbor:"20,keyasint,omitempty" json:"scanConcurrency,omitempty"`                                         // Maximum number of concurrent scan requests. Default is 1.
}

func (c CID) MarshalBinary() ([]byte, error) {
	return cid.Cid(c).MarshalBinary()
}

func (c *CID) UnmarshalBinary(b []byte) error {
	if len(b) == 0 {
		*c = CID(cid.Undef)
		return nil
	}
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

func (m ConfigMap) Value() (driver.Value, error) {
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

func (m *ConfigMap) Scan(src any) error {
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

func IsSecretConfigName(key string) bool {
	k := strings.ToLower(key)
	return strings.Contains(k, "secret") || strings.Contains(k, "pass") || strings.Contains(k, "token") || strings.Contains(k, "key")
}

func (m ConfigMap) String() string {
	if m == nil {
		return "<nil>"
	}
	var values []string
	for k, v := range m {
		if v == "" || v == "0" || v == "false" {
			continue
		}
		if IsSecretConfigName(k) {
			v = "*"
		}
		values = append(values, k+":"+v)
	}
	slices.Sort(values)
	return strings.Join(values, " ")
}

func (c ClientConfig) Value() (driver.Value, error) { //nolint:recvcheck
	return json.Marshal(c)
}

func (c ClientConfig) String() string {
	var values []string
	if c.ConnectTimeout != nil {
		values = append(values, "connectTimeout:"+c.ConnectTimeout.String())
	}
	if c.Timeout != nil {
		values = append(values, "timeout:"+c.Timeout.String())
	}
	if c.ExpectContinueTimeout != nil {
		values = append(values, "expectContinueTimeout:"+c.ExpectContinueTimeout.String())
	}
	if c.InsecureSkipVerify != nil {
		values = append(values, "insecureSkipVerify:"+strconv.FormatBool(*c.InsecureSkipVerify))
	}
	if c.NoGzip != nil {
		values = append(values, "noGzip:"+strconv.FormatBool(*c.NoGzip))
	}
	if c.UserAgent != nil {
		values = append(values, "userAgent:"+*c.UserAgent)
	}
	if len(c.CaCert) > 0 {
		values = append(values, "caCert:"+strings.Join(c.CaCert, ","))
	}
	if c.ClientCert != nil {
		values = append(values, "clientCert:"+*c.ClientCert)
	}
	if c.ClientKey != nil {
		values = append(values, "clientKey:"+*c.ClientKey)
	}
	if len(c.Headers) > 0 {
		values = append(values, "headers:<hidden>")
	}
	if c.DisableHTTP2 != nil {
		values = append(values, "disableHTTP2"+strconv.FormatBool(*c.DisableHTTP2))
	}
	if c.DisableHTTPKeepAlives != nil {
		values = append(values, "disableHTTPKeepAlives:"+strconv.FormatBool(*c.DisableHTTPKeepAlives))
	}
	if c.RetryMaxCount != nil {
		values = append(values, "retryMaxCount:"+strconv.Itoa(*c.RetryMaxCount))
	}
	if c.RetryDelay != nil {
		values = append(values, "retryDelay:"+c.RetryDelay.String())
	}
	if c.RetryBackoff != nil {
		values = append(values, "retryBackoff:"+c.RetryBackoff.String())
	}
	if c.RetryBackoffExponential != nil {
		values = append(values, "retryBackoffExponential:"+fmt.Sprint(*c.RetryBackoffExponential))
	}
	if c.SkipInaccessibleFile != nil {
		values = append(values, "skipInaccessibleFile:"+strconv.FormatBool(*c.SkipInaccessibleFile))
	}
	if c.UseServerModTime != nil {
		values = append(values, "useServerModTime:"+strconv.FormatBool(*c.UseServerModTime))
	}
	if c.LowLevelRetries != nil {
		values = append(values, "lowLevelRetries:"+strconv.Itoa(*c.LowLevelRetries))
	}
	if c.ScanConcurrency != nil {
		values = append(values, "scanConcurrency:"+strconv.Itoa(*c.ScanConcurrency))
	}
	return strings.Join(values, " ")
}

func (c *ClientConfig) Scan(src any) error {
	if src == nil {
		*c = ClientConfig{}
		return nil
	}

	source, ok := src.([]byte)
	if !ok {
		return ErrInvalidHTTPConfigEntry
	}

	return json.Unmarshal(source, c)
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
