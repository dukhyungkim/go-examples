package main

type Human struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (h *Human) IsEqual(human *Human) bool {
	if h == human {
		return true
	}

	if h == nil || human == nil {
		return false
	}

	if h.Name == human.Name && h.Age == human.Age {
		return true
	}

	return false
}
