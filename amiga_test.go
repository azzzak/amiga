package amiga

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulate(t *testing.T) {
	var holder [3]string

	type args struct {
		m map[string]string
		p map[string]*string
	}
	tests := []struct {
		name    string
		args    args
		want    [3]string
		wantErr error
	}{
		{
			name: "full",
			args: args{
				m: map[string]string{
					"Event":        "AgentComplete",
					"Privilege":    "agent,all",
					"DestPriority": "1",
				},
				p: map[string]*string{
					"Event":        &holder[0],
					"Privilege":    &holder[1],
					"DestPriority": &holder[2],
				},
			},
			want: [3]string{"AgentComplete", "agent,all", "1"},
		}, {
			name: "partial",
			args: args{
				m: map[string]string{
					"Event":        "AgentComplete",
					"Privilege":    "agent,all",
					"DestPriority": "1",
				},
				p: map[string]*string{
					"Event": &holder[0],
				},
			},
			want: [3]string{"AgentComplete", "", ""},
		}, {
			name: "error",
			args: args{
				m: map[string]string{
					"Event":        "AgentComplete",
					"Privilege":    "agent,all",
					"DestPriority": "1",
				},
				p: map[string]*string{
					"Reason":       &holder[1],
					"DestPriority": &holder[2],
				},
			},
			want:    [3]string{"", "", "1"},
			wantErr: fmt.Errorf("missing arguments: %s", "Reason"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Populate(tt.args.m, tt.args.p)
			assert.Equal(t, tt.wantErr, err)

			for n := range []int{0, 1, 2} {
				assert.Equal(t, tt.want[n], holder[n])
			}

			holder = [3]string{}
		})
	}
}
