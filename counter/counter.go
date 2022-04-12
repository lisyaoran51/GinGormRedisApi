package counter

type Counter interface {
	String() string
	Count() int
	GetCount() int
}