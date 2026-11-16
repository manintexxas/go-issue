package internal_test

import (
	"errors"
	"testing"

	in "receipt-service/internal"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name    string
		id      in.ID
		status  in.Status
		want    string
		wantErr error
	}{
		{
			name:   "paid",
			id:     in.ID("INV-0001"),
			status: in.PAID,
			want:   "Invoice: INV-0001\nStatus: PAID",
		},
		{
			name:   "pending",
			id:     in.ID("INV-0001"),
			status: in.PENDING,
			want:   "Invoice: INV-0001\nStatus: PENDING",
		},
		{
			name:   "failed",
			id:     in.ID("INV-0001"),
			status: in.FAILED,
			want:   "Invoice: INV-0001\nStatus: FAILED",
		},
		{
			name:    "idTooShort",
			id:      in.ID("INV-001"),
			status:  in.PAID,
			wantErr: in.ErrIDLengthNotValid,
		},
		{
			name:    "idEmpty",
			id:      "",
			status:  in.PAID,
			wantErr: in.ErrIDLengthNotValid,
		},
		{
			name:    "invalidPrefix",
			id:      in.ID("ABC-0001"),
			status:  in.PAID,
			wantErr: in.ErrIDPrefixNotValid,
		},
		{
			name:    "nonDigitID",
			id:      in.ID("INV-00A1"),
			status:  in.PAID,
			wantErr: in.ErrIDContainsNonNumber,
		},
		{
			name:    "unknownStatus",
			id:      in.ID("INV-0001"),
			status:  in.UNKNOWN,
			wantErr: in.ErrStatusNotValid,
		},
		{
			name:    "zeroStatus",
			id:      in.ID("INV-0001"),
			status:  in.Status(0),
			wantErr: in.ErrStatusNotValid,
		},
		{
			name:    "combinedStatus",
			id:      in.ID("INV-0001"),
			status:  in.Status(3),
			wantErr: in.ErrStatusNotValid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := in.Build(tt.id, tt.status)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf(
					"Build() error = %v, want error = %v",
					err,
					tt.wantErr,
				)
			}

			if tt.wantErr != nil {
				if got != "" {
					t.Errorf("Build() result = %q, want empty result", got)
				}
				return
			}

			if got != tt.want {
				t.Errorf("Build() = %q, want %q", got, tt.want)
			}
		})
	}
}