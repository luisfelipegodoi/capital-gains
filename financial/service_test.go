package financial

import "testing"

func TestCalculateCapitalGains(t *testing.T) {

	t.Run("when the file contains a buy with zero tax and sells with zero tax and total values less than 20000",
		func(t *testing.T) {

		})

	t.Run("when the file contains a buy with zero tax, a sell with 10000 tax and a sell with loss equal 25000",
		func(t *testing.T) {

		})

	t.Run("when the file contains a buy with zero tax, a sell with 25000 loss and a sell with tax equal 1000",
		func(t *testing.T) {

		})

	t.Run("when the file contains a buy with zero tax, a buy with zero tax and a sell with non gains",
		func(t *testing.T) {

		})

	t.Run("when the file contains a buy with zero tax, a buy with zero tax, a sell with non gains and a sell with 10000 tax",
		func(t *testing.T) {

		})

	t.Run("when the file contains a buy with zero tax, a sell with zero tax, a sells with zero tax and a last sell with tax equal 3000",
		func(t *testing.T) {

		})

	t.Run("when the file contains a buy with zero tax, a sell with 80000 tax, a buy with zero tax and a sell with tax equal 60000",
		func(t *testing.T) {

		})
}
