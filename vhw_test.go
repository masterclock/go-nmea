package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newVHW(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    VHW
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDVHW,1.1,T,2.2,M,3.3,N,4.4,K"),
			want: VHW{
				HeadingTrue:     1.1,
				True:            "T",
				HeadingMagnetic: 2.2,
				Magnetic:        "M",
				SpeedKnots:      3.3,
				Knots:           "N",
				SpeedKph:        4.4,
				Kph:             "K",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newVHW() error = %v", err)
				return
			}
			msg := m.(VHW)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newVHW() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
