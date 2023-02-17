package tools

// https://github.com/ernestosuarez/itertools combinations permutations

// GenIndexes generates, from two natural numbers n > r,
// all the possible combinations of r indexes taken from 0 to n-1.
// For example if n=3 and r=2, the result will be:
// [0,1], [0,2] and [1,2]
func GenIndexes(maxIndex, numberOfIndexes int) <-chan []int {
	if numberOfIndexes > maxIndex {
		panic("Invalid arguments")
	}

	ch := make(chan []int)

	go func() {
		runningIndexes := make([]int, numberOfIndexes)
		// initialize the running indexes like [0,1,2]
		for i := range runningIndexes {
			runningIndexes[i] = i
		}

		first := make([]int, numberOfIndexes)
		copy(first, runningIndexes)
		ch <- first

		for {
			// for each position in the result, starting with the last one
			for i := numberOfIndexes - 1; i >= 0; i-- {

				// as long as the number is lower than the max index
				// we are looking for minus the index in the result
				if runningIndexes[i] < (i-numberOfIndexes)+maxIndex {

					runningIndexes[i]++

					// increase the value in the positions behind
					// our current position, as they cannot be lower
					for j := 1; j < numberOfIndexes-i; j++ {
						runningIndexes[i+j] = runningIndexes[i] + j
					}

					result := make([]int, numberOfIndexes)
					copy(result, runningIndexes)

					ch <- result
					break
				}
			}
			// if the first position reaches the maximum value
			if runningIndexes[0] >= maxIndex-numberOfIndexes {
				break
			}
		}
		close(ch)

	}()
	return ch
}

// CombinationsInt generates all the combinations of r elements
// extracted from an slice of integers
// use it as a generator:
//
//	for v := range CombinationsInt(iterable, r) {
//	     fmt.Println(v)
//	}
func CombinationsInt(iterable []int, length int) chan []int {
	ch := make(chan []int)

	go func() {

		for indexes := range GenIndexes(len(iterable), length) {

			result := make([]int, length)
			for i, index := range indexes {
				result[i] = iterable[index]
			}

			ch <- result
		}

		close(ch)
	}()

	return ch
}

// CombinationsStr generates all the combinations of r elements
// extracted from an slice of strings
func CombinationsStr(iterable []string, length int) chan []string {
	ch := make(chan []string)

	go func() {
		for indexes := range GenIndexes(len(iterable), length) {
			result := make([]string, length)
			for i, val := range indexes {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}

// GenMapInts is the generator version of MapInts which can partially apply a function if you break early
func GenMapInts(ints []int, f func(index, value int) int) chan []int {
	ch := make(chan []int)

	length := len(ints)

	newInts := make([]int, length)
	copy(newInts, ints)

	go func() {
		for i, v := range ints {
			newInts[i] = f(i, v)

			// copy to prevent override
			temp := make([]int, length)
			copy(temp, newInts)
			ch <- temp
		}
		close(ch)
	}()
	return ch
}
