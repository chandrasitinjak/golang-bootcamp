package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFailSum(t *testing.T) {
	result := sum(10, 10)

	//if result != 40 {
	//	//t.Fail()
	//	//t.FailNow()
	//	//t.Error("Result has to be 40")
	//	//t.Fatal("result has to be 40")
	//}

	require.Equal(t, 40, result, "result must be 40")
	
	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := sum(10, 10)

	//if result != 20 {
	//	panic("Result should be 20")
	//}

	assert.Equal(t, 20, result, "Result must 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}
