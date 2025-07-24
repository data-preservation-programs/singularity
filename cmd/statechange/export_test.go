package statechange

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExportToCSV(t *testing.T) {
	// Create test data
	epochHeight := int32(123456)
	sectorID := "sector-123"
	stateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: "proposed",
			NewState:      "published",
			Timestamp:     time.Date(2023, 6, 15, 10, 30, 0, 0, time.UTC),
			EpochHeight:   &epochHeight,
			SectorID:      &sectorID,
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      `{"reason":"test"}`,
		},
		{
			ID:            2,
			DealID:        model.DealID(456),
			PreviousState: "published",
			NewState:      "active",
			Timestamp:     time.Date(2023, 6, 16, 11, 45, 0, 0, time.UTC),
			EpochHeight:   nil,
			SectorID:      nil,
			ProviderID:    "f05678",
			ClientAddress: "f1fedcba",
			Metadata:      "{}",
		},
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "test-export-*.csv")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Export to CSV
	err = exportToCSV(stateChanges, tmpFile.Name())
	require.NoError(t, err)

	// Read and verify CSV content
	file, err := os.Open(tmpFile.Name())
	require.NoError(t, err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	require.NoError(t, err)

	// Verify header
	expectedHeader := []string{
		"ID", "DealID", "PreviousState", "NewState", "Timestamp",
		"EpochHeight", "SectorID", "ProviderID", "ClientAddress", "Metadata",
	}
	assert.Equal(t, expectedHeader, records[0])

	// Verify first data row
	expectedRow1 := []string{
		"1", "123", "proposed", "published", "2023-06-15 10:30:00",
		"123456", "sector-123", "f01234", "f1abcdef", `{"reason":"test"}`,
	}
	assert.Equal(t, expectedRow1, records[1])

	// Verify second data row (with nil values)
	expectedRow2 := []string{
		"2", "456", "published", "active", "2023-06-16 11:45:00",
		"", "", "f05678", "f1fedcba", "{}",
	}
	assert.Equal(t, expectedRow2, records[2])

	// Should have header + 2 data rows
	assert.Len(t, records, 3)
}

func TestExportToJSON(t *testing.T) {
	// Create test data
	epochHeight := int32(123456)
	sectorID := "sector-123"
	stateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: "proposed",
			NewState:      "published",
			Timestamp:     time.Date(2023, 6, 15, 10, 30, 0, 0, time.UTC),
			EpochHeight:   &epochHeight,
			SectorID:      &sectorID,
			ProviderID:    "f01234",
			ClientAddress: "f1abcdef",
			Metadata:      `{"reason":"test"}`,
		},
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "test-export-*.json")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Export to JSON
	err = exportToJSON(stateChanges, tmpFile.Name())
	require.NoError(t, err)

	// Read and verify JSON content
	file, err := os.Open(tmpFile.Name())
	require.NoError(t, err)
	defer file.Close()

	var exportData struct {
		Metadata struct {
			ExportTime string `json:"exportTime"`
			TotalCount int    `json:"totalCount"`
		} `json:"metadata"`
		StateChanges []model.DealStateChange `json:"stateChanges"`
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&exportData)
	require.NoError(t, err)

	// Verify metadata
	assert.NotEmpty(t, exportData.Metadata.ExportTime)
	assert.Equal(t, 1, exportData.Metadata.TotalCount)

	// Verify state changes
	require.Len(t, exportData.StateChanges, 1)
	assert.Equal(t, uint64(1), exportData.StateChanges[0].ID)
	assert.Equal(t, model.DealID(123), exportData.StateChanges[0].DealID)
	assert.Equal(t, model.DealState("proposed"), exportData.StateChanges[0].PreviousState)
	assert.Equal(t, model.DealState("published"), exportData.StateChanges[0].NewState)
	assert.Equal(t, "f01234", exportData.StateChanges[0].ProviderID)
	assert.Equal(t, "f1abcdef", exportData.StateChanges[0].ClientAddress)
}

func TestExportStateChanges_UnsupportedFormat(t *testing.T) {
	stateChanges := []model.DealStateChange{}
	
	err := exportStateChanges(stateChanges, "xml", "test.xml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported export format: xml")
}

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

func TestExportToCSV_EmptyData(t *testing.T) {
	// Create temporary file
	tmpFile, err := os.CreateTemp("", "test-export-empty-*.csv")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Export empty data
	err = exportToCSV([]model.DealStateChange{}, tmpFile.Name())
	require.NoError(t, err)

	// Read and verify CSV content
	file, err := os.Open(tmpFile.Name())
	require.NoError(t, err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	require.NoError(t, err)

	// Should only have header row
	assert.Len(t, records, 1)
	expectedHeader := []string{
		"ID", "DealID", "PreviousState", "NewState", "Timestamp",
		"EpochHeight", "SectorID", "ProviderID", "ClientAddress", "Metadata",
	}
	assert.Equal(t, expectedHeader, records[0])
}

func TestExportToJSON_EmptyData(t *testing.T) {
	// Create temporary file
	tmpFile, err := os.CreateTemp("", "test-export-empty-*.json")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Export empty data
	err = exportToJSON([]model.DealStateChange{}, tmpFile.Name())
	require.NoError(t, err)

	// Read and verify JSON content
	file, err := os.Open(tmpFile.Name())
	require.NoError(t, err)
	defer file.Close()

	var exportData struct {
		Metadata struct {
			ExportTime string `json:"exportTime"`
			TotalCount int    `json:"totalCount"`
		} `json:"metadata"`
		StateChanges []model.DealStateChange `json:"stateChanges"`
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&exportData)
	require.NoError(t, err)

	// Verify metadata
	assert.NotEmpty(t, exportData.Metadata.ExportTime)
	assert.Equal(t, 0, exportData.Metadata.TotalCount)

	// Verify empty state changes
	assert.Len(t, exportData.StateChanges, 0)
}