package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Id       int
	Name     string
	Children []*Node
	Parent   int
	Size     int
}

var filesystem map[int]*Node

func (n *Node) FindChild(name string) *Node {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

// to calculate the size of a node, recursively add the size of all children or its own size if it's a file

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize scanner, relational mapper, root node, and counter
	sc := bufio.NewScanner(input)
	filesystem = make(map[int]*Node)
	filesystem[0] = &Node{Id: 0, Name: "/", Size: 0}
	currentDirectory := *filesystem[0]
	// i acts as a unique value to assign to each node to refer to later, since
	// unique names are not guaranteed
	i := 1

	// loop through the input
	for sc.Scan() {
		// take the input and split it into parameters to analyze
		output := sc.Text()
		items := strings.Split(output, " ")

		// if the first parameter is a $, it's a command
		if items[0] == "$" {
			if items[1] == "cd" {
				if items[2] == ".." {
					// navigate back to the parent
					currentDirectory = *filesystem[currentDirectory.Parent]
				} else if items[2] != "/" {
					// ^ edge case for root directory behavior
					// navigate to the child present in the current directory
					currentDirectory = *filesystem[currentDirectory.Id].FindChild(items[2])
				}
			}
		} else {
			if items[0] == "dir" {
				// we're viewing a directory, add it to the filesystem and add it to the current directory's children
				filesystem[i] = &Node{Id: i, Name: items[1], Parent: currentDirectory.Id, Size: 0}
				filesystem[currentDirectory.Id].Children = append(filesystem[currentDirectory.Id].Children, filesystem[i])
			} else {
				// we're viewing a file, take its size add it to the filesystem and add it to the current directory's children
				size, _ := strconv.Atoi(items[0])
				filesystem[i] = &Node{Id: i, Name: items[1], Parent: currentDirectory.Id, Size: size}
				filesystem[currentDirectory.Id].Children = append(filesystem[currentDirectory.Id].Children, filesystem[i])
			}
		}
		// increment the counter
		i++
	}

	// calculate the size of the filesystem starting at the root
	calculateSize(filesystem[0])
	// initialize the sum and the total size allowed
	sum, totalSize := 0, 100000

	// loop through the filesystem and add the size of all directories that are less than the total size
	for _, v := range filesystem {
		if len(v.Children) != 0 && v.Size <= totalSize {
			sum += v.Size
		}
	}
	fmt.Println(sum)
}

// recursively calculate the size of a node and its children
func calculateSize(node *Node) int {
	// if the node has no children, it's a file and its size is its own size
	if len(node.Children) == 0 {
		return node.Size
	}

	size := 0
	for _, child := range node.Children {
		size += calculateSize(child)
	}

	// add the size of this node to the relational mapper
	filesystem[node.Id].Size = size
	return size
}
