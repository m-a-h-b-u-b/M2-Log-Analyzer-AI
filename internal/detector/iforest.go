//! Module Name: iforest.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Isolation Forest anomaly detector with ONNX runtime.

package detector

import (
	"github.com/microsoft/onnxruntime-go"
)

type IForestDetector struct {
	session *onnxruntime.Session
}

func NewIForest(modelPath string) (*IForestDetector, error) {
	sess, err := onnxruntime.NewSession(modelPath)
	if err != nil {
		return nil, err
	}
	return &IForestDetector{session: sess}, nil
}

func (d *IForestDetector) Score(features []float32) (float32, error) {
	// TODO: implement ONNX inference
	return 0.0, nil
}
