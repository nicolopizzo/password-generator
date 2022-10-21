package main

func Map[A any, B any](f func(a A) B, xs []A) []B {
	res := make([]B, 0)
	for _, x := range xs {
		res = append(res, f(x))
	}

	return res
}
