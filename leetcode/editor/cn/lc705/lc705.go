/**
不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

 实现 MyHashSet 类：


 void add(key) 向哈希集合中插入值 key 。
 bool contains(key) 返回哈希集合中是否存在这个值 key 。
 void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。


 示例：


输入：
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove",
 "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
输出：
[null, null, null, true, false, null, true, null, false]

解释：
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // 返回 True
myHashSet.contains(3); // 返回 False ，（未找到）
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // 返回 True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // 返回 False ，（已移除）



 提示：


 0 <= key <= 10⁶
 最多调用 10⁴ 次 add、remove 和 contains


 Related Topics 设计 数组 哈希表 链表 哈希函数 👍 360 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

/*
哈希函数：能够将集合中任意可能的元素映射到一个固定范围的整数值，并将该元素存储到整数值对应的地址上。
冲突处理：由于不同元素可能映射到相同的整数值，因此需要在整数值出现「冲突」时，需要进行冲突处理。总的来说，有以下几种策略解决冲突：
链地址法：为每个哈希值维护一个链表，并将具有相同哈希值的元素都放入这一链表当中。
开放地址法：当发现哈希值 h 处产生冲突时，根据某种策略，从 h 出发找到下一个不冲突的位置。例如，一种最简单的策略是，不断地检查 h+1,h+2,h+3,… 这些整数对应的位置。
再哈希法：当发现哈希冲突后，使用另一个哈希函数产生一个新的地址。
扩容：当哈希表元素过多时，冲突的概率将越来越大，而在哈希表中查询一个元素的效率也会越来越低。因此，需要开辟一块更大的空间，来缓解哈希表中发生的冲突。
*/
package main

const base = 769 // 选择一个合适的素数

type MyHashSet struct {
	data [][]int // 使用切片数组代替 list.List
}

// 构造函数
func Constructor() MyHashSet {
	return MyHashSet{make([][]int, base)}
}

// 哈希函数
func (s *MyHashSet) hash(key int) int {
	return key % base
}

// 添加元素
func (s *MyHashSet) Add(key int) {
	h := s.hash(key)
	if !s.Contains(key) { // 只有不存在时才添加
		s.data[h] = append(s.data[h], key)
	}
}

// 删除元素
func (s *MyHashSet) Remove(key int) {
	h := s.hash(key)
	for i, v := range s.data[h] {
		if v == key {
			// 删除 key（在切片中删除元素）
			leftPart := s.data[h][:i]
			rightPart := s.data[h][i+1:]
			// append 拼接前后两部分，去掉索引 i 处的元素
			s.data[h] = append(leftPart, rightPart...)
			break
		}
	}
}

// 查找元素
func (s *MyHashSet) Contains(key int) bool {
	h := s.hash(key)
	for _, v := range s.data[h] {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// leetcode submit region end(Prohibit modification and deletion)
