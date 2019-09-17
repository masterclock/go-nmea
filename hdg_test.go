package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newHDG(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    HDG
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDHDG,5.0,100.1,E,9.00,W"),
			want: HDG{
				Heading:            5.0,
				Deviation:          100.1,
				DeviationDirection: "E",
				Variation:          9.00,
				VariationDirection: "W",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newHDG() error = %v", err)
				return
			}
			msg := m.(HDG)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newHDG() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
