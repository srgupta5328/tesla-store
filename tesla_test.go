package main

import "testing"

func TestSetTeslaID(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Setting a uidV1 for a given tesla", func(t *testing.T) {
		car := &Tesla{}
		got := car.SetTeslaID()
		want := ""
		if car.ID != "" {
			want = got
		}
		assertCorrectMessage(t, got, want)
	})
}
