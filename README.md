# gojsend
A JSend Builder and Response Writer for Go.

## JSend
[JSend](https://github.com/omniti-labs/jsend) is a specification that lays down some rules for how JSON responses from web servers should be formatted.

## Install
```shell
$ go get -u github.com/zishone/gojsend
```

## Usage
### JSendBuilder
The JSend Builder provides a convinient way for building JSend responses.

**Success**
```go
builder := gojsend.NewBuilder().
  // JSONEncoder(json.Marshal). // Use to overwrite default json encoder
  Success(map[string]interface{}{"foo": "bar"})
b, _ := builder.Build()
s := string(b)  // {"status":"success","data":{"foo":"bar"}}
m := builder.Response() // map[status:success data:map[foo:bar]]
```

**Fail**
```go
builder := gojsend.NewBuilder().
  // JSONEncoder(json.Marshal). // Use to overwrite default json encoder
  Fail(map[string]interface{}{"foo": "bar"})
b, _ := builder.Build()
s := string(b)  // {"status":"fail","data":{"foo":"bar"}}
m := builder.Response() // map[status:fail data:map[foo:bar]]
```

**Error**
```go
builder := gojsend.NewBuilder().
  Error("foobar").
  Code(1).
  // JSONEncoder(json.Marshal). // Use to overwrite default json encoder
  Data(map[string]interface{}{"foo": "bar"})
b, _ := builder.Build()
s := string(b)  // {"status":"error","message":"foobar","code":1,"data":{"foo":"bar"}}
m := builder.Response() // map[status:error message:foobar code:1 data:map[foo:bar]]
```

### JSendWriter
The JSend Writer extends the http.ResponseWriter with the JSend builder functions and a Send function to pass the built JSend response to http.ResponseWriter's Write function.

**Success**
```go
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
  gojsend.NewWriter(w).
    Success(map[string]interface{}{"foo": "bar"}).
    // JSONEncoder(json.Marshal). // Use to overwrite default json encoder
    // StatusCode(200). // Use to overwrite default status code
    Send()  // Response Body is the same with Builder example with HTTP Status Code 200
}
```

**Fail**
```go
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
  gojsend.NewWriter(w).
    Fail(map[string]interface{}{"foo": "bar"}).
    // JSONEncoder(json.Marshal). // Use to overwrite default json encoder
    // StatusCode(200). // Use to overwrite default status code
    Send()  // Response Body is the same with Builder example with HTTP Status Code 400
}
```

**Error**
```go
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
  gojsend.NewWriter(w).
    Error("foobar").
    Code(1).
    Data(map[string]interface{}{"foo": "bar"}).
    // JSONEncoder(json.Marshal). // Use to overwrite default json encoder
    // StatusCode(200). // Use to overwrite default status code
    Send()  // Response Body is the same with Builder example with HTTP Status Code 500
}
```

### JSendDecorator
Decorates handlers to replace http.ResponseWriter with JSendWriter
```go
func HandlerFunc(w gojsend.JSendWriter, r *http.Request) {
  w.Success(map[string]interface{}{"foo": "bar"}).
    // StatusCode(200).	// Use to overwrite default status code
    Send()  // Response Body is the same with Builder example with HTTP Status Code 200
}

func main() {
  d := gojsend.NewDecorator()// .
    // JSONEncoder(json.Marshal)  // Use to overwrite default json encoder

  http.HandleFunc("/", d.Decorate(HandlerFunc))
  http.ListenAndServe(":3000", nil)
}
```

## Authors
* **Zishran Julbert Garces**

See also the list of [contributors](https://github.com/zishone/gojsend/contributors) who participated in this project.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/zishone/gojsend/blob/master/LICENSE) file for details.