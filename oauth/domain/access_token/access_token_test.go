package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if ExpirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("brand new access token should not have defined access token")
	}

	if at.UserID != 0 {
		t.Error("brand new access token should not have defined user ID")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("empty access token should be expired")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	if at.IsExpired() {
		t.Error("access token expiring three hours from now should not be expired")
	}

}
