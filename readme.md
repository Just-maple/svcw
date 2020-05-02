# go-service-wrapper
> helps you to make interruption in every interface service call


## Installation

Install `svcw` by running:

```sh
go get github.com/Just-maple/svcw
```


## Usage example

If your application use `wire` or other `DI`、`IOC` design

the service call relation would seems like a tree

- `A.DoSomething()` may contain `A.B.DoSomething()`
- `B.DoSomething()` may contain `B.C.DoSomething()` and `B.E.DoSomething()`
- `C.DoSomething()` may contain `C.D.DoSomething()`

```go
func (this A) DoSomething(ctx context.Context){ 
    this.B.DoSomething(ctx)
}

func (this B) DoSomething(ctx context.Context){ 
    this.C.DoSomething(ctx)
    this.E.DoSomething(ctx)
}

func (this C) DoSomething(ctx context.Context){ 
    this.D.DoSomething(ctx)
}

...
```

in micro service design,we could use opentracing to set interruption in every service call

but in single application,if we want to know how the service call

we should write some repeat dull codes,like:

```go
func (this A) DoSomething(ctx context.Context){ 
    t := time.Now()
    logger.FromContext(ctx).Printf("call:A:start")
    defer func(){logger.FromContext(ctx).Printf("call:A:finished,time used:%v",time.Now().Since(t))}()
    this.B.DoSomething(ctx)
}

func (this B) DoSomething(ctx context.Context){ 
    t := time.Now()
    logger.FromContext(ctx).Printf("call:B:start")
    defer func(){logger.FromContext(ctx).Printf("call:B:finished,time used:%v",time.Now().Since(t))}()
    this.C.DoSomething(ctx)
}

func (this C) DoSomething(ctx context.Context){ 
    t := time.Now()
    logger.FromContext(ctx).Printf("call:C:start")
    defer func(){logger.FromContext(ctx).Printf("call:C:finished,time used:%v",time.Now().Since(t))}()
    this.D.DoSomething(ctx)
}

...
```


when your service and call comes complex,it's hard and dull to manage these codes

svcw helps you to gen these repeat and dull codes,but not in your service codes

with running in runtime

```go
err := svcw.Gen(&yourApplication, interruptionCode, true)
if err != nil {
	panic(err)
}
```
when `svcw` get your application ptr, it will analyse your service interface call relation

and it gens files `zitf.go` contain all structs that implement your service interface

look at the [example](./example/cxk/zitf.go)
 
you can design what should be done in every service call

look at the [example](./example/example.go) `interruptionCode`

and when your application start 

use the build tag `go build -tags svcw` to replace the service interface in runtime

these gen structs registered in a container that if you call: 

```go
wrapflask.Wrap(&yourApplication)
```

all the dependencies will be replaced by the wrapped service 

so your service call will do all you want in interruption