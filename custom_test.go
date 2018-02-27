package fortnox

import (
	"encoding/json"
	"testing"
)

func TestFloatish_UnmarshalJSON(t *testing.T) {

	testS := struct {
		FloatFieldFromStr   Floatish
		FloatFieldFromFloat Floatish
	}{}

	testPayload := `{"FloatFieldFromStr": "8.8888", "FloatFieldFromFloat": 9.9999}`

	if err := json.Unmarshal([]byte(testPayload), &testS); err != nil {
		t.Fatal(err)
	}

	if testS.FloatFieldFromFloat != 9.9999 {
		t.Fatalf("unexpected value %.04f", testS.FloatFieldFromFloat)
	}
	if testS.FloatFieldFromStr != 8.8888 {
		t.Fatalf("unexpected value %.04f", testS.FloatFieldFromStr)
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {

	testS := struct {
		FnoxDate Date
	}{}

	testPayload := `{"FnoxDate": "1988-03-18"}`

	if err := json.Unmarshal([]byte(testPayload), &testS); err != nil {
		t.Fatal(err)
	}

	if testS.FnoxDate.Year != 1988 {
		t.Fatalf("unexpected value %d", testS.FnoxDate.Year)
	}
	if testS.FnoxDate.Month != 3 {
		t.Fatalf("unexpected value %d", testS.FnoxDate.Month)
	}
	if testS.FnoxDate.Date != 18 {
		t.Fatalf("unexpected value %d", testS.FnoxDate.Date)
	}

}

func TestDate_String(t *testing.T) {

	testD := Date{2017, 05, 21}

	testStr := testD.String()
	if testStr != "2017-05-21" {
		t.Fatal("unexpected format", testStr)
	}
}

func TestIntish_Int(t *testing.T) {
	intish := Intish(99)

	i := intish.Int()

	if i != 99 {
		t.Fatal("unexpected value", i)
	}
}

func TestFloatish_Float64(t *testing.T) {
	ish := Floatish(99.99)

	f := ish.Float64()

	if f != 99.99 {
		t.Fatal("unexpected value", f)
	}

}

func TestStringIsh_UnmarshalJSON(t *testing.T) {

	testS := struct {
		SomeString StringIsh
	}{}

	testPayload := `{"SomeString": 1988}`

	if err := json.Unmarshal([]byte(testPayload), &testS); err != nil {
		t.Fatal(err)
	}

	if testS.SomeString != "1988" {
		t.Fatalf("unexpected value %d", testS.SomeString)
	}

	testPayload = `{"SomeString": "1989"}`

	if err := json.Unmarshal([]byte(testPayload), &testS); err != nil {
		t.Fatal(err)
	}

	if testS.SomeString != "1989" {
		t.Fatalf("unexpected value %d", testS.SomeString)
	}
}
