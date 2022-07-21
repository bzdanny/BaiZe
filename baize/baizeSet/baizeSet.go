package baizeSet

import "fmt"

// Set 利用泛型，定义一个泛型类型(set), 泛型的写法使用中括号，comparable是新增的一个内置类型
type Set[T comparable] map[T]struct{}

// Add 向set中添加元素
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// AddAll 向set中添加多个元素
func (s Set[T]) AddAll(vs ...T) {
	for i := 0; i < len(vs); i++ {
		s[vs[i]] = struct{}{}
	}
}

// Remove 移除set中某个元素
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// RemoveAll 移除set中多个元素
func (s Set[T]) RemoveAll(vs ...T) {
	for i := 0; i < len(vs); i++ {
		delete(s, vs[i])
	}
}

// Len 获取set的长度
func (s Set[T]) Len() int {
	return len(s)
}

// Contains 查询某个元素v是否在set中
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// 实现String方法，优化set的打印效果
func (s Set[T]) String() string {
	return fmt.Sprint(s.getAll())
}

// Clear 清空set
func (s Set[T]) Clear() Set[T] {
	return make(Set[T])
}

// ToSlice set转slice
func (s Set[T]) ToSlice() []T {
	return s.getAll()
}

func (s Set[T]) getAll() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
