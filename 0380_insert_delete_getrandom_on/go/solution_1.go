type RandomizedSet struct {
	set map[int]int
	num []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
		num: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if !ok {
		this.set[val] = len(this.num)
		this.num = append(this.num, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	idx_val, ok := this.set[val]
	if ok {
		var lnum int = len(this.num)
		if idx_val != lnum-1 {
			this.num[idx_val], this.num[lnum-1] = this.num[lnum-1], this.num[idx_val]
			this.set[this.num[idx_val]] = idx_val
		}
		this.num = this.num[:lnum-1]
		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.num[rand.Intn(len(this.num))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */