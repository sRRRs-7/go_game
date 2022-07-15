package load

func Collision() {
	for _, e := range Enemies {
		Restart(&e.position)
	}
}