package routine

type Routine interface {
	Next() (position, speed, delay int)
	Done() bool
}
