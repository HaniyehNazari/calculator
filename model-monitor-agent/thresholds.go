package main

type Thresholds struct {
	TTFT       float64
	TPOT       float64
	Throughput float64
}

var defaultThresholds = Thresholds{
	TTFT:       10,
	TPOT:       55,
	Throughput: 100000,
}
