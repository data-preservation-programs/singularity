package statechange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatOptionalInt32(t *testing.T) {
	// Test with nil value
	assert.Equal(t, "", formatOptionalInt32(nil))

	// Test with valid value
	value := int32(12345)
	assert.Equal(t, "12345", formatOptionalInt32(&value))

	// Test with negative value
	negValue := int32(-678)
	assert.Equal(t, "-678", formatOptionalInt32(&negValue))

	// Test with zero
	zeroValue := int32(0)
	assert.Equal(t, "0", formatOptionalInt32(&zeroValue))
}

func TestFormatOptionalString(t *testing.T) {
	// Test with nil value
	assert.Equal(t, "", formatOptionalString(nil))

	// Test with valid value
	value := "test-string"
	assert.Equal(t, "test-string", formatOptionalString(&value))

	// Test with empty string
	emptyValue := ""
	assert.Equal(t, "", formatOptionalString(&emptyValue))
}

func TestExportStateChanges_UnsupportedFormat(t *testing.T) {
	err := exportStateChanges(nil, "xml", "test.xml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported export format: xml")
}

// Basic smoke tests to ensure commands compile and have expected structure
func TestCommandsExist(t *testing.T) {
	assert.NotNil(t, ListCmd)
	assert.NotNil(t, GetCmd)
	assert.NotNil(t, StatsCmd)
	assert.NotNil(t, RepairCmd)

	// Verify command names
	assert.Equal(t, "list", ListCmd.Name)
	assert.Equal(t, "get", GetCmd.Name)
	assert.Equal(t, "stats", StatsCmd.Name)
	assert.Equal(t, "repair", RepairCmd.Name)

	// Verify repair subcommands exist
	assert.Greater(t, len(RepairCmd.Subcommands), 0)
}