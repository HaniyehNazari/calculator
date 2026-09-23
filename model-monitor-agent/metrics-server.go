package main

import (
	"fmt"
	"net/http"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "# HELP model_ttft_seconds Time to first token")
	fmt.Fprintln(w, "# TYPE model_ttft_seconds gauge")
	fmt.Fprintln(w, "model_ttft_seconds 8.5")
	fmt.Fprintln(w, "# HELP model_tpot_milliseconds Time per output token")
	fmt.Fprintln(w, "# TYPE model_tpot_milliseconds gauge")
	fmt.Fprintln(w, "model_tpot_milliseconds 42")

	fmt.Fprintln(w, "# HELP model_throughput_tokens_per_second Model throughput")
	fmt.Fprintln(w, "# TYPE model_throughput_tokens_per_second gauge")
	fmt.Fprintln(w, "model_throughput_tokens_per_second 108000")
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Println("Fake Model Metrics Server")
	fmt.Println("Listening on http://localhost:8081/metrics")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
