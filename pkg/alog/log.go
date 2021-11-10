// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package alog

import (
	"github.com/ethersphere/bee/pkg/logging"
)

var ALog logging.Logger

func SetLogger(newLogger logging.Logger) {
	ALog = newLogger
}
func Debug(args ...interface{}) {
	ALog.Debug(args)
}

func Info(args ...interface{}) {
	ALog.Info(args)
}

func Error(args ...interface{}) {
	ALog.Error(args)
}
func DebugF(format string, args ...interface{}) {
	ALog.Debugf(format, args)
}

func InfoF(format string, args ...interface{}) {
	ALog.Infof(format, args)
}

func errorF(format string, args ...interface{}) {
	ALog.Errorf(format, args)
}
