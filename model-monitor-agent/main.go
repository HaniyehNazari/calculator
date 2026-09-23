package main

import "fmt"

func main() {
	model := "26B"

	ttft, err := queryPrometheus("model_ttft_seconds")
	if err != nil {
		fmt.Println("Error reading TTFT:", err)
		return
	}

	tpot, err := queryPrometheus("model_tpot_milliseconds")
	if err != nil {
		fmt.Println("Error reading TPOT:", err)
		return
	}

	throughput, err := queryPrometheus("model_throughput_tokens_per_second")
	if err != nil {
		fmt.Println("Error reading throughput:", err)
		return
	}

	metrics := Metrics{
		TTFT:       ttft,
		TPOT:       tpot,
		Throughput: throughput,
	}

	analysis := analyzeMetrics(metrics, defaultThresholds)

	fmt.Println("AI Model Monitor Agent")
	fmt.Println("Checking model:", model)

	fmt.Println()
	fmt.Println("Metrics:")
	fmt.Println("TTFT:", metrics.TTFT, "seconds")
	fmt.Println("TPOT:", metrics.TPOT, "ms")
	fmt.Println("Throughput:", metrics.Throughput, "tokens/s")

	fmt.Println()
	fmt.Println("Status:")

	if analysis.TTFTOK {
		fmt.Println("TTFT: OK")
	} else {
		fmt.Println("TTFT: WARNING")
	}

	if analysis.TPOTOK {
		fmt.Println("TPOT: OK")
	} else {
		fmt.Println("TPOT: WARNING")
	}

	if analysis.ThroughputOK {
		fmt.Println("Throughput: OK")
	} else {
		fmt.Println("Throughput: WARNING")
	}

	fmt.Println()

	if analysis.OverallOK {
		fmt.Println("Overall Status: OK")
	} else {
		fmt.Println("Overall Status: WARNING")
		fmt.Println()
		fmt.Println("Problem:")
		fmt.Println("-", analysis.Problem)
		fmt.Println("Action:", analysis.Action)
	}
}
