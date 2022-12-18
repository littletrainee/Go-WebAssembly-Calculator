js.Global() = object
js.Globla().Get("document") = object
js.Global().Get("document").Call("getElementById", "xxx") = object

```
/*
* s := i[0].String() <= get value type
* s := i[0] 		 <= get value
*/
s := value[0]
```

```
/* like javascript
* document.getElementById.innerHTML = i[0]
* js.Global().Get("document").Call("getElementById", "Result").Set("innerHTML", s)
*/
```

```
go env -w GOOS=js GOARCH=wasm && go env GOOS GOARCH && go build -o ./Asserts/main.wasm ./Source/WebAssembly/main.go && go env -w GOOS=windows GOARCH=amd64 && go env GOOS GOARCH
```