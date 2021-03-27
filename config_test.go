package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	os.Setenv("PROJECT_NAME", "test-proj-name")
	os.Setenv("INSTANCE_NAME", "test-instance-name")
	os.Setenv("SQL_NAME", "test-sql-name")
}

func BenchmarkViperDoubleDip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ViperDoubleDip()
	}
}

func BenchmarkOSDoubleDip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OSSingleDip()
	}
}

func BenchmarkViperSingleDip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ViperSingleDip()
	}
}

func BenchmarkOSSingleDip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OSSingleDip()
	}
}

func BenchmarkViperGetBool(b *testing.B) {
	defer setUnsetEnv("TEST_BOOL", "true")()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ViperGetBool()
	}
}

func BenchmarkOSGetBool(b *testing.B) {
	defer setUnsetEnv("TEST_BOOL", "true")()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		OSGetBool()
	}
}

func BenchmarkViperGetBoolNotSet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ViperGetBool()
	}
}

func BenchmarkOSGetBoolNotSet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		OSGetBool()
	}
}

func BenchmarkViperGetBoolWithError(b *testing.B) {
	defer setUnsetEnv("TEST_BOOL", "true")()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ViperGetBoolWithError()
	}
}

func BenchmarkOSGetBoolWithError(b *testing.B) {
	defer setUnsetEnv("TEST_BOOL", "true")()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		OSGetBoolWithError()
	}
}

func BenchmarkViperGetBoolWithErrorNotSet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ViperGetBoolWithError()
	}
}

func BenchmarkOSGetBoolWithErrorNotSet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		OSGetBoolWithError()
	}
}

func BenchmarkOSGetBoolWithErrorNotBool(b *testing.B) {
	defer setUnsetEnv("TEST_BOOL", "abc")()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		OSGetBoolWithError()
	}
}

func setUnsetEnv(k, v string) func() {
	os.Setenv(k, v)

	return func() {
		os.Unsetenv(k)
	}
}
