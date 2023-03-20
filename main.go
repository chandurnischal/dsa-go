package main

import "dsa/linkedlist"

func main() {
	ll := linkedlist.LinkedList{}

	ll.Append(4)
	ll.Prepend(3)
	ll.PrintList()

}
