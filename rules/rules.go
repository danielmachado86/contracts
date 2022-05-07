package rules

type Rule interface {
	Compute() Rule
	Save()
}
