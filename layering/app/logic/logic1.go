package logic

func logic1(experiments map[string]string) {
	if !needLogic1Run(experiments) {
		return
	}

	// Далее логика не зависит от experiments

	prepareLogic1Params()
	logic1Stage1()
	logic1Stage2()
}

func needLogic1Run(experiments map[string]string) bool {
	if experiments["A"] == "1" {
		return true
	}
	return false
}

func prepareLogic1Params() {

}

func logic1Stage1() {
	if true {
		if true {
			if true {

			}
		}
	}
}

func logic1Stage2() {

}
