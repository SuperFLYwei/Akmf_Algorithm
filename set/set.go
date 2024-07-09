package set

// New gives new set
func New[T comparable](items ...T) Set[T] {
	set := set[T]{
		elements: make(map[T]bool),
	}
	for _, item := range items {
		set.Add(item)
	}
	return &set
}

type Set[T comparable] interface {
	// Add: add a new element to the set
	Add(item T)
	// Delete: delete the passed element from the set if present
	Delete(item T)
	// Len: gives the length of the set
	Len() int
	// GetItems: gives the array of elements of the set
	GetItems() []T
	// In: checks whether item is present in set or not
	In(item T) bool
	// IsSubsetOf: checks whether item is present in set or not
	IsSubsetOf(set2 Set[T]) bool
	// IsProperSubsetOf: checks whether set is proper subset of set2 or not.
	// ex: [1,2,3] proper subset of [1,2,3,4] -> true
	IsProperSubsetOf(set2 Set[T]) bool
	// Union: gives new union set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,3,4,5]
	Union(set2 Set[T]) Set[T]
	// Intersection: gives new intersection set of both sets.
	// ex: [1,2,3] Intersection [3,4,5] -> [3]
	Intersection(set2 Set[T]) Set[T]
	// Difference: gives new difference set of both sets.
	// ex: [1,2,3] Difference [3,4,5] -> [1,2]
	Difference(set2 Set[T]) Set[T]
	// SymmetricDifference: gives new symmetric difference set of both sets.
	// ex: [1,2,3] SymmetricDifference [3,4,5] -> [1,2,4,5]
	SymmetricDifference(set2 Set[T]) Set[T]
}

type set[T comparable] struct {
	elements map[T]bool
}

func (set *set[T]) Add(value T) {
	set.elements[value] = true
}

func (set *set[T]) Delete(value T) {
	delete(set.elements, value)
}

func (set *set[T]) Len() int {
	return len(set.elements)
}

func (set *set[T]) GetItems() []T {
	keys := make([]T, 0, len(set.elements))
	for k := range set.elements {
		keys = append(keys, k)
	}
	return keys
}

func (set *set[T]) In(value T) bool {
	if _ , in := set.elements[value]; in {
		return true
	}
	return false
}

func (set *set[T]) IsSubsetOf(set2 Set[T]) bool {
	if set.Len() > set2.Len() {
		return false
	}
	for _, item := range set.GetItems() {
		if !set2.In(item) {
			return false
		}
	}
	return true
}

func (set *set[T]) IsProperSubsetOf(set2 Set[T]) bool {
	if set.Len() == set2.Len() {
		return false
	}
	return set.IsSubsetOf(set2)
}

func (set *set[T]) Union(set2 Set[T]) Set[T] {
	unionSet := New[T]()
	for _, item := range set.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range set2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (set *set[T]) Intersection(st2 Set[T]) Set[T] {
	intersectionSet := New[T]()
	var minSet, maxSet Set[T]
	if set.Len() > st2.Len() {
		minSet = st2
		maxSet = set
	} else {
		minSet = set
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

func (set *set[T]) Difference(set2 Set[T]) Set[T] {
	differenceSet := New[T]()
	for _, item := range set.GetItems() {
		if !set2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (set *set[T]) SymmetricDifference(set2 Set[T]) Set[T] {
	symmetricDifferenceSet := New[T]()
	dropSet := New[T]()
	for _, item := range set.GetItems() {
		if set2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}
	for _, item := range set2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}
	return symmetricDifferenceSet
}