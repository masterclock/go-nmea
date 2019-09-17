package nmea

const (
	// TypeALR type for ALR sentences
	TypeALR = "ALR"
)

// ALR set alarm state
// http://aprs.gids.nl/nmea/#hdt
type ALR struct {
	BaseSentence
	Time      Time   // time of alarm condition change, UTC
	ID        string // unique alarm number (identifier) at alarm source
	Condition string // alarm condition
	ACK       string // alarm acknowledge state A=acknowledged, V=unacknowledged
	Text      string // alarm's description text
}

// newALR constructor
func newALR(s BaseSentence) (ALR, error) {
	p := newParser(s)
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
