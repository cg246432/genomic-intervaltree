# genomic-intervaltree

This implements an interval tree structure in Golang, with a genomic focus. A Genomic Interval Tree is actually a collection of IntervalTrees, one for each chromosome. It is represented as a map, or dictionary, where the chromosome name is the key, and the tree is the value. A tree can be generated via a collection of intervals. 