package db

//Tests array Test strukturasining jamlanma shakli
type Tests []*Test

//Push method Tests array uchun
func (ts *Tests) Push(t *Test) {
	*ts = append(*ts, t)
}
