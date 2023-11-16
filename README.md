# go-typed-splay

[![GoDoc](https://godoc.org/github.com/kode4food/go-typed-splay?status.svg)](https://godoc.org/github.com/kode4food/go-typed-splay)
[![Report](https://goreportcard.com/badge/github.com/kode4food/go-typed-splay)](https://goreportcard.com/report/github.com/kode4food/go-typed-splay)

The splay tree data structure in Go, with generic type support. 

[Splay trees](https://en.wikipedia.org/wiki/Splay_tree)
are self-balancing binary search trees.
Accessing a node in the tree causes it and its neighbors
to be splayed upwards by means of rotations.
Future accesses to this node or to its neighbors will then be cheaper.
Splay trees perform best when there is temporal or spatial
[locality](https://en.wikipedia.org/wiki/Locality_of_reference)
in the access patterns (insertions, lookups or removals).
Say, you enter person data by date+time of birth and then
you investigate specific age groups.  Or you manage virtual
memory pages and your applications tend to focus on groups
of related pages. Over time they will abandon some groups
of memory pages and move on to other groups of memory
pages. The implied locality of reference in these examples
indicate splay trees as the data structure of choice.

The worst case cost for a single access (insertions,
lookups or removals) is O(n), but amortized (averaged) over
a sequence of accesses, while starting from an empty tree,
the cost will be O(log n).  Inserting a sorted sequence
of 'n' elements into an empty tree has total cost O(n),
which is optimal.  Retrieving that same sequence of 'n'
elements in either ascending or descending order has total
cost O(n), which is again optimal.

## References

["Performance Analysis of BSTs in System Software"](http://benpfaff.org/papers/libavl.pdf)
by Ben Pfaff, compares the performance of 20 variants
of binary search trees (BSTs) and concludes that splay trees
perform best when accesses are sequential or clustered.

## Credits and Changes

This is a fork of Bert Gijsbers' [go-splaytree](https://github.com/gijsbers/go-splaytree) intended to support Go generics. As a result, some necessary changes have been made to the API and interface. They are enumerated as follows:

* The `Item` interface that the tree managed has been replaced by a generic type. The original interface required the programmer to implement a `Less` method. This has been replaced by a comparator function being provided to the `NewSplayTree` constructor. 
* The original implementation checked for nil Items and expected the programmer to check for nil Items. Because nil is a perfectly valid result for a generic type, this access pattern has been replaced in many places with one that adds a bool result to signal a successful operation. 

## Installation

        $ go get github.com/kode4food/go-typed-splay

## Copyright and License

Copyright ©2016 Bert Gijsbers except where otherwise noted. All rights reserved.
Use of this source code is governed by an Apache 2.0 license
that can be found in the LICENSE file.
