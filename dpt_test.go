package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newDPT(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    DPT
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDDPT,100,10.1,10000.200"),
			want: DPT{
				Depth:  100,
				Offset: 10.1,
				Range:  10000.200,
			},
			wantErr: false,
		},
		{
			name: "test2",
			raw:  "$SDDPT,,0.0,21.1*49",
			want: DPT{
				Depth:  0,
				Offset: 0,
				Range:  21.1,
			},
		},
		{
			name: "test3",
			raw:  "$SDDPT,10.7,0.0,21.1*51",
			want: DPT{
				Depth:  10.7,
				Offset: 0,
				Range:  21.1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newDPT() error = %v", err)
				return
			}
			msg := m.(DPT)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newDPT() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
