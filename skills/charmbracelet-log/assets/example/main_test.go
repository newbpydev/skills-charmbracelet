package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStructuredFields(t *testing.T) {
	var out bytes.Buffer
	newLogger(&out).With("request", "r1").Info("loaded", "count", 3)
	var row map[string]any
	if err := json.Unmarshal(out.Bytes(), &row); err != nil {
		t.Fatal(err)
	}
	if row["request"] != "r1" || row["count"] != float64(3) {
		t.Fatalf("fields lost: %v", row)
	}
	if bytes.Contains(out.Bytes(), []byte("\x1b")) {
		t.Fatal("JSON contains terminal styling")
	}
}
