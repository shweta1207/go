package disjointsets

type DisjointSets struct {
	parent []int // if elements are sparse in values, use maps
	// size []int for better performance
}

func NewDisjointSets(len int) *DisjointSets {
	return &DisjointSets{
		parent: make([]int, len),
	}
}

func (d *DisjointSets) MakeSet(val int) {

}
