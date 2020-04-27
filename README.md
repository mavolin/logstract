# logstract

Logstract is an abstract logger for structured logging.
It aims to provide maximum compatibility with as many structured logging frameworks, while introducing as low of an overhead as possible.
Additionally, there are some default implementations so whoever uses your library can get started without a hassle.

## Default Implementations

To get you started more quickly, we have created default implementations.
Currently, the two most-starred structured loggers are supported, namely [zap](https://github.com/uber-go/zap) and [logrus](https://github.com/sirupsen/logrus).
Their implementations can be found in the `impl` package.

## I want to use logstract in my library, what do I need to do?
### Implementing logstract

Using logstract is simple.
Typicallly, you create an exposed variable `Logger *Logger` with `logstract.Default` as value, in your core package.
This ensures that you can always log, without risking a nil dereference panic.

### Logging

If you want to log, simply use your `Logger` variable:
```go
package core

import (

"math/rand"


"github.com/mavolin/logstract/pkg/logstract"
)

var Logger = logstract.Default

func Icecream(free bool, flavor string)  {
    Logger.
        WithFields(logstract.Fields{
            "free": free,
            "flavor": flavor,
        }).
        Debug("yay, ice cream!")
    
    if free {
        Logger.Warn("free icecream for everyone, gonna be expensive 🍦")
    }

    if rand.Float64() <= 0.99 {
        Logger.
            WithField("mc_donalds", true).
            Fatal("the ice cream is machine broken")
    } else {
        Logger.Infof("%s ice cream ready for pickup", flavor)
    }
}
```

If the user doesn't set a `Logger` implementation, no logs will be created and if `Logger` is set, then logs will be generated.

## I want to use a library that uses logstract, what do I need to do?

If the library you are using is following convention, then using your favorite logger can be done in a few steps.

### Using a natively supported logger

This makes it even more easy:

Just set the `Logger` variable of your library to use one of the default implementations:

Example:
```go
func main() {
    l, _ := zap.NewProduction()
    sugar := l.Sugar()

    thatcoollib.Logger = logstract.NewLogger(impl.Zap(sugar))
}
```

### My logger isn't natively supported

Fear not!
Have a look at one of the default implementations (found in pkg/impl), and you'll have a wrapper up and running in no time.

### The logger I want to use does not feature structured logging. What should I do?

The author of the library your're using most likely choose logstract because of it's structured logging capabilities, as structured logging has mostly proven it's dominance over the more conservative just-message logging, at least in the go world.

If you can, change your logger to one, supporting structured logging.
However, if you don't want to switch, in my opinion your best option is to add the fields in parentheses behind:
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

func joinFields(f logstract.Fields, keyValSep, entrySep string) string {
    n := (len(f) - 1) * (len(keyValSep) + len(entrySep))

    for k, v := range f {
        n += len(k) + len(v)
    }

    var b strings.Builder
    b.Grow(n)
    
    first := true

    for k, v := range f {
        if !first {
            b.WriteString(entrySep)
        }

    	b.WriteString(k)
    	b.WriteString(keyValSep)
    	b.WriteString(v)

        first = false
    }
}
```
