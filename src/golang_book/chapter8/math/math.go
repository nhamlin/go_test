package math

// Average is capitalized, which allows it to be visible to other packages
// If we had named it average, it wouldn't be visible.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func random(xs []float64) float64 {
	return 1
}
