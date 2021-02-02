package main

import (
	"fmt"
	"testing"
)

/*
func TestAbs(_t *testing.T) {
	_t.Errorf("Abs(-1) = %d; want 1", 3)
}
*/

/*
	for i := 0; i < 100000; i++ {
		log.Printf("     sending %v", i)
		rply := API.RplyNtfnStrm {Rtn: 200, EvtId: int32(i)}
		if err := _stream.Send(&rply); err != nil {
			log.Printf("     NtfnStrm got error: %v", err)
			return err
		}

		time.Sleep(1000* time.Millisecond)
	}

func test_agnt() {
	agnt := c_agnt{uuid: "3"}

	fmt.Printf("hi, world %v", agnt.uuid)
	agnt.dojob_loop()
}
*/

func TestVnc(_t *testing.T) {

}

func BenchmarkVnc(_b *testing.B) {
	for i := 0; i < _b.N; i++ {
		fmt.Sprintf("hello")
	}
}
