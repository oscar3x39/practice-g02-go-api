# practice-g02-go-api

# GO升級到1.18錯誤解決方式
→ go run main.go 

```
# golang.org/x/sys/unix
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/syscall_darwin.1_13.go:25:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.1_13.go:27:3: //go:linkname must refer to declared function or variable
../../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.1_13.go:40:3: //go:linkname must refer to declared function or variable
```

`go get -u golang.org/x/sys`