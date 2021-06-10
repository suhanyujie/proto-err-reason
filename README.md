# proto-err-reason
trans proto error reason into err code 

将 proto 中的 error reason 转换为 go error code 文件

比如你的 error reason proto 文件定义如下：

```
syntax = "proto3";

package role.v1;

option go_package = "rolesvc/api/v1;v1";

enum ErrorReason {
  ERROR_REASON_UNSPECIFIED = 0; // unknown
  USER_NOT_FOUND = 1;// not found
  Role_Create_Fail = 300001001;// 新增角色失败
}
```

通过执行 errorTrans -file=example/t1.proto -o=example/output_error_code.go -pkg=example

可以将其转换为：

```go
package example

type ErrCode int

const (
	ERROR_REASON_UNSPECIFIED ErrCode = 0 // unknown
	USER_NOT_FOUND ErrCode = 1 // not found
	Role_Create_Fail ErrCode = 300001001 // 新增角色失败

)
```

至此，再通过 `go generate stringer -type ErrCode -linecomment -output code_string.go` 可以得到错误 `String()` 的实现。

## reference
* https://github.com/yoheimuta/go-protoparser