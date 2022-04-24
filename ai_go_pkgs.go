package ai_go_pkgs

import "fmt"

type GoEngineer struct {
	name, company, country string
}

func Sample() {
	fmt.Println("Hello world")
}

func (ge GoEngineer) PrintDetails() {
	fmt.Printf("My name is %v, developer at %v , from %v", ge.name, ge.company, ge.country)
	fmt.Println()
}

func PrintGoEnginners(ge []GoEngineer) {
	goEngineers := []GoEngineer{
		{
			name:    "Antony Injila",
			company: "Ariya Finergy",
			country: "Kenya",
		},
		{
			name:    "Mohameed Mwiji",
			company: "Lori Systems",
			country: "Kenya",
		},
	}
	for _, engineer := range goEngineers {
		engineer.PrintDetails()
	}
}
