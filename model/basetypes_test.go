package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestClientConfigMarshal(t *testing.T) {
	c := ClientConfig{
		ConnectTimeout:          ptr.Of(time.Second),
		Timeout:                 ptr.Of(time.Second),
		ExpectContinueTimeout:   ptr.Of(time.Second),
		InsecureSkipVerify:      ptr.Of(true),
		NoGzip:                  ptr.Of(true),
		UserAgent:               ptr.Of("x"),
		CaCert:                  []string{"x"},
		ClientCert:              ptr.Of("x"),
		ClientKey:               ptr.Of("x"),
		Headers:                 map[string]string{"x": "x"},
		DisableHTTP2:            ptr.Of(true),
		DisableHTTPKeepAlives:   ptr.Of(true),
		RetryMaxCount:           ptr.Of(10),
		RetryDelay:              ptr.Of(time.Second),
		RetryBackoff:            ptr.Of(time.Second),
		RetryBackoffExponential: ptr.Of(1.0),
		SkipInaccessibleFile:    ptr.Of(true),
		UseServerModTime:        ptr.Of(true),
		LowLevelRetries:         ptr.Of(10),
		ScanConcurrency:         ptr.Of(10),
	}
	data, err := c.Value()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var c2 ClientConfig
	err = c2.Scan(data)
	require.NoError(t, err)
	require.EqualValues(t, c, c2)

	str := c.String()
	require.Equal(t, "connectTimeout:1s timeout:1s expectContinueTimeout:1s insecureSkipVerify:true noGzip:true userAgent:x caCert:x clientCert:x clientKey:x headers:<hidden> disableHTTP2:true disableHTTPKeepAlives:true retryMaxCount:10 retryDelay:1s retryBackoff:1s retryBackoffExponential:1 skipInaccessibleFile:true useServerModTime:true lowLevelRetries:10 scanConcurrency:10", str)
}

func TestConfigMapRedaction(t *testing.T) {
	m := ConfigMap{
		"access_key_id":     "AKIAPUBLIC",
		"secret_access_key": "supersecret",
		"region":            "us-east-1",
	}

	// API-facing JSON marshal redacts secret-named values, keeps the rest
	data, err := json.Marshal(m)
	require.NoError(t, err)
	var out map[string]string
	require.NoError(t, json.Unmarshal(data, &out))
	require.Equal(t, redactedSecret, out["secret_access_key"])
	require.Equal(t, redactedSecret, out["access_key_id"])
	require.Equal(t, "us-east-1", out["region"])

	// persistence keeps the real values
	v, err := m.Value()
	require.NoError(t, err)
	var back ConfigMap
	require.NoError(t, back.Scan(v))
	require.Equal(t, "supersecret", back["secret_access_key"])
	require.Equal(t, "AKIAPUBLIC", back["access_key_id"])
}

func TestClientConfigHeaderRedaction(t *testing.T) {
	c := ClientConfig{Headers: map[string]string{"Authorization": "Bearer tok", "X-Trace": "1"}}

	// API-facing JSON marshal hides header values (they carry credentials)
	data, err := json.Marshal(c)
	require.NoError(t, err)
	var out struct {
		Headers map[string]string `json:"headers"`
	}
	require.NoError(t, json.Unmarshal(data, &out))
	require.Equal(t, redactedSecret, out.Headers["Authorization"])
	require.Equal(t, redactedSecret, out.Headers["X-Trace"])

	// persistence keeps the real values
	v, err := c.Value()
	require.NoError(t, err)
	var back ClientConfig
	require.NoError(t, back.Scan(v))
	require.Equal(t, "Bearer tok", back.Headers["Authorization"])
	require.Equal(t, "1", back.Headers["X-Trace"])
}

var TestCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))

func TestCIDMarshalBinary(t *testing.T) {
	c := CID(TestCid)
	data, err := c.MarshalBinary()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var c2 CID
	err = c2.UnmarshalBinary(data)
	require.NoError(t, err)
	require.Equal(t, c, c2)
}

func TestCIDMarshalBinary_Empty(t *testing.T) {
	c := CID(cid.Undef)
	data, err := c.MarshalBinary()
	require.NoError(t, err)
	require.Len(t, data, 0)

	var c2 CID
	err = c2.UnmarshalBinary(data)
	require.NoError(t, err)
	require.Equal(t, c, c2)
}

func TestCIDMarshalJSON(t *testing.T) {
	c := CID(TestCid)
	data, err := c.MarshalJSON()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var c2 CID
	err = c2.UnmarshalJSON(data)
	require.NoError(t, err)
	require.Equal(t, c, c2)
}

func TestCIDMarshalJSON_Empty(t *testing.T) {
	c := CID(cid.Undef)
	data, err := c.MarshalJSON()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var c2 CID
	err = c2.UnmarshalJSON(data)
	require.NoError(t, err)
	require.Equal(t, c, c2)
}

func TestCIDStringer(t *testing.T) {
	require.Equal(t, "bafkreie7q3iidccmpvszul7kudcvvuavuo7u6gzlbobczuk5nqk3b4akba", CID(TestCid).String())
	require.Equal(t, "", CID(cid.Undef).String())
}

func TestCIDValueScan(t *testing.T) {
	c := CID(TestCid)
	data, err := c.Value()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var c2 CID
	err = c2.Scan(data)
	require.NoError(t, err)
	require.Equal(t, c, c2)
}

func TestCIDValueScan_Empty(t *testing.T) {
	c := CID(cid.Undef)
	data, err := c.Value()
	require.NoError(t, err)
	require.Len(t, data, 0)

	var c2 CID
	err = c2.Scan(data)
	require.NoError(t, err)
	require.Equal(t, c, c2)
}

func TestStringSliceValueScan(t *testing.T) {
	s := StringSlice{"test"}
	data, err := s.Value()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var s2 StringSlice
	err = s2.Scan(data)
	require.NoError(t, err)
	require.Equal(t, s, s2)
}

func TestStringSliceValueScan_Nil(t *testing.T) {
	s := StringSlice(nil)
	data, err := s.Value()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var s2 StringSlice
	err = s2.Scan(data)
	require.NoError(t, err)
	require.Equal(t, s, s2)
}

func TestStringMapValueScan(t *testing.T) {
	s := ConfigMap{"key": "value"}
	data, err := s.Value()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var s2 ConfigMap
	err = s2.Scan(data)
	require.NoError(t, err)
	require.Equal(t, s, s2)
}

func TestStringMapValueScan_Nil(t *testing.T) {
	s := ConfigMap(nil)
	data, err := s.Value()
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var s2 ConfigMap
	err = s2.Scan(data)
	require.NoError(t, err)
	require.Equal(t, s, s2)
}
