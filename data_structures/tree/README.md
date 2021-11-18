# Definition

A collection of items where the structure represents a tree. The root item is the only one without parent.

* Node: every item in the tree
* Nodes have data and a list of children nodes
* Iterating through a tree = traversal
* there are different orders of traversal

## Breadth-First Traversal

Iterate every level of the tree from left to right then go to next level

```go
    1
   / \
  2   3
 / \
4   5

[1, 2, 3, 4, 5]
```

## Depth-First Traversal

Iterate from top to the bottom starting from the left and moving to the right

```go
    1
   / \
  2   3
 / \
4  5

[1, 2, 4, 5, 3]
```

# Implementation

## Node

* create a Node class that has 2 properties Data and Children
* when the constructor is called, initialize the Data prop with the argument and initialize children to an empty list
* the Add method should append the list a new node
* the Remove method should remove from the Children list all nodes that have the same Data as the passed by argument


## Tree

* create a Class that has 1 property, Root
* when the constructor is called, initialize the Root to null

## Breadth-First Traversal

* create an empty array
* add the root node to the array
* while the array is not empty
	* grab and remove the first node from the array
	* take the first node in the array and append all its children to the **end** of the array
	* call the function on the first node


## Depth-First Traversal

* create an empty array
* add the root node to the array
* while the array is not empty
	* grab and remove the first node from the array
	* take the first node in the array and append all its children to the **start** of the array
	* call the function on the first node