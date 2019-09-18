package nmea

import "github.com/mitchellh/mapstructure"

const (
	// TypeALF type for ALF sentences
	TypeALF = "ALF"
)

// ALF alert sentence
// http://aprs.gids.nl/nmea/#hdt
type ALF struct {
	BaseSentence   `json:"base_sentence,omitempty" mapstrucure:"-"`
	TotalNum       int64  `json:"total_num,omitempty" mapstrucure:"total_num,omitempty"`                // total number of alf senteces this message (1 - 2)
	SentenceNum    int64  `mapstructure:"sentence_num,omitempty" json:"sentence_num,omitempty"`         // sentence number (1 - 2)
	SeqID          string `mapstructure:"seq_id,omitempty" json:"seq_id,omitempty"`                     // sequential message identifier (0 - 9)
	LastChangeTime Time   `mapstructure:"last_change_time,omitempty" json:"last_change_time,omitempty"` // time of last change, hhmmss.ss or null
	AlertCatogory  string `mapstructure:"alert_catogory,omitempty" json:"alert_catogory,omitempty"`     // alert category, A = Alert category A, B = Alert category B, null
	AlertPriority  string `mapstructure:"alert_priority,omitempty" json:"alert_priority,omitempty"`     // alert priority A=Alarm W=Warning C=Caution, null when SentenceNum=2
	AlertState     string `mapstructure:"alert_state,omitempty" json:"alert_state,omitempty"`           // alert state V=Not Acked, S=Silence, A=Acked, O/U=Resolved,Not Acked, N=normal state, null when SentenceNum=2
	MCode          string `mapstructure:"m_code,omitempty" json:"m_code,omitempty"`                     // manufactrure mnemonic code FEC/null
	AlertID        string `mapstructure:"alert_id,omitempty" json:"alert_id,omitempty"`                 // alert identifier 000-999999
	AlertInstance  string `mapstructure:"alert_instance,omitempty" json:"alert_instance,omitempty"`     // alert instance null
	Revision       string `mapstructure:"revision,omitempty" json:"revision,omitempty"`                 // revision counter
	Escalation     string `mapstructure:"escalation,omitempty" json:"escalation,omitempty"`             // escalation counter
	AlertText      string `mapstructure:"alert_text,omitempty" json:"alert_text,omitempty"`             // alert text max. 16 characters for 1st sentence, maximum length of the field for 2nd sentence later
}

func (s ALF) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := mapstructure.Decode(s, &m)
	if err != nil {
		return m, err
	}
	bm, err := s.BaseSentence.toMap()
	if err != nil {
		return m, err
	}
	for k, v := range bm {
		m[k] = v
	}
	return m, nil
}

// newALF constructor
func newALF(s BaseSentence) (ALF, error) {
	p := NewParser(s)
	p.AssertType(TypeALF)
	m := ALF{
		BaseSentence:   s,
		TotalNum:       p.Int64(0, "TotalNum"),
		SentenceNum:    p.Int64(1, "SentenceNum"),
		SeqID:          p.String(2, "SeqID"),
		LastChangeTime: p.Time(3, "LastChangeTime"),
		AlertCatogory:  p.String(4, "AlertCategory"),
		AlertPriority:  p.String(5, "AlertPriority"),
		AlertState:     p.String(6, "AlertState"),
		MCode:          p.String(7, "MCode"),
		AlertID:        p.String(8, "AlertID"),
		AlertInstance:  p.String(9, "AlertInstance"),
		Revision:       p.String(10, "Revision"),
		Escalation:     p.String(11, "Escalation"),
		AlertText:      p.String(12, "AlertText"),
	}
	return m, p.Err()
}
