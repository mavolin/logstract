# logstract

[![Go Report Card](https://goreportcard.com/badge/github.com/mavolin/logstract)](https://goreportcard.com/report/github.com/mavolin/logstract)
[![godoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/mavolin/logstract)
[![GitHub](https://img.shields.io/github/license/mavolin/dismock)](https://github.com/mavolin/logstract/blob/master/LICENSE)

-----

Logstract is an abstract logger for structured logging.
It aims to provide maximum compatibility with as many structured logging frameworks, while introducing as low of an overhead as possible.

## Using logstract in your library

If you want to use logstract in your library, there is nothing you need to do.
Alll configuration is done by the user.

```go
import log "github.com/mavolin/logstract"

func IceCream(free bool, flavor string)  {
    e := log.
        WithFields(logstract.Fields{
            "free": free,
            "flavor": flavor,
        }) // use entries with previously added fields
    
    if free {
        log.Warn("free ice cream for everyone, gonna be expensive 🍦")
    }

    if rand.Float64() <= 0.99 {
        e.
            WithField("mc_donalds", true).
            Fatal("the ice cream is machine broken")
    } else {
        e.Infof("%s ice cream ready for pickup", flavor)
    }
}
```

By default, `logstract.Logger` is a no-op logger.
Therefore, logs will only be created if it is replaced with an actual implementation.

## Using a library that uses logstract

To use enable logging in a library that uses logstract, simply create a `LogFunc` for your logger and store it in the `logstract.Logger` variable.

Below are examples for both [zap](https://github.com/uber-go/zap) and [logrus](https://github.com/sirupsen/logrus/).

### Examples

#### logrus

```go
func Logrus(lvl logstract.Lvl, msg string, fields logstract.Fields) {
    e := logrus.NewEntry(logrus.StandardLogger())
    
    if len(fields) != 0 {
        e.WithFields(logrus.Fields(fields))
    }
    
    switch lvl {
    case logstract.LvlDebug:
        e.Debug(msg)
    case logstract.LvlInfo:
        e.Info(msg)
    case logstract.LvlWarn:
        e.Warn(msg)
    case logstract.LvlError:
        e.Error(msg)
    case logstract.LvlFatal:
        e.Panic(msg)
    }
}
```

#### zap

```go
func Zap() logstract.LogFunc {
    l, _ := zap.NewProduction()
    sugar := l.Sugar()

    return func(lvl logstract.Lvl, msg string, fields logstract.Fields) {
        s := make([]interface{}, 0)
        
        if len(fields) != 0 {
            s = fieldsToSlice(fields)
        }
        
        switch lvl {
        case logstract.LvlDebug:
            l.Debugw(msg, s)
        case logstract.LvlInfo:
            l.Infow(msg, s)
        case logstract.LvlWarn:
            l.Warnw(msg, s)
        case logstract.LvlError:
            l.Errorw(msg, s)
        case logstract.LvlFatal:
            l.Panicw(msg, s)
        }
    }
}

// fieldsToSlice converts the passed Fields to a slice, as used by zap.
func fieldsToSlice(f logstract.Fields) (s []interface{}) {
    s = make([]interface{}, 0, 2*len(f))
    
    for k, v := range f {
        s = append(s, k, v)
    }
    
    return
}
```
