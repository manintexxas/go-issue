
package internal

import (
	"errors"
	"strings"
	"testing"
)

func isolateDict(t *testing.T) {
	t.Helper()

	original := Dict
	Dict = make(map[ShelterID]*ShelterRecord)

	t.Cleanup(func() {
		Dict = original
	})
}

func assertCreatedRecordError(
	t *testing.T,
	source CapacitySource,
	wantErr error,
) {
	t.Helper()

	got, err := CreatedRecord(source)

	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}

	if got != nil {
		t.Error("record should be nil")
	}

	if len(Dict) != 0 {
		t.Errorf("Dict length = %d, want 0", len(Dict))
	}
}

func TestCreatedRecordSuccess(t *testing.T) {
	tests := []struct {
		name    string
		source  CapacitySource
		wantID  ShelterID
		wantCap ShelterCapacity
		wantTag Tags
	}{
		{
			name: "mobile request",
			source: MobileRequest{
				ShelterID: "MBL-0001",
				Capacity:  120,
			},
			wantID:  "MBL-0001",
			wantCap: 120,
			wantTag: MOBILE,
		},
		{
			name: "legacy CSV",
			source: LegacyCSVRow{
				ShelterID: "CSV-0001",
				Capacity:  "250",
			},
			wantID:  "CSV-0001",
			wantCap: 250,
			wantTag: CSV,
		},
		{
			name: "partner payload",
			source: PartnerPayload{
				ShelterID: "NGO-0001",
				Capacity:  uint16(10_000),
			},
			wantID:  "NGO-0001",
			wantCap: 10_000,
			wantTag: NGO,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDict(t)

			got, err := CreatedRecord(tt.source)

			if err != nil {
				t.Fatalf("CreatedRecord() unexpected error: %v", err)
			}

			if got == nil {
				t.Fatal("CreatedRecord() returned nil record")
			}

			if got.ID != tt.wantID {
				t.Errorf("ID = %v, want %v", got.ID, tt.wantID)
			}

			if got.Cap != tt.wantCap {
				t.Errorf("Cap = %v, want %v", got.Cap, tt.wantCap)
			}

			if got.Tag != tt.wantTag {
				t.Errorf("Tag = %v, want %v", got.Tag, tt.wantTag)
			}

			if got.Status != PENDING {
				t.Errorf("Status = %v, want %v", got.Status, PENDING)
			}

			if got.Update.IsZero() {
				t.Error("Update should not be zero")
			}

			if Dict[tt.wantID] != got {
				t.Error("record was not stored correctly in Dict")
			}

			if len(Dict) != 1 {
				t.Errorf("Dict length = %d, want 1", len(Dict))
			}
		})
	}
}

func TestCreatedRecordCapacityTypes(t *testing.T) {
	tests := []struct {
		name string
		cap  any
		want ShelterCapacity
	}{
		{"int8", int8(1), 1},
		{"uint8", uint8(2), 2},
		{"int16", int16(3), 3},
		{"uint16", uint16(4), 4},
		{"int32", int32(5), 5},
		{"uint32", uint32(6), 6},
		{"int", int(7), 7},
		{"uint", uint(8), 8},
		{"int64", int64(9), 9},
		{"uint64", uint64(10), 10},
		{"string", "11", 11},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDict(t)

			id := ShelterID("NGO-1001")

			source := PartnerPayload{
				ShelterID: string(id),
				Capacity:  tt.cap,
			}

			got, err := CreatedRecord(source)

			if err != nil {
				t.Fatalf("CreatedRecord() unexpected error: %v", err)
			}

			if got.Cap != tt.want {
				t.Errorf("Cap = %v, want %v", got.Cap, tt.want)
			}

			if got.ID != id {
				t.Errorf("ID = %v, want %v", got.ID, id)
			}

			if Dict[id] != got {
				t.Error("record was not stored correctly")
			}

			_ = i
		})
	}
}

func TestCreatedRecordEmptySource(t *testing.T) {
	tests := []struct {
		name   string
		source PartnerPayload
	}{
		{
			name: "empty shelter ID",
			source: PartnerPayload{
				ShelterID: nil,
				Capacity:  uint16(100),
			},
		},
		{
			name: "empty capacity",
			source: PartnerPayload{
				ShelterID: "NGO-0002",
				Capacity:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDict(t)
			assertCreatedRecordError(t, tt.source, ErrRequestEmpty)
		})
	}
}

func TestCreatedRecordInvalidID(t *testing.T) {
	tests := []struct {
		name string
		id   any
		want error
	}{
		{"ID is not string", 123, ErrIDFormatInvalid},
		{"invalid ID length", "MBL-001", ErrIDLengthNotValid},
		{"invalid ID prefix", "BAD-0001", ErrIDPrefixInvalid},
		{"ID contains non-number", "MBL-00A1", ErrShelterContainsNonNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDict(t)

			source := PartnerPayload{
				ShelterID: tt.id,
				Capacity:  uint16(100),
			}

			assertCreatedRecordError(t, source, tt.want)
		})
	}
}

func TestCreatedRecordInvalidCapacity(t *testing.T) {
	tests := []struct {
		name string
		cap  any
		want error
	}{
		{"int8 below range", int8(0), ErrShelterCapOutOfRange},
		{"uint8 below range", uint8(0), ErrShelterCapOutOfRange},

		{"int16 below range", int16(-1), ErrShelterCapOutOfRange},
		{"int16 above range", int16(10_001), ErrShelterCapOutOfRange},

		{"uint16 below range", uint16(0), ErrShelterCapOutOfRange},
		{"uint16 above range", uint16(10_001), ErrShelterCapOutOfRange},

		{"int32 below range", int32(-1), ErrShelterCapOutOfRange},
		{"int32 above range", int32(10_001), ErrShelterCapOutOfRange},

		{"uint32 below range", uint32(0), ErrShelterCapOutOfRange},
		{"uint32 above range", uint32(10_001), ErrShelterCapOutOfRange},

		{"int below range", int(-1), ErrShelterCapOutOfRange},
		{"int above range", int(10_001), ErrShelterCapOutOfRange},

		{"uint below range", uint(0), ErrShelterCapOutOfRange},
		{"uint above range", uint(10_001), ErrShelterCapOutOfRange},

		{"int64 below range", int64(-1), ErrShelterCapOutOfRange},
		{"int64 above range", int64(10_001), ErrShelterCapOutOfRange},

		{"uint64 below range", uint64(0), ErrShelterCapOutOfRange},
		{"uint64 above range", uint64(10_001), ErrShelterCapOutOfRange},

		{"empty string", "", ErrShelterCapCannotEmpty},
		{"string zero", "0", ErrShelterCapOutOfRange},
		{"string above range", "10001", ErrShelterCapOutOfRange},
		{"string contains non-number", "12A", ErrShelterContainsNonNumber},
		{"string integer overflow", strings.Repeat("9", 100), ErrIDContainsNonNumber},

		{"unsupported type", true, ErrShelterCapFormatInavalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDict(t)

			source := PartnerPayload{
				ShelterID: "NGO-0003",
				Capacity:  tt.cap,
			}

			assertCreatedRecordError(t, source, tt.want)
		})
	}
}

func TestCreatedRecordDuplicateID(t *testing.T) {
	isolateDict(t)

	id := ShelterID("NGO-0004")

	existing := &ShelterRecord{
		ID:     id,
		Cap:    100,
		Tag:    NGO,
		Status: PENDING,
	}

	Dict[id] = existing

	source := PartnerPayload{
		ShelterID: string(id),
		Capacity:  uint16(200),
	}

	got, err := CreatedRecord(source)

	if !errors.Is(err, ErrIDAlreadyBeenReport) {
		t.Fatalf("error = %v, want %v", err, ErrIDAlreadyBeenReport)
	}

	if got != nil {
		t.Error("record should be nil")
	}

	if Dict[id] != existing {
		t.Error("existing record should not be replaced")
	}

	if len(Dict) != 1 {
		t.Errorf("Dict length = %d, want 1", len(Dict))
	}
}
