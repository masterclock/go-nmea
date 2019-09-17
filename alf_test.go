package nmea

import (
	"testing"

	"github.com/go-test/deep"
)

func Test_newALF(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    ALF
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			raw:  makeSentence("$BDALF,1,1,0,012345.78,A,W,V,FEC,999999,null,99,9,alarming"),
			want: ALF{
				TotalNum:    1,
				SentenceNum: 1,
				SeqID:       "0",
				LastChangeTime: Time{
					Valid:       true,
					Hour:        1,
					Minute:      23,
					Second:      45,
					Millisecond: 780,
				},
				AlertCatogory: "A",
				AlertPriority: "W",
				AlertState:    "V",
				MCode:         "FEC",
				AlertID:       "999999",
				AlertInstance: "null",
				Revision:      "99",
				Escalation:    "9",
				AlertText:     "alarming",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newALF() error = %v", err)
				return
			}
			msg := m.(ALF)
			msg.BaseSentence = BaseSentence{}
			if diff := deep.Equal(msg, tt.want); diff != nil {
				t.Errorf("newALF() = %#v, want %#v, dif = %v", msg, tt.want, diff)
			}
		})
	}
}
