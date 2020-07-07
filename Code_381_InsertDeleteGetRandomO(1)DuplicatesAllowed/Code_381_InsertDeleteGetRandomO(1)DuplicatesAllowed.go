package main


type RandomizedCollection struct {
	list []int
	indexKeyMap map[int]int
	size int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	list := make([]int, 0)
	indexKeyMap := make(map[int]int, 0)
	size := 0
	return RandomizedCollection{list, indexKeyMap, size}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.list = append(this.list, val)
	this.indexKeyMap[this.size] = val
	this.size++
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {

}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {

}
