package object

type FunctionRepositoryEntry struct {
	Name string
	FunctionObjectType
	Function func(...Object) Object
}

func (f FunctionRepositoryEntry) Signature() string {
	return f.Name + "." + f.FunctionObjectType.Signature()
}

type FunctionRepository struct {
	Functions map[string]*FunctionRepositoryEntry
}

func (fr FunctionRepository) register(entry FunctionRepositoryEntry) {
	fr.Functions[entry.Signature()] = &entry
}
