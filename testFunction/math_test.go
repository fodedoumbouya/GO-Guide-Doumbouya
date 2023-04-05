package math

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

var subTests = []addTest{
	{4, 2, 2},
	{10, 2, 8},
	{50, 2, 48},
	{3, 3, 0},
}

func TestAdd(t *testing.T) {

	//------------------Add
	gotAdd := Add(4, 6)
	wantAdd := 10
	CheckTestInt(gotAdd, wantAdd, t)

	for _, v := range addTests {
		CheckTestInt(Add(v.arg1, v.arg2), v.expected, t)
	}

	//-------------------------Sub
	gotSub := Subtract(40, 30)
	wantSub := 10
	CheckTestInt(gotSub, wantSub, t)

	for _, v := range subTests {
		CheckTestInt(Subtract(v.arg1, v.arg2), v.expected, t)
	}

}

func CheckTestInt(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
