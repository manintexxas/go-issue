package internal

import (
	"errors"
	"reflect"
	"testing"
)

// ============================================================
// CHANNEL CONFIGURATION API
// ============================================================

func TestConfig(t *testing.T) {
	initial := Channel{true, false, true, false}

	tests := []struct {
		name    string
		index   int
		enabled bool
		want    Channel
		wantErr error
	}{
		{
			name:    "enable SMS",
			index:   0,
			enabled: true,
			want:    Channel{true, false, true, false},
		},
		{
			name:    "disable Cell Broadcast",
			index:   1,
			enabled: false,
			want:    Channel{true, false, true, false},
		},
		{
			name:    "enable Radio",
			index:   2,
			enabled: true,
			want:    Channel{true, false, true, false},
		},
		{
			name:    "enable Satellite",
			index:   3,
			enabled: true,
			want:    Channel{true, false, true, true},
		},
		{
			name:    "negative channel index",
			index:   -1,
			enabled: true,
			want:    initial,
			wantErr: ErrChannelIndexOutOfRange,
		},
		{
			name:    "channel index overflow",
			index:   4,
			enabled: true,
			want:    initial,
			wantErr: ErrChannelIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := initial

			got, err := Config(original, tt.index, tt.enabled)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected channel %v, got %v", tt.want, got)
			}

			if !reflect.DeepEqual(original, initial) {
				t.Errorf("original channel was modified: got %v", original)
			}
		})
	}
}

// ============================================================
// RECIPIENT QUEUE API
// ============================================================

func TestRecipients(t *testing.T) {
	t.Run("append recipients and preserve order", func(t *testing.T) {
		queue := Queue{"REG-001", "REG-002"}
		incoming := Incoming{"REG-003", "REG-004"}

		got := Recipients(queue, incoming)

		want := Queue{
			"REG-001",
			"REG-002",
			"REG-003",
			"REG-004",
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("expected queue %v, got %v", want, got)
		}
	})

	t.Run("append with sufficient capacity", func(t *testing.T) {
		queue := make(Queue, 2, 4)
		queue[0] = "REG-001"
		queue[1] = "REG-002"

		got := Recipients(queue, Incoming{"REG-003"})

		if len(got) != 3 {
			t.Fatalf("expected length 3, got %d", len(got))
		}

		// The returned slice shares the original backing array.
		got[0] = "REG-UPDATED"

		if queue[0] != "REG-UPDATED" {
			t.Errorf("expected shared backing array, got queue %v", queue)
		}

		// Inspect the original backing array beyond its current length.
		expandedQueue := queue[:cap(queue)]

		if expandedQueue[2] != "REG-003" {
			t.Errorf("expected appended value in shared backing array, got %v", expandedQueue)
		}
	})

	t.Run("append with insufficient capacity", func(t *testing.T) {
		queue := make(Queue, 2, 2)
		queue[0] = "REG-001"
		queue[1] = "REG-002"

		got := Recipients(queue, Incoming{"REG-003"})

		if len(got) != 3 {
			t.Fatalf("expected length 3, got %d", len(got))
		}

		// Append must allocate a new backing array.
		got[0] = "REG-UPDATED"

		if queue[0] != "REG-001" {
			t.Errorf("expected original queue to remain unchanged, got %v", queue)
		}
	})

	t.Run("empty incoming", func(t *testing.T) {
		queue := Queue{"REG-001", "REG-002"}

		got := Recipients(queue, Incoming{})

		if !reflect.DeepEqual(got, queue) {
			t.Errorf("expected unchanged queue %v, got %v", queue, got)
		}
	})

	t.Run("nil queue and nil incoming", func(t *testing.T) {
		var queue Queue
		var incoming Incoming

		got := Recipients(queue, incoming)

		if got != nil {
			t.Errorf("expected nil queue, got %v", got)
		}
	})

	t.Run("append to nil queue", func(t *testing.T) {
		var queue Queue
		incoming := Incoming{"REG-001", "REG-002"}

		got := Recipients(queue, incoming)

		want := Queue{"REG-001", "REG-002"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected queue %v, got %v", want, got)
		}
	})
}

// ============================================================
// LIVE PAGE VIEW API
// ============================================================

func TestView(t *testing.T) {
	t.Run("create page view with shared backing array", func(t *testing.T) {
		queue := Queue{
			"REG-001",
			"REG-002",
			"REG-003",
			"REG-004",
		}

		got, err := View(queue, 1, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := Queue{"REG-002", "REG-003"}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("expected view %v, got %v", want, got)
		}

		if len(got) != 2 {
			t.Errorf("expected length 2, got %d", len(got))
		}

		if cap(got) != 2 {
			t.Errorf("expected capacity 2, got %d", cap(got))
		}

		// Element mutation must be reflected in the original queue.
		got[0] = "REG-UPDATED"

		if queue[1] != "REG-UPDATED" {
			t.Errorf("expected queue element to change, got %v", queue)
		}

		// Append must not overwrite elements outside the page.
		got = append(got, "REG-005")

		if queue[3] != "REG-004" {
			t.Errorf("append modified an element outside the page: %v", queue)
		}
	})

	t.Run("empty page at queue boundary", func(t *testing.T) {
		queue := Queue{"REG-001", "REG-002"}

		got, err := View(queue, 2, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 0 || cap(got) != 0 {
			t.Errorf("expected empty view with zero capacity, got len=%d cap=%d", len(got), cap(got))
		}
	})

	t.Run("nil queue", func(t *testing.T) {
		var queue Queue

		got, err := View(queue, 0, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 0 {
			t.Errorf("expected empty view, got %v", got)
		}
	})

	tests := []struct {
		name    string
		start   int
		end     int
		wantErr error
	}{
		{
			name:    "negative start",
			start:   -1,
			end:     1,
			wantErr: ErrIndexCannotBeNegative,
		},
		{
			name:    "negative end",
			start:   0,
			end:     -1,
			wantErr: ErrIndexCannotBeNegative,
		},
		{
			name:    "start overflow",
			start:   5,
			end:     5,
			wantErr: ErrIndexOverflow,
		},
		{
			name:    "end overflow",
			start:   0,
			end:     5,
			wantErr: ErrIndexOverflow,
		},
		{
			name:    "invalid range",
			start:   3,
			end:     2,
			wantErr: ErrIndexInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := Queue{"REG-001", "REG-002", "REG-003", "REG-004"}

			_, err := View(queue, tt.start, tt.end)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

// ============================================================
// SNAPSHOT API
// ============================================================

func TestAudit(t *testing.T) {
	t.Run("snapshot has independent backing array", func(t *testing.T) {
		queue := Queue{
			"REG-001",
			"REG-002",
			"REG-003",
			"REG-004",
		}

		got, err := Audit(queue, 1, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := SnapShot{"REG-002", "REG-003"}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("expected snapshot %v, got %v", want, got)
		}

		// Snapshot mutation must not affect the original queue.
		got[0] = "SNAPSHOT-UPDATED"

		if queue[1] != "REG-002" {
			t.Errorf("snapshot mutation affected original queue: %v", queue)
		}

		// Original queue mutation must not affect the snapshot.
		queue[2] = "QUEUE-UPDATED"

		if got[1] != "REG-003" {
			t.Errorf("queue mutation affected snapshot: %v", got)
		}
	})

	t.Run("snapshot with empty range", func(t *testing.T) {
		queue := Queue{"REG-001", "REG-002"}

		got, err := Audit(queue, 2, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 0 {
			t.Errorf("expected empty snapshot, got %v", got)
		}
	})

	t.Run("snapshot from nil queue", func(t *testing.T) {
		var queue Queue

		got, err := Audit(queue, 0, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 0 {
			t.Errorf("expected empty snapshot, got %v", got)
		}
	})

	tests := []struct {
		name    string
		start   int
		end     int
		wantErr error
	}{
		{
			name:    "negative start",
			start:   -1,
			end:     1,
			wantErr: ErrIndexCannotBeNegative,
		},
		{
			name:    "negative end",
			start:   0,
			end:     -1,
			wantErr: ErrIndexCannotBeNegative,
		},
		{
			name:    "start overflow",
			start:   5,
			end:     5,
			wantErr: ErrIndexOverflow,
		},
		{
			name:    "end overflow",
			start:   0,
			end:     5,
			wantErr: ErrIndexOverflow,
		},
		{
			name:    "invalid range",
			start:   3,
			end:     2,
			wantErr: ErrIndexInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := Queue{"REG-001", "REG-002", "REG-003", "REG-004"}

			_, err := Audit(queue, tt.start, tt.end)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
