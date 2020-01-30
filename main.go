package main

import (
	"github.com/iesreza/foundation-example/components/sample"
	"github.com/iesreza/foundation/system"
)

func main()  {
	system.PreBoot()

	//Register modules
	sample.Register()

	system.Boot()
	system.ListenCLI()
}
