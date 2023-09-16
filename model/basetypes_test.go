package model

import (
	"testing"

	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

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
