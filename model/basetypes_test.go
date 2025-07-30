package model

import (
	"testing"
	"time"

	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestClientConfigMarshal(t *testing.T) {
	c := ClientConfig{
		ConnectTimeout:          ptr.Of(int64(time.Second)),
		Timeout:                 ptr.Of(int64(time.Second)),
		ExpectContinueTimeout:   ptr.Of(int64(time.Second)),
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
		RetryDelay:              ptr.Of(int64(time.Second)),
		RetryBackoff:            ptr.Of(int64(time.Second)),
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
	require.Equal(t, "connectTimeout:1s timeout:1s expectContinueTimeout:1s insecureSkipVerify:true noGzip:true userAgent:x caCert:x clientCert:x clientKey:x headers:<hidden> disableHTTP2true disableHTTPKeepAlives:true retryMaxCount:10 retryDelay:1s retryBackoff:1s retryBackoffExponential:1 skipInaccessibleFile:true useServerModTime:true lowLevelRetries:10 scanConcurrency:10", str)
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
