package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newDBK(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    DBK
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDDBK,10,f,100,M,1000,F"),
			want: DBK{
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
				t.Errorf("newDBK() error = %v", err)
				return
			}
			msg := m.(DBK)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newDBK() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
