package tracking_event

import (
	"crypto/sha256"
	"errors"
	"net/url"
)

// Define a Value Object following Domain Driven Design guidelines

var ErrEmptyUid = errors.New("the field Uid can not be empty")

type Uid struct {
	value string
}

func (uid Uid) String() string {
	return uid.value
}

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
	- Hash url + ordered query params + ordered headers
	The purpose is to have an unique identifier per url
 **/
func (url Url) Hash() string {
	h := sha256.New()
	h.Write([]byte(url.value))

	return string(h.Sum(nil))
}

type TrackingEvent struct {
	uid Uid
	url Url
}

func NewUid(value string) (Uid, error) {
	if value == "" {
		return Uid{}, ErrEmptyUid
	}

	return Uid{
		value: value,
	}, nil
}

func NewUrl(value string) (Url, error) {
	_, error := url.Parse(value)
	if error != nil {
		return Url{}, error
	}

	return Url{
		value: value,
	}, nil
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
