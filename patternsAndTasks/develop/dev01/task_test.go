package main

import (
	"testing"
	"time"
)

func TestParseMsg(t *testing.T) {
	t.Log("Test getNTPTime behavior at different host.")
	{
		testID := 0
		t.Logf("\tTest %d:\twhen host is empty.", testID)
		{

			err := getNTPTime("")
			if err == nil {
				t.Error("Expected connection refused")
			}
		}

		testID++
		t.Logf("\tTest %d:\twhen host is unknown.", testID)
		{
			err := getNTPTime("onfsdonofnw")
			if err == nil {
				t.Error("Expected no such host")
			}
		}

		testID++
		t.Logf("\tTest %d:\twhen host is a random site.", testID)
		{
			err := getNTPTime("wb.ru")
			if err == nil {
				t.Error("Expected timeout")
			}
		}

		testID++
		t.Logf("\tTest %d:\twhen host is apple time server.", testID)
		{
			err := getNTPTime("time.apple.com") //иногда можно поймать таймаут
			if err != nil {
				t.Error("Expect to get current time.")
			}
		}

		testID++
		t.Logf("\tTest %d:\twhen host is default server from beevik ntp.", testID)
		{
			//можно было бы найти значение таймаут и скоректировать ожидаемый результат, а без него придется просто перезапускать тест
			err := getNTPTime("0.beevik-ntp.pool.ntp.org")
			if err != nil {
				t.Errorf("Expect to get current time: %v.", time.Now())
			}
		}

	}
}
