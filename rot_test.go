package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newROT(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    ROT
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDROT,100.1,A"),
			want: ROT{
				Rate:   100.1,
				Status: "A",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newROT() error = %v", err)
				return
			}
			msg := m.(ROT)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newROT() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
