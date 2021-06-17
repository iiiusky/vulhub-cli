/*
Copyright © 2020 iiusky sky@03sec.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"github.com/iiiusky/slog"
	"go.uber.org/zap"
)

func Logger() *zap.Logger {
	return slog.Logger(&slog.SLoggerSetting{
		AppName:    "vulhub-cli",
		Path:       "./",
		IsDebug:    false,
		CallerSkip: 1,
	})
}

func DebugF(msg string, args ...interface{}) {
	Logger().Debug(fmt.Sprintf(msg, args...))
}

func Debug(msg string, fields ...zap.Field) {
	Logger().Debug(msg, fields...)
}

func InfoF(msg string, args ...interface{}) {
	Logger().Info(fmt.Sprintf(msg, args...))
}

func Info(msg string, fields ...zap.Field) {
	Logger().Info(msg, fields...)
}

func WarnF(msg string, args ...interface{}) {
	Logger().Warn(fmt.Sprintf(msg, args...))
}

func Warn(msg string, fields ...zap.Field) {
	Logger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger().Error(msg, fields...)
}

func DPanicF(msg string, args ...interface{}) {
	Logger().DPanic(fmt.Sprintf(msg, args...))
}

func DPanic(msg string, fields ...zap.Field) {
	Logger().DPanic(msg, fields...)
}

func PanicF(msg string, args ...interface{}) {
	Logger().Panic(fmt.Sprintf(msg, args...))
}

func Panic(msg string, fields ...zap.Field) {
	Logger().Panic(msg, fields...)
}

func FatalF(msg string, args ...interface{}) {
	Logger().Fatal(fmt.Sprintf(msg, args...))
}

func Fatal(msg string, fields ...zap.Field) {
	Logger().Fatal(msg, fields...)
}
