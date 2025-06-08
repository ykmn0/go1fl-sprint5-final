package actioninfo

import "log"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Parse error: %v\n", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ActionInfo error: %v\n", err)
			continue
		}

		println(info)
	}
}
