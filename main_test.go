package main

import "testing"

func TestMsisdn(t *testing.T) {
	want := Parsed{
		MnoIdentifier:     "Reliance Jio",
		CountryCode:       91,
		SubscriberNumber:  "083691 10173",
		CountryIdentifier: "IN",
	}

	got, _ := ParseMsisdn("918369110173")
	if want != *got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestLeadingZeroes(t *testing.T) {
	want := Parsed{
		MnoIdentifier:     "Reliance Jio",
		CountryCode:       91,
		SubscriberNumber:  "083691 10173",
		CountryIdentifier: "IN",
	}

	got, _ := ParseMsisdn("00918369110173")
	if want != *got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestLeadingPlus(t *testing.T) {
	want := Parsed{
		MnoIdentifier:     "Reliance Jio",
		CountryCode:       91,
		SubscriberNumber:  "083691 10173",
		CountryIdentifier: "IN",
	}

	got, _ := ParseMsisdn("+918369110173")
	if want != *got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestLeadingPlusAndZeroes(t *testing.T) {
	want := Parsed{
		MnoIdentifier:     "Reliance Jio",
		CountryCode:       91,
		SubscriberNumber:  "083691 10173",
		CountryIdentifier: "IN",
	}

	got, _ := ParseMsisdn("+00918369110173")
	if want != *got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
