package nmea

import (
	"strconv"
)

const (
	// TypeALC type for ALC sentences
	TypeALC = "ALC"
)

type ALCAlertEntry struct {
	MCode         string // manufacture mnemonic code(FEC, null)
	AlertID       string // alert identifier 000 - 999999
	AlertInstance string // alert instance null
	Revision      string // revision counter 1 - 99
}

// ALC cyclic alert list
type ALC struct {
	BaseSentence
	TotalNum    int64 // total number of sentences this message 01 to 99
	SentenceNum int64 // sentence index number 01-99
	Index       int64 // sequential message index
	AlertNum    int64 // Number of alert entries 0 - n
	Alerts      []ALCAlertEntry
}

// newALC constructor
func newALC(s BaseSentence) (ALC, error) {
	p := newParser(s)
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
