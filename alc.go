package nmea

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const (
	// TypeALC type for ALC sentences
	TypeALC = "ALC"
)

type ALCAlertEntry struct {
	MCode         string `json:"m_code,omitempty"`         // manufacture mnemonic code(FEC, null)
	AlertID       string `json:"alert_id,omitempty"`       // alert identifier 000 - 999999
	AlertInstance string `json:"alert_instance,omitempty"` // alert instance null
	Revision      string `json:"revision,omitempty"`       // revision counter 1 - 99
}

// ALC cyclic alert list
type ALC struct {
	BaseSentence `json:"base_sentence,omitempty" mapstructure:"-"`
	TotalNum     int64           `json:"total_num,omitempty" mapstructure:"total_num"`       // total number of sentences this message 01 to 99
	SentenceNum  int64           `json:"sentence_num,omitempty" mapstructure:"sentence_num"` // sentence index number 01-99
	Index        int64           `json:"index,omitempty" mapstructure:"index"`               // sequential message index
	AlertNum     int64           `json:"alert_num,omitempty" mapstructure:"alert_num"`       // Number of alert entries 0 - n
	Alerts       []ALCAlertEntry `json:"alerts,omitempty" mapstructure:"-"`
}

func (s ALC) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := mapstructure.Decode(s, &m)
	if err != nil {
		return map[string]interface{}{}, err
	}
	bm, err := s.BaseSentence.toMap()
	if err != nil {
		return m, err
	}
	for k, v := range bm {
		m[k] = v
	}
	alerts := make([]map[string]interface{}, len(s.Alerts))
	for idx, alert := range s.Alerts {
		alerts[idx] = alert.toMap()
	}
	m["alerts"] = alerts
	return m, nil
}

func (s ALCAlertEntry) toMap() map[string]interface{} {
	return map[string]interface{}{
		"m_code":         s.MCode,
		"alert_id":       s.AlertID,
		"alert_instance": s.AlertInstance,
		"revision":       s.Revision,
	}
}

// newALC constructor
func newALC(s BaseSentence) (ALC, error) {
	p := NewParser(s)
	p.AssertType(TypeALC)

	entries := []ALCAlertEntry{}

	m := ALC{
		BaseSentence: s,
		TotalNum:     p.Int64(0, "TotalNum"),
		SentenceNum:  p.Int64(1, "SentenceNum"),
		Index:        p.Int64(2, "Index"),
		AlertNum:     0,
		Alerts:       entries,
	}

	num := int(p.Int64(3, "AlertNum"))
	err := p.Err()
	if err != nil {
		return m, err
	}
	for i := 0; i < num; i++ {
		idxStr := strconv.FormatInt(int64(i), 10)
		entry := ALCAlertEntry{
			MCode:         p.String(i*4+4+0, "MCode "+idxStr),
			AlertID:       p.String(i*4+4+1, "AlertID "+idxStr),
			AlertInstance: p.String(i*4+4+2, "AlertInstance "+idxStr),
			Revision:      p.String(i*4+4+3, "Revision "+idxStr),
		}
		if p.Err() == nil {
			entries = append(entries, entry)
		} else {
			return m, p.Err()
		}
	}
	m.AlertNum = int64(num)
	m.Alerts = entries

	return m, p.Err()
}
