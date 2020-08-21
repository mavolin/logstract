# logstract

[![Go Report Card](https://goreportcard.com/badge/github.com/mavolin/logstract)](https://goreportcard.com/report/github.com/mavolin/logstract)
[![godoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/mavolin/logstract)
[![GitHub](https://img.shields.io/github/license/mavolin/dismock)](https://github.com/mavolin/logstract/blob/master/LICENSE)

-----

Logstract is an abstract logger for structured logging.
It aims to provide maximum compatibility with as many structured logging frameworks, while introducing as low of an overhead as possible.

## I want to use logstract in my library, what do I need to do?

Nothing. 
Logstract was designed to be as easy to use as possible, this means you have to create no global variable, that might cause you headaches about import cycles.
Instead, you just call logstract's exposed functions:

```go
func Icecream(free bool, flavor string)  {
    logstract.
        WithFields(logstract.Fields{
            "free": free,
            "flavor": flavor,
        }).
        Debug("yay, ice cream!")
    
    if free {
        logstract.Warn("free icecream for everyone, gonna be expensive 🍦")
    }

    if rand.Float64() <= 0.99 {
        logstract.
            WithField("mc_donalds", true).
            Fatal("the ice cream is machine broken")
    } else {
        logstract.Infof("%s ice cream ready for pickup", flavor)
    }
}
```

By default, `logstract.Logger` is a no-op logger.
Therefore, logs will only be creates if it is replaced with an actual implementation.

## I want to use a library that uses logstract, what do I need to do?

To use logstract, simply create a `LogFunc` for your logger and store it in the `logstract.Logger` variable.

Below are examples for both [zap](https://github.com/uber-go/zap) and [logrus](https://github.com/sirupsen/logrus/), however, here are no default implementations to keep logstract dependency free.

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


### The logger I want to use does not feature structured logging. What should I do?

The author of the library your're using most likely choose logstract because of it's structured logging capabilities, as structured logging has mostly proven it's dominance over the more conservative just-message logging, at least in the go world.

If you can, change your logger to structured logger.
However, if you don't want to switch, in my opinion your best option is to add the fields in a parentheses behind:
```go
func MyLogFunc(lvl logstract.Lvl, msg string, fields logstract.Fields) {
    if fields != nil {
        msg += " (" + joinFields(fields, ": ", ", ") + ")"
    }
    
    switch lvl {
    case logstract.LvlDebug:
        myLogger.Debug(msg)
    case logstract.LvlInfo:
        myLogger.Info(msg)
    ...
    }
}

func joinFields(f logstract.Fields, keyValSep, entrySep string) (s string) {    
    first := true
    
    for k, v := range f {
        if !first {
            s += entrySep
        }
        
        s += k + keyValSep + fmt.Sprimtf("%+v", v")
        
        first = false
    }
    
    return
}
```
