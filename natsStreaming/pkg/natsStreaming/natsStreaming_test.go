package natsStreaming

import (
	"encoding/json"
	"testing"
)

func TestParseMsg(t *testing.T) {
	s := "Invalid query. Cos this is originally string, not JSON."
	m, _ := json.Marshal(s)

	order := ParseMsg(m)
	if order.OrderUid != "" {
		t.Error("parseMsg return is wrong. ")
	}
}
