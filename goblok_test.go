package goblok

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	customProfaneWords := []string{"bad", "words"}
	g := New(customProfaneWords)

	if !reflect.DeepEqual(g.ProfaneWords, customProfaneWords) {
		t.Errorf("Expected profane words: %v, got: %v", customProfaneWords, g.ProfaneWords)
	}

	defaultG := New()
	if !reflect.DeepEqual(defaultG.ProfaneWords, ProfaneWords) {
		t.Errorf("Expected default profane words: %v, got: %v", ProfaneWords, defaultG.ProfaneWords)
	}
}

func TestAdd(t *testing.T) {
	g := New()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty_string", "", ""},
		{"new_word", "badword", "badword"},
		{"existing_word", "badword", ""},
	}

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			actual := g.Add(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	g := New([]string{"badword"})

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty_string", "", ""},
		{"non_existing_word", "word", ""},
		{"existing_word", "badword", "badword"},
		{"remove_after_remove_should_return_empty", "badword", ""},
	}

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			actual := g.Remove(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
func TestFilter(t *testing.T) {
	g := New()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty_string", "", ""},
		{"no_profanity", "Hidupku sangat indah.", "Hidupku sangat indah."},
		{"single_profanity", "Anjing lo", "****** lo"},
		{"single_profanity_with_leet_speak", "4nj1ng lo", "****** lo"},
		{"multiple_profanity", "Anjing lo babi", "****** lo ****"},
		{"multiple_profanity_with_leet_speak", "4nj1ng lo b4b1", "****** lo ****"},
	}

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			actual := g.Filter(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestAnalyze(t *testing.T) {
	g := New()

	cases := []struct {
		name          string
		input         string
		expectedError bool
		expected      Analysis
	}{
		{"empty_string", "", true, Analysis{}},
		{"no_profanity", "Hidupku sangat indah.", false, Analysis{
			Input:     "Hidupku sangat indah.",
			Filtered:  "Hidupku sangat indah.",
			Profane:   []string{},
			Locations: []Location{},
		}},
		{"single_profanity", "Anjing lo", false, Analysis{
			Input:     "Anjing lo",
			Filtered:  "****** lo",
			Profane:   []string{"anjing"},
			Locations: []Location{{Word: "anjing", Index: 0}},
		}},
		{"multiple_profanity", "Anjing lo babi", false, Analysis{
			Input:     "Anjing lo babi",
			Filtered:  "****** lo ****",
			Profane:   []string{"anjing", "babi"},
			Locations: []Location{{Word: "anjing", Index: 0}, {Word: "babi", Index: 10}},
		}},
		{"multiple_profanity_with_leet_speak", "4nj1ng lo b4b1", false, Analysis{
			Input:     "4nj1ng lo b4b1",
			Filtered:  "****** lo ****",
			Profane:   []string{"anjing", "babi"},
			Locations: []Location{{Word: "anjing", Index: 0}, {Word: "babi", Index: 10}},
		}},
	}

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			actual, err := g.Analyze(tc.input)
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
			if !tc.expectedError && !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %+v, got %+v", tc.expected, actual)
			}
		})
	}
}

func TestContains(t *testing.T) {
	g := New()

	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty_string", "", false},
		{"no_profanity", "Hidupku sangat indah.", false},
		{"single_profanity", "Anjing lo", true},
		{"multiple_profanity", "Anjing lo babi", true},
	}

	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			actual := g.Contains(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
