package detector

import (
	"math"
	"sync"
)

// RollingZScoreDetector detects anomalies based on rolling mean/std
type RollingZScoreDetector struct {
	window    []float64
	size      int
	threshold float64
	mutex     sync.Mutex
}

// NewRollingZScoreDetector creates a new detector
func NewRollingZScoreDetector(windowSize int, threshold float64) *RollingZScoreDetector {
	return &RollingZScoreDetector{
		window:    make([]float64, 0, windowSize),
		size:      windowSize,
		threshold: threshold,
	}
}

// AddValue adds a new value and returns true if anomaly
func (d *RollingZScoreDetector) AddValue(val float64) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if len(d.window) >= d.size {
		d.window = d.window[1:]
	}
	d.window = append(d.window, val)

	mean, std := d.meanStd()
	if std == 0 {
		return false
	}
	z := math.Abs((val - mean) / std)
	return z > d.threshold
}

// meanStd computes mean and standard deviation
func (d *RollingZScoreDetector) meanStd() (float64, float64) {
	n := float64(len(d.window))
	if n == 0 {
		return 0, 0
	}
	sum := 0.0
	for _, v := range d.window {
		sum += v
	}
	mean := sum / n

	variance := 0.0
	for _, v := range d.window {
		variance += (v - mean) * (v - mean)
	}
	std := math.Sqrt(variance / n)
	return mean, std
}
