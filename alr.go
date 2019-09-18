package nmea

const (
	// TypeALR type for ALR sentences
	TypeALR = "ALR"
)

// ALR set alarm state
// http://aprs.gids.nl/nmea/#hdt
type ALR struct {
	BaseSentence `mapstructure:"-,omitempty" json:"base_sentence,omitempty"`
	Time         Time   `mapstructure:"-,omitempty" json:"time,omitempty"`              // time of alarm condition change, UTC
	ID           string `mapstructure:"id,omitempty" json:"id,omitempty"`               // unique alarm number (identifier) at alarm source
	Condition    string `mapstructure:"condition,omitempty" json:"condition,omitempty"` // alarm condition
	ACK          string `mapstructure:"ack,omitempty" json:"ack,omitempty"`             // alarm acknowledge state A=acknowledged, V=unacknowledged
	Text         string `mapstructure:"text,omitempty" json:"text,omitempty"`           // alarm's description text
}

func (s ALR) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"time":       s.Time.String(),
		"time_valid": s.Time.Valid,
		"id":         s.ID,
		"condition":  s.Condition,
		"ack":        s.ACK,
		"text":       s.Text,
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

// newALR constructor
func newALR(s BaseSentence) (ALR, error) {
	p := NewParser(s)
	p.AssertType(TypeALR)
	m := ALR{
		BaseSentence: s,
		Time:         p.Time(0, "Time"),
		ID:           p.String(1, "ID"),
		Condition:    p.String(2, "Condition"),
		ACK:          p.String(3, "ACK"),
		Text:         p.String(4, "Text"),
	}
	return m, p.Err()
}
