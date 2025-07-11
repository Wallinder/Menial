package state

type Session struct {
    SessionId string
    Seq       int64
}

func (s *Session) UpdateSessionId(seq string) {
    s.SessionId = sessionId
}

func (s *Session) UpdateSeq(seq int64) {
    s.seq = seq
}

func (s *Session) GetSessionId() string {
    return s.SessionId
}

func (s *Session) GetSeq() int64 {
    return s.Seq
}