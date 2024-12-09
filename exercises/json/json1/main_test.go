// json1
// Make the tests pass!

package main

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJSONEncoding(t *testing.T) {
	person := Person{Name: "John", Age: 30} // build the person struct to satisfy the expectJSON result

	// Encoding
	data, err := json.Marshal(person)
	if err != nil {
		t.Errorf("JSON encoding failed: %v", err)
	}

	// Check the encoded JSON
	expectedJSON := `{"name":"John","age":30}`
	if string(data) != expectedJSON {
		t.Errorf("Expected JSON: %s, Got: %s", expectedJSON, string(data))
	}
}

func TestDecodeJSON(t *testing.T) {
	jsonData := []byte(`{"name": "John", "age": 30}`)

	var person Person
	err := json.Unmarshal(jsonData, &person) // this function need to be able to modify the person!
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}

	if person.Name != "John" {
		t.Errorf("Expected name to be 'John', got '%s'", person.Name)
	}

	if person.Age != 30 {
		t.Errorf("Expected age to be 30, got %d", person.Age)
	}
}

func TestDecodeJSONWithMissingField(t *testing.T) {
	jsonData := []byte(`{"name": "Alice"}`) // this age field should be missing to pass the test

	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}

	if person.Name != "Alice" {
		t.Errorf("Expected name to be 'Alice', got '%s'", person.Name)
	}

	// Age should be the default value (0)
	if person.Age != 0 {
		t.Errorf("Expected age to be 0, got %d", person.Age)
	}
}
