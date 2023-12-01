package filesplitter

import (
	"os"
	"testing"
)

func TestSplitJoinFile(t *testing.T) {
	f, err := os.Open("test.pdf")
	if err != nil {
		t.Errorf("%v", err)
	}
	err = SplitFile(f, "split.pdf", "./", 1<<20)
	if err != nil {
		t.Errorf("%v", err)
	}
	f.Close()
	f, err = os.Create("join_test.pdf")
	if err != nil {
		t.Errorf("%v", err)
	}
	err = JoinFile("split.pdf.part0", f)
	if err != nil {
		t.Errorf("%v", err)
	}
}
