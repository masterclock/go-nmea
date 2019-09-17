package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newHBT(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    HBT
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDHBT,100.1,A,9"),
			want: HBT{
				Interval: 100.1,
				Status:   "A",
				ID:       "9",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newHBT() error = %v", err)
				return
			}
			msg := m.(HBT)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newHBT() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
