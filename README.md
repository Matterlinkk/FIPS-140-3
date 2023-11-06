# FIPS-140-3 in Go
Covers these tests:
- Monobit
- Maximum series length
- Series length
- Poker
  
### <i>Test takes 20000 bits in output</i>

## Structure description
- `operations/operations.go` - exists useful funcs(random sequence generation and func for a simpler work with Poker's test)
- `staticTests/maxlength.go` - series(length) implementation
- `staticTests/maxrun.go` - maximum series(run) test implementation
- `staticTests/monobit.go` - monobit test implementation.
- `staticTests/poker.go` - poker test implementation
- `main_test.go` - a file with autotests to check the correctness of the functionality
