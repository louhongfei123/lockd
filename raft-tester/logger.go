package main

import (
	"encoding/json"

	"github.com/xiaonanln/lockd/raft"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	demoLogger   *zap.SugaredLogger
	assertLogger assert.TestingT
)

func init() {
	var err error
	cfgJson := []byte(`{
		"level": "debug",
		"outputPaths": ["stderr"],
		"errorOutputPaths": ["stderr"],
		"encoding": "console",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)
	var logConfig zap.Config
	if err = json.Unmarshal(cfgJson, &logConfig); err != nil {
		panic(err)
	}
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	defaultLogger, err := logConfig.Build()
	demoLogger = defaultLogger.Sugar()
	assertLogger = raft.MakeAssertLogger(demoLogger)
}
