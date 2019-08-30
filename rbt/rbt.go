/*

rbt.go

An implementation of a Red-Black Tree.
For each RBT there is two interface that should be implemented:
  - A comparator interface to compare two nodes in RBT.
  - A dumper interface to dump RBT for debugging.

MIT License

Copyright (c) 2018 rezamirz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package rbt

const (
	Black = false
	Red   = true
)

type Comparator interface {
	/*
	 * Compares two keys and returns -1, 0, or 1.
	 * The return value is used to sort the keys.
	 */
	Compare(k1, k2 interface{}) int
}

type Dumper interface {
	/*
	 * Dumps key and value, useful for debugging
	 */
	Dump(k, v interface{})
}

type RBTNode struct {
	key   interface{} /* Key */
	value interface{} /* Associated value */
	left  *RBTNode
	right *RBTNode
	n     uint32 /* # of nodes in subtree */
	color bool
}

type RBT struct {
	root *RBTNode

	comparator Comparator
	dumper     Dumper
}

func isRed(node *RBTNode) bool {
	if node == nil {
		return false
	}

	return node.color == Red
}

func NewRBT(comparator Comparator, dumper Dumper) *RBT {
	rbt := &RBT{
		comparator: comparator,
		dumper:     dumper,
	}

	return rbt
}

func rbtSize(node *RBTNode) uint32 {
	if node == nil {
		return 0
	}

	return node.n
}

func (rbt *RBT) Size() uint32 {
	return rbtSize(rbt.root)
}

func (rbt *RBT) IsEmpty() bool {
	return rbt.root == nil
}

func rotateLeft(h *RBTNode) *RBTNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = Red
	x.n = h.n
	h.n = 1 + rbtSize(h.left) + rbtSize(h.right)
	return x
}

func rotateRight(h *RBTNode) *RBTNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = Red
	x.n = h.n
	h.n = 1 + rbtSize(h.left) + rbtSize(h.right)
	return x
}

func flipColors(h *RBTNode) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func rbtGet(node *RBTNode, key interface{}, comparator Comparator) interface{} {

	if node == nil {
		return nil
	}

	cmp := comparator.Compare(key, node.key)
	if cmp < 0 {
		return rbtGet(node.left, key, comparator)
	} else if cmp > 0 {
		return rbtGet(node.right, key, comparator)
	}
	return node.value
}

func (rbt *RBT) Get(key interface{}) interface{} {
	if rbt == nil {
		return nil
	}

	return rbtGet(rbt.root, key, rbt.comparator)
}

func (rbt *RBT) Contains(key interface{}) bool {
	return rbt.Get(key) != nil
}

func rbtPut(node *RBTNode, key interface{}, value interface{}, comparator Comparator) *RBTNode {
	if node == nil {
		node = &RBTNode{
			key:   key,
			value: value,
			n:     1,
			color: Red,
		}
		return node
	}

	cmp := comparator.Compare(key, node.key)
	if cmp < 0 {
		node.left = rbtPut(node.left, key, value, comparator)
	} else if cmp > 0 {
		node.right = rbtPut(node.right, key, value, comparator)
	} else {
		node.value = value
	}

	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}

	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}

	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}

	node.n = rbtSize(node.left) + rbtSize(node.right) + 1
	return node
}

func (rbt *RBT) Put(key interface{}, value interface{}) {
	rbt.root = rbtPut(rbt.root, key, value, rbt.comparator)
	rbt.root.color = Black
}

/* Returns min key */
func rbtMin(node *RBTNode) *RBTNode {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	}

	return rbtMin(node.left)
}

func (rbt *RBT) Min() interface{} {
	node := rbtMin(rbt.root)
	if node != nil {
		return node.key
	}

	return nil
}

func rbtMax(node *RBTNode) interface{} {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node.key
	}

	return rbtMax(node.right)
}

func (rbt *RBT) Max() interface{} {
	return rbtMax(rbt.root)
}

func rbtFloor(node *RBTNode, key interface{}, comparator Comparator) *RBTNode {
	if node == nil {
		return nil
	}

	cmp := comparator.Compare(key, node.key)
	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		t := rbtFloor(node.left, key, comparator)
		return t
	}

	t := rbtFloor(node.right, key, comparator)
	if t != nil {
		return t
	}

	return node
}

/* Returns the largest key in the symbol table less than or equal to key. */
func (rbt *RBT) Floor(key interface{}) interface{} {
	t := rbtFloor(rbt.root, key, rbt.comparator)
	if t != nil {
		return t.key
	}
	return nil
}

func rbtCeiling(node *RBTNode, key interface{}, comparator Comparator) *RBTNode {

	if node == nil {
		return nil
	}

	cmp := comparator.Compare(key, node.key)
	if cmp == 0 {
		return node
	}

	if cmp > 0 {
		t := rbtCeiling(node.right, key, comparator)
		return t
	}

	t := rbtCeiling(node.left, key, comparator)
	if t != nil {
		return t
	}

	return node
}

/* Returns the smallest key in the symbol table greater than or equal to key. */
func (rbt *RBT) Ceiling(key interface{}) interface{} {
	t := rbtCeiling(rbt.root, key, rbt.comparator)
	if t != nil {
		return t.key
	}
	return nil
}

func rbtSelect(node *RBTNode, rank uint32) *RBTNode {
	if node == nil {
		return nil
	}

	t := rbtSize(node.left)
	if t > rank {
		return rbtSelect(node.left, rank)
	} else if t < rank {
		return rbtSelect(node.right, rank-t-1)
	}
	return node
}

func (rbt *RBT) Select(rank uint32) interface{} {
	node := rbtSelect(rbt.root, rank)
	return node.key
}

func rbtRank(key interface{}, node *RBTNode, comparator Comparator) uint32 {

	if node == nil {
		return 0
	}

	cmp := comparator.Compare(key, node.key)
	if cmp < 0 {
		return rbtRank(key, node.left, comparator)
	} else if cmp > 0 {
		return 1 + rbtSize(node.left) + rbtRank(key, node.right, comparator)
	}
	return rbtSize(node.left)
}

func (rbt *RBT) Rank(key interface{}) uint32 {
	return rbtRank(key, rbt.root, rbt.comparator)
}

/*
 * Assuming that h is red and both h->left and h->left->left
 * are black, make h->left or one of its children red.
 */
func moveRedLeft(h *RBTNode) *RBTNode {
	flipColors(h)

	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}

	return h
}

/*
 * Assuming that h is red and both h->right and h->right->left
 * are black, make h->right or one of its children red.
 */
func moveRedRight(h *RBTNode) *RBTNode {
	flipColors(h)

	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}

	return h
}

/* restore red-black tree invariant */
func balance(h *RBTNode) *RBTNode {
	if isRed(h.right) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}

	h.n = rbtSize(h.left) + rbtSize(h.right) + 1
	return h
}

func rbtDelete_min(node *RBTNode) (*RBTNode, interface{}, interface{}) {
	if node.left == nil {
		h := node.right
		return h, node.key, node.value
	}

	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}

	var key, value interface{}
	node.left, key, value = rbtDelete_min(node.left)
	return balance(node), key, value
}

func (rbt *RBT) DeleteMin() (bool, interface{}, interface{}) {
	if rbt.IsEmpty() {
		return false, nil, nil
	}

	root := rbt.root
	if !isRed(root.left) && !isRed(root.right) {
		root.color = Red
	}

	root, key, value := rbtDelete_min(root)
	if root != nil {
		root.color = Black
	}
	rbt.root = root
	return true, key, value
}

func rbtDelete_max(node *RBTNode) (*RBTNode, interface{}, interface{}) {
	if isRed(node.left) {
		node = rotateRight(node)
	}

	if node.right == nil {
		h := node.left
		return h, node.key, node.value
	}

	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}

	var key, value interface{}
	node.right, key, value = rbtDelete_max(node.right)
	return balance(node), key, value
}

func (rbt *RBT) DeleteMax() (bool, interface{}, interface{}) {
	if rbt.IsEmpty() {
		return false, nil, nil
	}

	root := rbt.root
	/* if both children of root are black, set root to red */
	if !isRed(root.left) && !isRed(root.right) {
		root.color = Red
	}

	var key, value interface{}
	root, key, value = rbtDelete_max(root)
	if root != nil {
		root.color = Black
	}
	rbt.root = root
	return true, key, value
}

func rbtDelete(node *RBTNode, key interface{}, comparator Comparator) *RBTNode {

	if node == nil {
		return nil
	}

	cmp := comparator.Compare(key, node.key)
	if cmp < 0 {
		if !isRed(node.left) && !isRed(node.left.left) {
			node = moveRedLeft(node)
		}
		node.left = rbtDelete(node.left, key, comparator)
	} else {
		if isRed(node.left) {
			node = rotateRight(node)
		}
		if comparator.Compare(key, node.key) == 0 && node.right == nil {
			return nil
		}
		if !isRed(node.right) && !isRed(node.right.left) {
			node = moveRedRight(node)
		}
		if comparator.Compare(key, node.key) == 0 {
			node.right, node.key, node.value = rbtDelete_min(node.right)
		} else {
			node.right = rbtDelete(node.right, key, comparator)
		}
	}

	return balance(node)
}

func (rbt *RBT) Delete(key interface{}) bool {
	if key == nil {
		return false
	}

	if !rbt.Contains(key) {
		return false
	}

	if !isRed(rbt.root.left) && !isRed(rbt.root.right) {
		rbt.root.color = Red
	}

	rbt.root = rbtDelete(rbt.root, key, rbt.comparator)
	if rbt.root != nil {
		rbt.root.color = Black
	}
	return true
}

func rbtHeight(node *RBTNode) int {
	if node == nil {
		return -1
	}
	lh := rbtHeight(node.left)
	rh := rbtHeight(node.right)
	if rh > lh {
		return 1 + rh
	}
	return 1 + lh
}

func (rbt *RBT) Height() int {
	return rbtHeight(rbt.root)
}

func isBST(node *RBTNode, minkey interface{}, maxkey interface{}, comparator Comparator) bool {
	if node == nil {
		return true
	}

	if minkey != nil && comparator.Compare(node.key, minkey) <= 0 {
		return false
	}

	if maxkey != nil && comparator.Compare(node.key, maxkey) >= 0 {
		return false
	}
	return isBST(node.left, minkey, node.key, comparator) && isBST(node.right, node.key, maxkey, comparator)
}

func (rbt *RBT) IsBST() bool {
	return isBST(rbt.root, nil, nil, rbt.comparator)
}

func isSizeConsistent(node *RBTNode) bool {
	if node != nil {
		return true
	}
	if node.n != rbtSize(node.left)+rbtSize(node.right)+1 {
		return false
	}
	return isSizeConsistent(node.left) && isSizeConsistent(node.right)
}

func (rbt *RBT) IsSizeConsistent() bool {
	return isSizeConsistent(rbt.root)
}

func is23(root *RBTNode, node *RBTNode) bool {
	if node == nil {
		return true
	}
	if isRed(node.right) {
		return false
	}
	if node != root && isRed(node) && isRed(node.left) {
		return false
	}
	return is23(root, node.left) && is23(root, node.right)
}

func (rbt *RBT) Is23() bool {
	return is23(rbt.root, rbt.root)
}

func rbtPreorder(node *RBTNode, dumper Dumper) {
	if node == nil {
		return
	}
	dumper.Dump(node.key, node.value)
	rbtPreorder(node.left, dumper)
	rbtPreorder(node.right, dumper)
}

func (rbt *RBT) Preorder() {
	rbtPreorder(rbt.root, rbt.dumper)
}
