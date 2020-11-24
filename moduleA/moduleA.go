package moduleA

type ModuleA struct {
	name        string
	packageName string
}

func CreateModuleA(name string, packageName string) *ModuleA  {
	return &ModuleA{
		name:        name,
		packageName: packageName,
	}
}

func (module *ModuleA) GetName() string  {
	return module.name
}

func (module *ModuleA) GetPackage() string  {
	return module.packageName
}