package nmea

const (
	// TypeARC type for ARC sentences
	TypeARC = "ARC"
)

// ARC alarm command refused
// http://aprs.gids.nl/nmea/#ARC
type ARC struct {
	BaseSentence
	Time     Time   // release time of alert command refused
	Reserved string // used for proprietary alerts, defined by the manufactrure(FEC, null)
	ID       string // alarm identifier 000-999999
	Instance string // alarm instance null
	Command  string // refused alart command A=Ackknowlege
}

func (s ARC) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"time":       s.Time.String(),
		"time_valid": s.Time.Valid,
		"reserves":   s.Reserved,
		"id":         s.ID,
		"instance":   s.Instance,
		"command":    s.Command,
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

// newARC constructor
func newARC(s BaseSentence) (ARC, error) {
	p := NewParser(s)
	p.AssertType(TypeARC)
	m := ARC{
		BaseSentence: s,
		Time:         p.Time(0, "Time"),
		Reserved:     p.String(1, "Reserved"),
		ID:           p.String(2, "ID"),
		Instance:     p.String(3, "Instance"),
		Command:      p.String(4, "Command"),
	}
	return m, p.Err()
}
