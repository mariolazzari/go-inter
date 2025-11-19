package sort

type Sortable interface {
	Less(i, j int) bool
	Swao(i, j int)
	Len() int
}

func Sort(s Sortable) {
	// TODO: sort
}
