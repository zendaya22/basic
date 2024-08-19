package internal

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	name := "dimas"
	must := fmt.Sprintf("nama saya %v", name)
	result := Sayhello(name)

	t.Logf(result)

	if result != must {
		t.Errorf("hasil yang diharapkan : %v\nhasil anda : %v", must, result)
	}
	
}