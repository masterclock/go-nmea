package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newDBS(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    DBS
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDDBS,10,f,100,M,1000,F"),
			want: DBS{
				DepthFeet:   10,
				Feet:        "f",
				DepthMeters: 100,
				Meters:      "M",
				DepthFathom: 1000,
				Fathom:      "F",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newDBS() error = %v", err)
				return
			}
			msg := m.(DBS)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newDBS() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
