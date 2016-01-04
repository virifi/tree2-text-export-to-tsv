package main

import (
	"bytes"
	"testing"
)

func TestCountPrefixTabsOneTab(t *testing.T) {
	actual := countPrefixTabs("\thoge")
	if actual != 1 {
		t.Fatalf("wont 1 got %d", actual)
	}
}

func TestCountPrefixTabsTwoTabs(t *testing.T) {
	actual := countPrefixTabs("\t\thoge")
	if actual != 2 {
		t.Fatalf("wont 2 got %d", actual)
	}
}

func TestCountPrefixTabsOneTabAndMidTab(t *testing.T) {
	actual := countPrefixTabs("\thoge\tfuga")
	if actual != 1 {
		t.Fatalf("wont 1 got %d", actual)
	}
}

func TestCountPrefixTabsNoTab(t *testing.T) {
	actual := countPrefixTabs("hogefuga")
	if actual != 0 {
		t.Fatalf("wont 0 got %d", actual)
	}
}

func TestRemovePrefixTabs(t *testing.T) {
	actual := removePrefixTabs("\t\thoge\tfuga")
	if actual != "hoge\tfuga" {
		t.Fatalf("wont 'hoge\\tfuga' got '%v'", actual)
	}
}

func TestProcessFile(t *testing.T) {
	input := bytes.NewBufferString("hoge\n\tfuga\n\thoge\n")
	output := &bytes.Buffer{}
	err := processFile(input, output)
	if err != nil {
		t.Fatalf("got error : %v", err)
	}
	expected := "hoge\tfuga\n\thoge\n"
	actual := string(output.Bytes())
	if actual != expected {
		t.Fatalf("wont\n%v\ngot\n%v", expected, actual)
	}
}

func TestProcessFile2(t *testing.T) {
	input := bytes.NewBufferString("hoge\n\tfuga\n\t\thoge\n")
	output := &bytes.Buffer{}
	err := processFile(input, output)
	if err != nil {
		t.Fatalf("got error : %v", err)
	}
	expected := "hoge\tfuga\thoge\n"
	actual := string(output.Bytes())
	if actual != expected {
		t.Fatalf("wont\n%v\ngot\n%v", expected, actual)
	}
}

func TestRemovePrefixNumbers(t *testing.T) {
	actual := removePrefixNumbers("1.2.3 hogehoge 3.4.5 fuga 9")
	expected := "hogehoge 3.4.5 fuga 9"
	if actual != expected {
		t.Fatalf("got '%v' wont '%v'", actual, expected)
	}
}

func TestRemovePrefixNumbersNonSpaceBetweenPrefixNumberAndContent(t *testing.T) {
	actual := removePrefixNumbers("1.2.3hogehoge 3.4.5 fuga 9")
	expected := "1.2.3hogehoge 3.4.5 fuga 9"
	if actual != expected {
		t.Fatalf("got '%v' wont '%v'", actual, expected)
	}
}

func TestRemovePrefixNumbersEmptyString(t *testing.T) {
	actual := removePrefixNumbers("")
	expected := ""
	if actual != expected {
		t.Fatalf("got '%v' wont '%v'", actual, expected)
	}
}
