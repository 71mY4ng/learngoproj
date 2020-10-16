package second

import learngo "timyang.com/learngoproj/learngo"

func Hello() string {
	return "second Help"
}

func HelloWithFirst() string {
	return "second Hello with \"" + learngo.SomeFn() + "\""
}
