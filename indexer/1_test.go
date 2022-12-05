package main

import (
	"os"
	"testing"
)

func TestPushing(t *testing.T) {

	//Change name of the JSON file for testing purposes

	JSONname = "test.json"

	//Remove previous JSON files if existing

	os.Remove(JSONname)

	//Create JSON file

	f, err := os.Create(JSONname)

	check(err)

	f.WriteString("{\"index\": \"testing\",\"records\":[")

	f.Close()

	searchInside("mailtesting")

	f, err = os.OpenFile(JSONname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	check(err)

	ending := []byte("{}]}")

	f.Write(ending)

	f.Close()

	content, err1 := os.ReadFile(JSONname)
	expected, err2 := os.ReadFile("expected.json")

	check(err1)
	check(err2)

	if string(content) != string(expected) {
		t.Fatalf("Expected is different")
	}
}

func BenchmarkPushing(b *testing.B) {

	JSONname = "bench.json"

	os.Remove(JSONname)

	//Create JSON file

	_, err := os.Create(JSONname)

	check(err)

	for i := 0; i < b.N; i++ {
		searchInside("mailtesting")
	}
}
