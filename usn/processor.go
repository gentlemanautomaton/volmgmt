package usn

// A Processor is capable of processing USN change journal records.
type Processor func(Record)

// Process runs p(r) if p is non-nil.
func (p Processor) Process(r Record) {
	if p != nil {
		p(r)
	}
}
