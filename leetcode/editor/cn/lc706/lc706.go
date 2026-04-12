/**
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。

 实现 MyHashMap 类：


 MyHashMap() 用空映射初始化对象
 void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更
新其对应的值 value 。
 int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
 void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。




 示例：


输入：
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
输出：
[null, null, null, 1, -1, null, 1, null, -1]

解释：
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]




 提示：


 0 <= key, value <= 10⁶
 最多调用 10⁴ 次 put、get 和 remove 方法


 Related Topics 设计 数组 哈希表 链表 哈希函数 👍 449 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)

type elem struct {
	key   int
	value int
}

type MyHashMap struct {
	data [][]elem
	base int
}

func Constructor() MyHashMap {
	const base = 857
	return MyHashMap{
		data: make([][]elem, base),
		base: base,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for i := range this.data[h] {
		item := this.data[h][i]
		if item.key == key {
			this.data[h][i].value = value
			return
		}
	}
	this.data[h] = append(this.data[h], elem{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for _, e := range this.data[h] {
		if e.key == key {
			return e.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for i := range this.data[h] {
		item := this.data[h][i]
		if item.key == key {
			left := this.data[h][:i]
			right := this.data[h][i+1:]
			this.data[h] = append(left, right...)
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// leetcode submit region end(Prohibit modification and deletion)
