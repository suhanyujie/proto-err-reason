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

通过执行 `eTrans -file=example/t1.proto -o=example/output_error_code.go -pkg=example`

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

至此，再执行 `cd example && go generate` 会生成 code_string.go 文件：

```go
// Code generated by "stringer -type ErrCode -linecomment -output code_string.go"; DO NOT EDIT.

package example

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ERROR_REASON_UNSPECIFIED-0]
	_ = x[USER_NOT_FOUND-1]
	_ = x[Role_Create_Fail-300001001]
}

const (
	_ErrCode_name_0 = "unknownnot found"
	_ErrCode_name_1 = "新增角色失败"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 7, 16}
)

func (i ErrCode) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case i == 300001001:
		return _ErrCode_name_1
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
```

利用 stringer 工具，可以为 output_error_code.go 中的错误定义生成对应的描述实现（`String()` 方法）。

当你使用它时，可以直接用 `fmt.Printf("%s", example.Role_Create_Fail)` 就能打印出错误描述。

综上，你只需在 proto 文件中（如 `example/t1.proto`）进行错误定义，就能直接生成一套 go 代码中可以使用的**错误码**和**错误描述**。

## reference
* https://github.com/yoheimuta/go-protoparser
