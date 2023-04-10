package agi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParseTasksFromString(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    Tasks
		wantErr bool
	}{
		{
			name: "empty",
			s:    "",
			want: nil,
		},
		{
			name: "single task",
			s:    `1. task1`,
			want: Tasks{
				{
					ID:   MakeTaskID(),
					Name: "task1",
				},
			},
		},
		{
			name: "multiple tasks",
			s: `1. task1
2. task2
3. task3`,
			want: Tasks{
				{
					ID:   MakeTaskID(),
					Name: "task1",
				},
				{
					ID:   MakeTaskID(),
					Name: "task2",
				},
				{
					ID:   MakeTaskID(),
					Name: "task3",
				},
			},
		},
		{
			name: "multiple tasks with empty lines",
			s: `1. task1

2. task2

3. task3`,
			want: Tasks{
				{
					ID:   MakeTaskID(),
					Name: "task1",
				},
				{
					ID:   MakeTaskID(),
					Name: "task2",
				},
				{
					ID:   MakeTaskID(),
					Name: "task3",
				},
			},
		},
		{
			name: "multiple tasks with empty lines and spaces",
			s: `1. task1

2. task2

3. task3`,
			want: Tasks{
				{
					ID:   MakeTaskID(),
					Name: "task1",
				},
				{
					ID:   MakeTaskID(),
					Name: "task2",
				},
				{
					ID:   MakeTaskID(),
					Name: "task3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTasksFromString(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTasksFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, cmpopts.IgnoreFields(Task{}, "ID")) {
				t.Errorf("ParseTasksFromString() = %v, want %v; diff %s", got, tt.want, cmp.Diff(tt.want, got))
			}
		})
	}
}
