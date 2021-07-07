package sim

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkSim_NewUuid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.New()
	}
}

func BenchmarkSim_Run1(b *testing.B) {
	b.ResetTimer()
	sm := NewSim(1)
	for i := 0; i < b.N; i++ {
		sm.Run()
	}
}

func BenchmarkSim_Run2(b *testing.B) {
	b.ResetTimer()
	sm := NewSim(2)
	for i := 0; i < b.N; i++ {
		sm.Run()
	}
}

func BenchmarkSim_Run3(b *testing.B) {
	b.ResetTimer()
	sm := NewSim(3)
	for i := 0; i < b.N; i++ {
		sm.Run()
	}
}

func BenchmarkSim_Run4(b *testing.B) {
	b.ResetTimer()
	sm := NewSim(4)
	for i := 0; i < b.N; i++ {
		sm.Run()
	}
}

func BenchmarkSim_Run5(b *testing.B) {
	b.ResetTimer()
	sm := NewSim(5)
	for i := 0; i < b.N; i++ {
		sm.Run()
	}
}

func BenchmarkSim_Run6(b *testing.B) {
	b.ResetTimer()
	sm := NewSim(6)
	for i := 0; i < b.N; i++ {
		sm.Run()
	}
}
