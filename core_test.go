package amiga

import (
	"bytes"
	"net/textproto"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_payload(t *testing.T) {
	type args struct {
		action string
		m      map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_1",
			args: args{
				action: "Originate",
				m: map[string]string{
					"Channel":     "Local/loop@test",
					"Application": "Playback",
					"Data":        "demo-congrats",
				},
			},
			want: "Action: Originate\r\nChannel: Local/loop@test\r\nApplication: Playback\r\nData: demo-congrats\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := payload(tt.args.action, tt.args.m)
			actual := strings.SplitAfter(got, "\r\n")
			expected := strings.SplitAfter(tt.want, "\r\n")
			sort.Strings(actual)
			sort.Strings(expected)
			assert.Equal(t, expected, actual)
		})
	}
}

type mock struct {
	bytes.Buffer
}

func (m mock) Close() error {
	return nil
}

func Test_dispatcher(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name:   "test_1",
			source: "Response: Success\r\nActionID: 25df-cf70-1228\r\n\r\nResponse: Success\r\nActionID: 7a67-b519-09fa\r\n\r\nActionID: f461-f3f5-f0a9\r\nResponse: Success\r\n\r\nResponse: Success\r\nActionID: d127-5966-7cd4\r\n\r\n",
		},
	}
	for _, tt := range tests {

		ami := New()

		t.Run(tt.name, func(t *testing.T) {
			c1 := make(chan map[string]string)
			c2 := make(chan map[string]string)
			c3 := make(chan map[string]string)

			ami.actions = map[string]chan map[string]string{
				"d127-5966-7cd4": c3,
				"f461-f3f5-f0a9": c2,
				"25df-cf70-1228": c1,
			}

			mock := new(mock)
			conn := textproto.NewConn(mock)

			mock.WriteString(tt.source)
			go dispatcher(ami, conn)

			var (
				res []string
				v   map[string]string
			)
			i := 3

			for {
				select {
				case v = <-c1:
				case v = <-c2:
				case v = <-c3:
				}
				res = append(res, v["ActionID"])
				i--
				if i == 0 {
					break
				}
			}

			assert.Equal(t, []string{
				"25df-cf70-1228",
				"f461-f3f5-f0a9",
				"d127-5966-7cd4",
			}, res)
		})
	}
}

func Test_readFields(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		want    map[string]string
		wantErr bool
	}{
		{
			name:   "test_1",
			source: "Event: AgentComplete\r\nActionID: 25df-cf70-1228\r\nPrivilege: agent,all\r\nChannelStateDesc: Up\r\n\r\n",
			want: map[string]string{
				"Event":            "AgentComplete",
				"Privilege":        "agent,all",
				"ChannelStateDesc": "Up",
				"ActionID":         "25df-cf70-1228",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := new(mock)
			conn := textproto.NewConn(mock)

			mock.WriteString(tt.source)
			got, err := readFields(conn)

			if assert.Nil(t, err) == tt.wantErr {
				t.Errorf("readFields() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, got, tt.want)
		})
	}
}

var event = map[string]string{
	// "Event":                "AgentComplete",
	"Privilege":            "agent,all",
	"Channel":              "Console/dsp",
	"ChannelState":         "6",
	"ChannelStateDesc":     "Up",
	"CallerIDNum":          "<unknown>",
	"CallerIDName":         "<unknown>",
	"ConnectedLineNum":     "103",
	"Language":             "en",
	"Context":              "queues",
	"Exten":                "3011",
	"Priority":             "6",
	"DestChannelState":     "6",
	"DestChannelStateDesc": "Up",
	"DestCallerIDNum":      "103",
	"DestLanguage":         "ru",
	"DestContext":          "full",
	"DestExten":            "3011",
	"DestPriority":         "1",
	"DestUniqueid":         "1608123110.45",
	"DestLinkedid":         "1608123109.42",
	"Queue":                "queue3011",
	"Interface":            "SIP/103",
	"MemberName":           "103",
	"HoldTime":             "2",
	"TalkTime":             "9",
	"Reason":               "caller",
}

func BenchmarkPayload(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = payload("AgentComplete", event)
	}
}
