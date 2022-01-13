package uuid_test

import (
	"testing"

	"github.com/phr3nzy/rescounts-api/internals/utils/ids/uuid"
)

func TestGenerateUUID(t *testing.T) {
}

func BenchmarkGenerateUUID(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		id := uuid.New()
		l = int64(len(id))
	}
	b.SetBytes(l)
}

func BenchmarkGenerateUUIDInParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			id := uuid.New()
			l = int64(len(id))
		}
	})
	b.SetBytes(l)
}
