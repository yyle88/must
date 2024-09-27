package errors_must

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done 期望无错误，假如传入的错误非空时 panic，认为起名为 NoError 比较冗长而且暗示期望有错，这不符合设计预期，设计的预期就是必然无错
func Done(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

// None 期望无错，当传入错误值时 panic
func None(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

func OK(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

func Ok(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

func No(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

func NO(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}
