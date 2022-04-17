package math

// Найти среднее в массиве чисел.

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Найти сумму в массиве.
func Sum(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}

func Max(xs []float64) float64 {
	max: = xs[0]
	for x := range xs {
		if xs[x] > max {
			max = xs[x]
		}

	}
	return max
}
