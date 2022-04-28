package rules

type Rule interface {
	Compute() Rule
	Save()
}

var rules = []Rule{
	&SignatureDate{},
	&StartDate{},
	&EndDate{},
	&PeriodicPayment{},
	&Termination{},
}
