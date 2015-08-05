package mudoo

import (
    "crypto/rand"
    "io"
    "os"
)

type SessionID string

const (
    // Length of the session ids.
    SessionIDLength = 16

    // Charset from which to build the session ids.
    SessionIDCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// NewSessionID creates a new ~random session that is SessionIDLength long and
// consists of random characters from the SessionIDCharset.
func NewSessionID() (sid SessionID, err os.Error) {
    b := make([]byte, SessionIDLength)

    if _, err = io.ReadFull(rand.Reader, b); err != nil {
        return
    }

    for i := 0; i < SessionIDLength; i++ {
        b[i] = SessionIDCharset[b[i]%uint8(len(SessionIDCharset))]
    }

    sid = SessionID(b)
    return
}
