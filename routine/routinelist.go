package routine

type RoutineNode struct {
	Val  Routine
	Next *RoutineNode
}

type RoutineList struct {
	first *RoutineNode
	last  *RoutineNode
}

func (list *RoutineList) Push(routine Routine) {
	r := new(RoutineNode)
	r.Val = routine
	r.Next = nil

	if list.first == nil {
		list.first = r
	}

	if list.last == nil {
		list.last = r
	} else {
		list.last.Next = r
	}
}

func (list *RoutineList) Pop() Routine {
	if list.first == nil {
		return nil
	}

	r := list.first
	if r == nil {
		return nil
	}

	list.first = r.Next
	return r.Val
}
