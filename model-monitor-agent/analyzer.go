package main

type Analysis struct {
	TTFTOK       bool
	TPOTOK       bool
	ThroughputOK bool
	OverallOK    bool
	Problem      string
	Action       string
}

func analyzeMetrics(metrics Metrics, thresholds Thresholds) Analysis {
	ttftOK := metrics.TTFT < thresholds.TTFT
	tpotOK := metrics.TPOT < thresholds.TPOT
	throughputOK := metrics.Throughput >= thresholds.Throughput

	analysis := Analysis{
		TTFTOK:       ttftOK,
		TPOTOK:       tpotOK,
		ThroughputOK: throughputOK,
		OverallOK:    ttftOK && tpotOK && throughputOK,
	}

	if !ttftOK {
		analysis.Problem = "TTFT is above the threshold."
		analysis.Action = "Investigate queue time and GPU utilization."
	}

	if !tpotOK {
		analysis.Problem = "TPOT is above the threshold."
		analysis.Action = "Investigate model serving performance."
	}

	if !throughputOK {
		analysis.Problem = "Throughput is below the threshold."
		analysis.Action = "Investigate model load and serving capacity."
	}

	return analysis
}
