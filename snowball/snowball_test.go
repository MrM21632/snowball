package snowball

import "testing"

func TestInitNode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")

	t.Setenv("SNOWBALL_NODE_ID", "32")
	_, err := InitNode(false)
	if err != nil {
		t.Fatalf("Error occurred running InitNode: %s", err)
	}

	t.Setenv("SNOWBALL_NODE_ID", "4000")
	_, err = InitNode(false)
	if err == nil {
		t.Fatalf("No error occurred running InitNode with NODE_ID of 4000")
	}
}

func TestGenerateDuplicateIDs(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode(false)
	var x, y SnowballID
	for i := 0; i < 1000000; i++ {
		y = node.GenerateID()
		if x == y {
			t.Errorf("IDs x(%d) and y(%d) are identical", x, y)
		}
		x = y
	}
}

func BenchmarkGenerateID(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode(false)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = node.GenerateID()
	}
}
