package nmea

const (
	// TypeROT type for ROT sentences
	TypeROT = "ROT"
)

// ROT rate of turn
// http://aprs.gids.nl/nmea/#hdt
type ROT struct {
	BaseSentence
	Rate   float64 // rate of turn, degrees/minute, "-" bow turns to port
	Status string
}

// newROT constructor
func newROT(s BaseSentence) (ROT, error) {
	p := newParser(s)
	p.AssertType(TypeROT)
	m := ROT{
		BaseSentence: s,
		Rate:         p.Float64(0, "Rate"),
		Status:       p.String(1, "Status"),
	}
	return m, p.Err()
}
