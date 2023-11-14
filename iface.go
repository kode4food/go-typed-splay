package splaytree

// Interface defines the public interface of splay trees
type Interface[Item any] interface {
	// Remove all elements of the tree
	Clear()
	// Count the number of elements in the tree
	Count() int
	// Remove an item from the tree
	Delete(item Item) (Item, bool)
	// Remove all given items from the tree
	DeleteAll(items []Item) int
	// Remove the largest element of the tree
	DeleteMax() (Item, bool)
	// Remove the smallest element of the tree
	DeleteMin() (Item, bool)
	// Remove and return the element at the root
	DeleteRoot() (Item, bool)
	// Clone the tree
	Duplicate() Interface[Item]
	// Compute the height of the tree
	Height() int
	// Infimum gives the closest smaller item
	Infimum(item Item) (Item, bool)
	// Add a new item to the tree (if unique)
	Insert(item Item) bool
	// Add a number of items to the tree
	InsertAll(items []Item) int
	// Create a new iterator over this tree.
	Iterator() func() (Item, bool)
	// Join two trees with cost O(N + M), which is optimal.
	Join(other Interface[Item])
	// Lookup an item
	Lookup(item Item) (Item, bool)
	// Give the largest element
	Max() (Item, bool)
	// Give the smallest element
	Min() (Item, bool)
	// Test if the tree contains at least one element
	NonEmpty() bool
	// Create an iterator for a limited range of items
	RangeIterator(lower Item, upper Item) func() (Item, bool)
	// Insert or replace an element
	Replace(item Item) bool
	// Insert or replace a number of elements
	ReplaceAll(items []Item) int
	// Create an iterator which iterates in descending order
	ReverseIterator() func() (Item, bool)
	// Give the current root element
	Root() (Item, bool)
	// Supremum gives the closest larger item
	Supremum(item Item) (Item, bool)
	// Traverse a tree inorder with a function
	Traverse(func(Item))
}
