package common

import "container/heap"

// Sorts artifacts by date
type artifactHeap []*TicketArtifact

var _ heap.Interface = &artifactHeap{}

func (h artifactHeap) Len() int {
	return len(h)
}

func (h artifactHeap) Less(i,j int) bool {
	return h[i].Date.Before(h[j].Date)
}

func (h artifactHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *artifactHeap) Push(v interface{}) {
	*h = append(*h, v.(*TicketArtifact))
}

func (h *artifactHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}
