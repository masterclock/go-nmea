package nmea

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
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

func TestALC_ToMap(t *testing.T) {
	type fields struct {
		BaseSentence BaseSentence
		TotalNum     int64
		SentenceNum  int64
		Index        int64
		AlertNum     int64
		Alerts       []ALCAlertEntry
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				BaseSentence: BaseSentence{
					Talker:   "BD",
					Type:     "ALC",
					Fields:   []string{},
					Checksum: "FF",
					Raw:      "Raw",
				},
				TotalNum:    10,
				SentenceNum: 11,
				Index:       1,
				AlertNum:    2,
				Alerts: []ALCAlertEntry{
					{
						MCode:         "FEC",
						AlertID:       "0001",
						AlertInstance: "null",
						Revision:      "9",
					},
					{
						MCode:         "FEC2",
						AlertID:       "0002",
						AlertInstance: "null2",
						Revision:      "92",
					},
				},
			},
			want: map[string]interface{}{
				"talker":       "BD",
				"type":         "ALC",
				"fields":       []string{},
				"checksum":     "FF",
				"raw":          "Raw",
				"total_num":    int64(10),
				"sentence_num": int64(11),
				"index":        int64(1),
				"alert_num":    int64(2),
				"alerts": []map[string]interface{}{
					{
						"m_code":         "FEC",
						"alert_id":       "0001",
						"alert_instance": "null",
						"revision":       "9",
					},
					{
						"m_code":         "FEC2",
						"alert_id":       "0002",
						"alert_instance": "null2",
						"revision":       "92",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ALC{
				BaseSentence: tt.fields.BaseSentence,
				TotalNum:     tt.fields.TotalNum,
				SentenceNum:  tt.fields.SentenceNum,
				Index:        tt.fields.Index,
				AlertNum:     tt.fields.AlertNum,
				Alerts:       tt.fields.Alerts,
			}
			got, err := s.ToMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("ALC.ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("ALC.ToMap() = %v, want %v, diff = %v", got, tt.want, diff)
			}
		})
	}
}
