package nmea

const (
	// TypeALF type for ALF sentences
	TypeALF = "ALF"
)

// ALF alert sentence
// http://aprs.gids.nl/nmea/#hdt
type ALF struct {
	BaseSentence
	TotalNum       int64  // total number of alf senteces this message (1 - 2)
	SentenceNum    int64  // sentence number (1 - 2)
	SeqID          string // sequential message identifier (0 - 9)
	LastChangeTime Time   // time of last change, hhmmss.ss or null
	AlertCatogory  string // alert category, A = Alert category A, B = Alert category B, null
	AlertPriority  string // alert priority A=Alarm W=Warning C=Caution, null when SentenceNum=2
	AlertState     string // alert state V=Not Acked, S=Silence, A=Acked, O/U=Resolved,Not Acked, N=normal state, null when SentenceNum=2
	MCode          string // manufactrure mnemonic code FEC/null
	AlertID        string // alert identifier 000-999999
	AlertInstance  string // alert instance null
	Revision       string // revision counter
	Escalation     string // escalation counter
	AlertText      string // alert text max. 16 characters for 1st sentence, maximum length of the field for 2nd sentence later
}

// newALF constructor
func newALF(s BaseSentence) (ALF, error) {
	p := newParser(s)
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
