package platform

import (
	"testing"
	"time"
)

// Fixed known-answer vector. ref is the HMAC-SHA256 of "<ts>.<body>" keyed by
// secret, computed independently (openssl dgst -sha256 -hmac). Locking the exact
// hex here is what proves this Go port and the TS SDK agree with the platform's
// Util/webhooksig — a transcription error would change the digest and fail here.
const (
	whSecret = "whsec_test"
	whTS     = int64(1700000000)
	whBody   = `{"id":"e1","type":"contest.completed","tenant_id":"t","created_at":"2026-07-13T00:00:00Z","data":{"contest_id":"c1"}}`
	whRef    = "e84001bbe9b6773c8bb0c38729228c8d558e5331c4d7decc9268cee7c45ab753"
)

func whHeader() string { return "t=1700000000,v1=" + whRef }
func whNow() time.Time { return time.Unix(whTS, 0) }

func TestConstructWebhookEvent_Valid(t *testing.T) {
	evt, err := ConstructWebhookEvent(whSecret, whHeader(), []byte(whBody), DefaultWebhookTolerance, whNow())
	if err != nil {
		t.Fatalf("valid delivery rejected: %v", err)
	}
	if evt.Type != "contest.completed" || evt.TenantId != "t" || evt.Id != "e1" {
		t.Fatalf("bad envelope parse: %+v", evt)
	}
	if got := string(evt.Data); got != `{"contest_id":"c1"}` {
		t.Fatalf("data not preserved raw: %s", got)
	}
}

func TestVerifyWebhookSignature_Failures(t *testing.T) {
	tests := []struct {
		name      string
		secret    string
		header    string
		body      []byte
		tolerance time.Duration
		now       time.Time
		want      error
	}{
		{"tampered body", whSecret, whHeader(), []byte(whBody + " "), DefaultWebhookTolerance, whNow(), ErrWebhookSignatureMismatch},
		{"wrong secret", "whsec_other", whHeader(), []byte(whBody), DefaultWebhookTolerance, whNow(), ErrWebhookSignatureMismatch},
		{"empty secret", "", whHeader(), []byte(whBody), DefaultWebhookTolerance, whNow(), ErrWebhookMissingSecret},
		{"stale", whSecret, whHeader(), []byte(whBody), 5 * time.Minute, time.Unix(whTS+301, 0), ErrWebhookTimestampOutOfWindow},
		{"malformed - no v1", whSecret, "t=1700000000", []byte(whBody), DefaultWebhookTolerance, whNow(), ErrWebhookMalformedSignature},
		{"malformed - non-numeric t", whSecret, "t=abc,v1=" + whRef, []byte(whBody), DefaultWebhookTolerance, whNow(), ErrWebhookMalformedSignature},
		{"malformed - empty", whSecret, "", []byte(whBody), DefaultWebhookTolerance, whNow(), ErrWebhookMalformedSignature},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := VerifyWebhookSignature(tc.secret, tc.header, tc.body, tc.tolerance, tc.now); err != tc.want {
				t.Fatalf("got %v, want %v", err, tc.want)
			}
		})
	}
}

func TestVerifyWebhookSignature_OrderIndependentAndToleranceOff(t *testing.T) {
	// Fields reversed + padded — parseSignature is order-independent.
	if err := VerifyWebhookSignature(whSecret, " v1="+whRef+", t=1700000000 ", []byte(whBody), DefaultWebhookTolerance, whNow()); err != nil {
		t.Fatalf("order-independent header rejected: %v", err)
	}
	// tolerance <= 0 disables the timestamp check (far-future now still accepted).
	if err := VerifyWebhookSignature(whSecret, whHeader(), []byte(whBody), 0, time.Unix(whTS+1_000_000, 0)); err != nil {
		t.Fatalf("tolerance-off should skip timestamp check: %v", err)
	}
}

// A bad payload that still carries a valid signature is an invalid-envelope error,
// not a signature error — the two failure classes stay distinct.
func TestConstructWebhookEvent_InvalidPayload(t *testing.T) {
	body := []byte(`not json`)
	// Re-sign the bad body so the signature is valid but the payload isn't an event.
	hdr := "t=1700000000,v1=" + signWebhook(whSecret, whTS, body)
	if _, err := ConstructWebhookEvent(whSecret, hdr, body, DefaultWebhookTolerance, whNow()); err != ErrWebhookInvalidPayload {
		t.Fatalf("got %v, want ErrWebhookInvalidPayload", err)
	}
}
