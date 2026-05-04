package main

import (
	"strings"
	"testing"
)
func TestCheckQuotaNotifiesUser(t *testing.T) {
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("Уведомлен (%s) вместо %s", notifiedUser, user)
	}
	
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("неожиданное уведомление <<%s>>, " + "want substring %q", notifiedMsg, wantSubstring)
	}
}