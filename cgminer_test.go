package cgminer

import (
	"testing"
)

func Test_Summary(t *testing.T) {
	miner := New("192.168.1.9", 4028)
	summary, err := miner.Summary()
	if err != nil {
		t.Error(err)
		return
	}
	if summary == nil {
		t.Error("Summary returned nil")
		return
	}
	// TODO: Make some assertions. Need to mock out the data source first?
}
