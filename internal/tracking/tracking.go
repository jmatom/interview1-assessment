package tracking

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
)

// Define a Value Object following Domain Driven Design guidelines
var ErrEmptyUid = errors.New("the field Uid can not be empty")
var ErrEmptyUrl = errors.New("the field Url can not be empty")

// Uid Value Object
type Uid struct {
	value string
}

func (uid Uid) String() string {
	return uid.value
}

func NewUid(value string) (Uid, error) {
	if value == "" {
		return Uid{}, ErrEmptyUid
	}

	return Uid{
		value: value,
	}, nil
}

// Url Value Object
type Url struct {
	value string
}

func (url Url) String() string {
	return url.value
}

/**
 * Given an url, generates an unique hash. In real life, this method should do somethiing like:
	- Order query params alphabetically
	- Order headers
		- not necessary according to this assessment, because endpoint to retrieve number of visits expects an url
	- Hash url + ordered query params + ordered headers
	The purpose is to have an unique identifier per url
 **/
func (url Url) Hash() string {
	hash := sha256.Sum256([]byte(url.value))

	return hex.EncodeToString(hash[:])
}

func NewUrl(value string) (Url, error) {
	if value == "" {
		return Url{}, ErrEmptyUrl
	}

	_, error := url.ParseRequestURI(value)
	if error != nil {
		return Url{}, error
	}

	return Url{
		value: value,
	}, nil
}

// TrackingEvent Model
type TrackingEvent struct {
	uid Uid
	url Url
}

func (t *TrackingEvent) Uid() Uid {
	return t.uid
}

func (t *TrackingEvent) Url() Url {
	return t.url
}

func NewTrackingEvent(uid, url string) (TrackingEvent, error) {
	uidVO, err := NewUid(uid)
	if err != nil {
		return TrackingEvent{}, err
	}

	urlVO, err := NewUrl(url)
	if err != nil {
		return TrackingEvent{}, err
	}

	return TrackingEvent{
		uid: uidVO,
		url: urlVO,
	}, nil
}

// VisitCounterRepository defines the expeected behaviour frorm a counter storage
type CounterRepository interface {
	AddVisit(trackingEvent TrackingEvent) error
	GetVisits(url Url) uint64
}
