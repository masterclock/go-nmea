package nmea

import (
	"reflect"
	"testing"
)

func Test_newALC(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    ALC
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "0 alert entry",
			raw:  makeSentence("$BDALC,01,01,00,0"),
			want: ALC{
				TotalNum:    1,
				SentenceNum: 1,
				Index:       0,
				AlertNum:    0,
				Alerts:      []ALCAlertEntry{},
			},
			wantErr: false,
		},
		{
			name: "1 alert entry",
			raw:  makeSentence("$BDALC,01,01,00,1,FEC,000,null,1"),
			want: ALC{
				TotalNum:    1,
				SentenceNum: 1,
				Index:       0,
				AlertNum:    1,
				Alerts: []ALCAlertEntry{
					{
						MCode:         "FEC",
						AlertID:       "000",
						AlertInstance: "null",
						Revision:      "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple alert entry",
			raw:  makeSentence("$BDALC,01,01,00,3,FEC,000,null,1,FEC,001,null,2,FEC,999999,null,99"),
			want: ALC{
				TotalNum:    1,
				SentenceNum: 1,
				Index:       0,
				AlertNum:    3,
				Alerts: []ALCAlertEntry{
					{
						MCode:         "FEC",
						AlertID:       "000",
						AlertInstance: "null",
						Revision:      "1",
					},
					{
						MCode:         "FEC",
						AlertID:       "001",
						AlertInstance: "null",
						Revision:      "2",
					},
					{
						MCode:         "FEC",
						AlertID:       "999999",
						AlertInstance: "null",
						Revision:      "99",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if err != nil {
				t.Errorf("newACL() error = %v", err)
				return
			}
			msg := m.(ALC)
			msg.BaseSentence = BaseSentence{}
			if !reflect.DeepEqual(msg, tt.want) {
				t.Errorf("newALC() = %v, want %v", msg, tt.want)
			}
		})
	}
}
