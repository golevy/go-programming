## os.Create()
- 创建文件
```go
os.Create("/Users/levylv/desktop/a.txt")
```

## file.WriteString()
- 向文件中写入字符串
```go
file.WriteString("Hello World")
```

## file.Write()
- 向文件中写入数据
- 需要将字符串转换成字节切片 `[]byte()`
```go
file.Write([]byte(str))
```

## file.WriteAt()
- 向文件中指定位置写入数据
```go
file.WriteAt([]byte(str), num)
```

## os.OpenFile()
- OpenFile() 有三个参数
- 第一个参数：`文件路径`
- 第二个参数：模式，常见的模式有
  - `O_RDONLY`(只读模式)
  - `O_WRONLY`(只写模式)
  - `O_RDWR`(可读可写模式)
  - `O_APPEND`(追加模式)
- 第三个参数：`权限`，取值（0-7），通常使用八进制，例如（0666）
  - `0`：没有任何权限
  - `1`：执行权限（如果是可执行文件，是可以运行的）
  - `2`：写权限
  - `3`：写权限与执行权限
  - `4`：读权限
  - `5`：读权限与执行权限
  - `6`：读权限与写权限
  - `7`：读权限、写权限、执行权限
  - `0666`：
    - 第一位代表文件所有者的权限
    - 第二位代表与文件所有者同一用户组的其他用户的权限
    - 第三位代表系统中其他用户的权限。

## file.Read()
- 读取数据