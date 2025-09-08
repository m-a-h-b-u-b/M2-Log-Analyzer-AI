//! Module Name: stats.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Rolling z-score anomaly detection for log metrics.

package detector

import "math"

type ZScoreDetector struct {
	window []float64
	size   int
}

func NewZScoreDetector(size int) *ZScoreDetector {
	return &ZScoreDetector{size: size}
}

func (d *ZScoreDetector) Add(value float64) bool {
	d.window = append(d.window, value)
	if len(d.window) > d.size {
		d.window = d.window[1:]
	}

	mean, std := d.meanStd()
	if std == 0 {
		return false
	}
	z := (value - mean) / std
	return math.Abs(z) > 3.0
}

func (d *ZScoreDetector) meanStd() (float64, float64) {
	sum := 0.0
	for _, v := range d.window {
		sum += v
	}
	mean := sum / float64(len(d.window))

	variance := 0.0
	for _, v := range d.window {
		variance += (v - mean) * (v - mean)
	}
	std := math.Sqrt(variance / float64(len(d.window)))
	return mean, std
}
