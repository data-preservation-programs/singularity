package statechange

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
)

// exportStateChanges exports state changes to the specified format and file path
func exportStateChanges(stateChanges []model.DealStateChange, format, outputPath string) error {
	switch format {
	case "csv":
		return exportToCSV(stateChanges, outputPath)
	case "json":
		return exportToJSON(stateChanges, outputPath)
	default:
		return errors.Errorf("unsupported export format: %s", format)
	}
}

// exportToCSV exports state changes to a CSV file
func exportToCSV(stateChanges []model.DealStateChange, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "failed to create CSV file")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{
		"ID",
		"DealID",
		"PreviousState",
		"NewState",
		"Timestamp",
		"EpochHeight",
		"SectorID",
		"ProviderID",
		"ClientAddress",
		"Metadata",
	}
	if err := writer.Write(header); err != nil {
		return errors.Wrap(err, "failed to write CSV header")
	}

	// Write state change records
	for _, change := range stateChanges {
		record := []string{
			strconv.FormatUint(change.ID, 10),
			strconv.FormatUint(uint64(change.DealID), 10),
			string(change.PreviousState),
			string(change.NewState),
			change.Timestamp.Format("2006-01-02 15:04:05"),
			formatOptionalInt32(change.EpochHeight),
			formatOptionalString(change.SectorID),
			change.ProviderID,
			change.ClientAddress,
			change.Metadata,
		}
		if err := writer.Write(record); err != nil {
			return errors.Wrap(err, "failed to write CSV record")
		}
	}

	return nil
}

// exportToJSON exports state changes to a JSON file
func exportToJSON(stateChanges []model.DealStateChange, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "failed to create JSON file")
	}
	defer file.Close()

	// Create export structure with metadata
	exportData := struct {
		Metadata struct {
			ExportTime string `json:"exportTime"`
			TotalCount int    `json:"totalCount"`
		} `json:"metadata"`
		StateChanges []model.DealStateChange `json:"stateChanges"`
	}{
		StateChanges: stateChanges,
	}

	exportData.Metadata.ExportTime = time.Now().Format(time.RFC3339)
	exportData.Metadata.TotalCount = len(stateChanges)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(exportData); err != nil {
		return errors.Wrap(err, "failed to encode JSON")
	}

	return nil
}

// Helper functions for formatting optional fields
func formatOptionalInt32(value *int32) string {
	if value == nil {
		return ""
	}
	return strconv.FormatInt(int64(*value), 10)
}

func formatOptionalString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
