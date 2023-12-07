package main

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
)

// 设置哈希数组默认大小16
const DefaultSize = 16

// 设置种子，保证不同哈希函数有不同的计算方式
var seeds = []uint{7, 11, 13, 31, 37, 61}

// 布隆过滤器结构，包括二进制数组和多个哈希函数
type BloomFilter struct {
	set       *bitset.BitSet
	hashFuncs [6]func(seed uint, value string) uint
}

// 构造一个布隆过滤器
func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	bf.set = bitset.New(DefaultSize)
	for i := 0; i < len(bf.hashFuncs); i++ {
		bf.hashFuncs[i] = createHash()
	}
	return bf
}

// 构造6个哈希函数
func createHash() func(seed uint, value string) uint {
	return func(seed uint, value string) uint {
		var result uint = 0
		for i := 0; i < len(value); i++ {
			result = result*seed + uint(value[i])
		}
		return result & (DefaultSize - 1)
	}
}

// 添加元素
func (b *BloomFilter) add(value string) {
	for i, f := range b.hashFuncs {
		b.set.Set(f(seeds[i], value))
	}
}

// 判断元素是否存在
func (b *BloomFilter) contains(value string) bool {
	for i, f := range b.hashFuncs {
		if !b.set.Test(f(seeds[i], value)) {
			return false
		}
	}
	return true
}

func main() {
	filter := NewBloomFilter()
	filter.add("asd")
	fmt.Println(filter.contains("123212"))
	fmt.Println(filter.contains("asd"))
}
