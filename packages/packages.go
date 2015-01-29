package packages

type Package struct {
	Name string
}

func CreatePackage(packageName string) Package {

	return Package{
		Name: packageName,
	}
}
