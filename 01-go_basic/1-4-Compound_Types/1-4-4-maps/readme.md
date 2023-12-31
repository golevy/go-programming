## map 概念

- map 是一种 `无序` 的 `键值对` 的集合

- map 最重要的一点是通过 `key` 来快速检索数据，`key` 类似于索引，指向数据的值

- 遍历一个 map 时，键值对的返回顺序是不确定的，且可能在每次迭代中都不同。这是因为 map 在Go语言中是基于哈希表实现的，旨在提供快速的键查找，而不是维护元素的顺序。

## map 创建

```go
var map名字 map[键的类型]值的类型
```

```go
map名字 := map[键的类型]值的类型{}
```

```go
make(map[键的类型]值的类型)
```

## map 初始化

```go
var map名字 = map[键的类型]值的类型{键:值} // 键不允许重复
```

```go
map名字 := map[键的类型]值的类型{键:值} // 键不允许重复
```

```go
make(map[键的类型]值的类型)
map名字[键] = 值 // 键的名字如果一致则完成值的修改
```

## map 键与值

- 通过 key 获取值

```go
map名字[键]
```

- 通过 key 获取值时，判断值是否存在

```go
// 变量1：这是从映射（map）中检索的键对应的值，如果映射中存在该键，变量1 将被赋予该键对应的值
// 变量2：这是一个布尔值，用于指示键是否在映射中存在
// 如果键存在，变量2 将被设置为 true；如果不存在，变量2 将被设置为 false
变量1, 变量2 := map名字[键]
```

- 通过循环方式获取值

```go
for key, value := range m {
	fmt.Println(key)
	fmt.Println(value)
}
```

- 通过 key 删除某个值

```go
delete(map名字, 键)
```