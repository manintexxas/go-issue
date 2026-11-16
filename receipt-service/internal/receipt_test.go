package internal_test

import (
	"testing"

	in "receipt-service/internal"

)

func TestBuild(t *testing.T) {
	tests := []struct {
		name	string
		id		in.ID
		status	in.Status
		want 	string
		wantErr	bool
	}{
		{
		name:	"paid",
		id:		in.ID("INV-0001"),
		status:	in.PAID,
		want:	"Invoice: INV-0001\nStatus: PAID",
		},
		{
		name:	"pending",
		id:		in.ID("INV-0001"),
		status:	in.PENDING,
		want:	"Invoice: INV-0001\nStatus: PENDING",
		},
		{
		name:	"failed",
		id:		in.ID("INV-0001"),
		status:	in.FAILED,
		want:	"Invoice: INV-0001\nStatus: FAILED",
		},
		{
		name:	"idTooShort",
		id:		in.ID("INV-001"),
		status:	in.PAID,
		wantErr:true,
		},
		{
		name:	"idEmpty",
		id:		"",
		status:	in.PAID,
		wantErr:true,
		},
		{
		name:	"unkownStatus",
		id:		in.ID("INV-0001"),
		status:	in.UNKNOWN,
		wantErr:true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := in.Build(tt.id, tt.status)

			if tt.want != "" {
				if err != nil {
					t.Errorf("%v", err)
				}

				if got != tt.want {
					t.Errorf("\ngot: %q\nwant: %q", got, tt.want)
				}
			}

			if tt.wantErr {
				if err == nil {
					t.Errorf("Unexpected success: %q", got)
				}
			}
		})
	}
	
}