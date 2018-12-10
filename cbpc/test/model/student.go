package model

//Student xuesheng
type student struct {
	Name string
	Score float32
}

//NewStudent new stu
func NewStudent(name string,score float32) *student {
	return &student{
		Name:name,
		Score:score,
	}
}