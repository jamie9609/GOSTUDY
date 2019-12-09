package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T)  {
	print1()
	time.Sleep(1*time.Millisecond)

}

func TestGoprint1(t *testing.T)  {
	goprint1()
	time.Sleep(1*time.Millisecond)
}

func BenchmarkPrint1(b *testing.B)  {
	for i := 0; i<b.N ;i ++  {
		print1()
	}

}

func BenchmarkGoprint1(b *testing.B)  {
	for i := 0; i<b.N ;i ++  {
		goprint1()
	}
                                                                                                                                                                                                                                                                                                                                                                             
}