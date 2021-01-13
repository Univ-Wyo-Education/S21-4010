package index

import "testing"

func Test_ReadIndex(t *testing.T) {
	fn := "testdata/index.json"
	idx, err := ReadIndex(fn)
	if err != nil {
		t.Fatalf("Failed to read index [%s] - fatal. error [%s]", fn, err)
	}
	_ = idx
}
