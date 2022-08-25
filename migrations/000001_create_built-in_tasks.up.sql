-- users
INSERT INTO users (created_at, updated_at, username, password_hash, is_disabled) VALUES
(NOW(), NOW(), 'built-in', '$2a$10$fZlvB/5ByHmJOesNhBw7sOWchjTdDbk6hRWxUDwSdakAWMTRgbTce' /* 'built-in' */, TRUE);

-- tasks
INSERT INTO tasks (created_at, updated_at, user_id, title, description, boilerplate_code, test_cases) VALUES
(
  NOW() + INTERVAL '1 minute',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Sum of Two Numbers (Test Task)',
  'Write a function that adds two numbers.',
  'package main

// TODO: finish this function
func add(x int, y int) int {
}

func main() {
  var x, y int
  fmt.Scan(&x, &y)

  sum := add(x, y)
  fmt.Println(sum)
}',
  '[
    { "Input": "5 12", "ExpectedOutput": "17\n" },
    { "Input": "23 42", "ExpectedOutput": "65\n" }
  ]'
),
(
  NOW() + INTERVAL '2 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Sum of Array',
  'Write a function that calculates the sum of an array of integers.',
  'package main

// TODO: finish this function
func sumOfArray(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  sum := sumOfArray(numbers)
  fmt.Print(sum)
}',
  '[
    { "Input": "1 2 3 4 5", "ExpectedOutput": "15" },
    { "Input": "1 -2 3 -4 5", "ExpectedOutput": "3" },
    { "Input": "-1 -2 -3 -4 -5", "ExpectedOutput": "-15" },
    { "Input": "", "ExpectedOutput": "0" },
    { "Input": "1 1 2 -3 4 -5", "ExpectedOutput": "0" },
    { "Input": "1 0 3 0 5", "ExpectedOutput": "9" },
    { "Input": "0 0 0 0 0", "ExpectedOutput": "0" }
  ]'
),
(
  NOW() + INTERVAL '3 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Product of Array',
  'Write a function that calculates the product of an array of integers. Note that the product of an empty array is 1: https://en.wikipedia.org/wiki/Empty_product',
  'package main

// TODO: finish this function
func productOfArray(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  product := productOfArray(numbers)
  fmt.Print(product)
}',
  '[
    { "Input": "1 2 3 4 5", "ExpectedOutput": "120" },
    { "Input": "1 -2 3 -4 5", "ExpectedOutput": "120" },
    { "Input": "-1 2 -3 4 -5", "ExpectedOutput": "-120" },
    { "Input": "-1 -2 -3 -4 -5", "ExpectedOutput": "-120" },
    { "Input": "", "ExpectedOutput": "1" },
    { "Input": "0 1 2 3 4 5", "ExpectedOutput": "0" },
    { "Input": "1 0 3 0 5", "ExpectedOutput": "0" },
    { "Input": "0 0 0 0 0", "ExpectedOutput": "0" },
    { "Input": "1 2 1 4 1", "ExpectedOutput": "8" },
    { "Input": "1 1 1 1 1", "ExpectedOutput": "1" }
  ]'
),
(
  NOW() + INTERVAL '4 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Minimum of Array',
  'Write a function that finds the minimum element of an array. You can assume that the array is not empty.',
  'package main

// TODO: finish this function
func findMinimumOfArray(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  minimum := findMinimumOfArray(numbers)
  fmt.Print(minimum)
}',
  '[
    { "Input": "1 2 3 4 5", "ExpectedOutput": "1" },
    { "Input": "5 4 3 2 1", "ExpectedOutput": "1" },
    { "Input": "5 3 1 2 4", "ExpectedOutput": "1" },
    { "Input": "5 1 3 1 2 1 4", "ExpectedOutput": "1" },
    { "Input": "1 1 1 1 1", "ExpectedOutput": "1" },
    { "Input": "1", "ExpectedOutput": "1" }
  ]'
),
(
  NOW() + INTERVAL '5 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Maximum of Array',
  'Write a function that finds the maximum element of an array. You can assume that the array is not empty.',
  'package main

// TODO: finish this function
func findMaximumOfArray(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  maximum := findMaximumOfArray(numbers)
  fmt.Print(maximum)
}',
  '[
    { "Input": "5 4 3 2 1", "ExpectedOutput": "5" },
    { "Input": "1 2 3 4 5", "ExpectedOutput": "5" },
    { "Input": "1 3 5 4 2", "ExpectedOutput": "5" },
    { "Input": "1 5 3 5 4 5 2", "ExpectedOutput": "5" },
    { "Input": "5 5 5 5 5", "ExpectedOutput": "5" },
    { "Input": "5", "ExpectedOutput": "5" }
  ]'
),
(
  NOW() + INTERVAL '6 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Index of Array Minimum',
  'Write a function that finds the index of the minimum element of an array. If the minimum element occurs more than once in the array, return the index of the first occurrence. You can assume that the array is not empty.',
  'package main

// TODO: finish this function
func findIndexOfMinimumOfArray(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  indexOfMinimum := findIndexOfMinimumOfArray(numbers)
  fmt.Print(indexOfMinimum)
}',
  '[
    { "Input": "1 2 3 4 5", "ExpectedOutput": "0" },
    { "Input": "5 4 3 2 1", "ExpectedOutput": "4" },
    { "Input": "5 3 1 2 4", "ExpectedOutput": "2" },
    { "Input": "5 1 3 1 2 1 4", "ExpectedOutput": "1" },
    { "Input": "1 1 1 1 1", "ExpectedOutput": "0" },
    { "Input": "1", "ExpectedOutput": "0" }
  ]'
),
(
  NOW() + INTERVAL '7 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Index of Array Maximum',
  'Write a function that finds the index of the maximum element of an array. If the maximum element occurs more than once in the array, return the index of the first occurrence. You can assume that the array is not empty.',
  'package main

// TODO: finish this function
func findIndexOfMaximumOfArray(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  indexOfMaximum := findIndexOfMaximumOfArray(numbers)
  fmt.Print(indexOfMaximum)
}',
  '[
    { "Input": "5 4 3 2 1", "ExpectedOutput": "0" },
    { "Input": "1 2 3 4 5", "ExpectedOutput": "4" },
    { "Input": "1 3 5 4 2", "ExpectedOutput": "2" },
    { "Input": "1 5 3 5 4 5 2", "ExpectedOutput": "1" },
    { "Input": "5 5 5 5 5", "ExpectedOutput": "0" },
    { "Input": "5", "ExpectedOutput": "0" }
  ]'
),
(
  NOW() + INTERVAL '8 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Search in Array',
  'Write a function that finds the index of the required element in the array. If that element occurs more than once in the array, return the index of the first occurrence. If that element is absent in the array, return -1.',
  'package main

// TODO: finish this function
func findNumberInArray(numbers []int, requiredNumber int) int {
}

func main() {
  var totalNumbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    totalNumbers = append(totalNumbers, number)
  }

  numbers, requiredNumber :=
    totalNumbers[:len(totalNumbers)-1], totalNumbers[len(totalNumbers)-1]
  indexOfRequiredNumber := findNumberInArray(numbers, requiredNumber)
  fmt.Print(indexOfRequiredNumber)
}',
  '[
    { "Input": "1 2 3 4 5 1", "ExpectedOutput": "0" },
    { "Input": "1 2 3 4 5 5", "ExpectedOutput": "4" },
    { "Input": "1 2 3 4 5 3", "ExpectedOutput": "2" },
    { "Input": "1 3 2 3 4 3 5 3", "ExpectedOutput": "1" },
    { "Input": "3 3 3 3 3 3", "ExpectedOutput": "0" },
    { "Input": "1 2 3 4 5 6", "ExpectedOutput": "-1" },
    { "Input": "1", "ExpectedOutput": "-1" }
  ]'
),
(
  NOW() + INTERVAL '9 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Search in Array (Last Occurrence)',
  'Write a function that finds the index of the required element in the array. If that element occurs more than once in the array, return the index of the last occurrence. If that element is absent in the array, return -1.',
  'package main

// TODO: finish this function
func findLastNumberInArray(numbers []int, requiredNumber int) int {
}

func main() {
  var totalNumbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    totalNumbers = append(totalNumbers, number)
  }

  numbers, requiredNumber :=
    totalNumbers[:len(totalNumbers)-1], totalNumbers[len(totalNumbers)-1]
  indexOfRequiredNumber := findLastNumberInArray(numbers, requiredNumber)
  fmt.Print(indexOfRequiredNumber)
}',
  '[
    { "Input": "1 2 3 4 5 1", "ExpectedOutput": "0" },
    { "Input": "1 2 3 4 5 5", "ExpectedOutput": "4" },
    { "Input": "1 2 3 4 5 3", "ExpectedOutput": "2" },
    { "Input": "1 3 2 3 4 3 5 3", "ExpectedOutput": "5" },
    { "Input": "3 3 3 3 3 3", "ExpectedOutput": "4" },
    { "Input": "1 2 3 4 5 6", "ExpectedOutput": "-1" },
    { "Input": "1", "ExpectedOutput": "-1" }
  ]'
),
(
  NOW() + INTERVAL '10 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Filtering of Array',
  'Write a function that creates a copy of the array without the specified element.',
  'package main

// TODO: finish this function
func filterArray(numbers []int, unwantedNumber int) []int {
}

func main() {
  var totalNumbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    totalNumbers = append(totalNumbers, number)
  }

  numbers, unwantedNumber :=
    totalNumbers[:len(totalNumbers)-1], totalNumbers[len(totalNumbers)-1]
  filteredNumbers := filterArray(numbers, unwantedNumber)

  var filteredNumbersForOutput []interface{}
  for _, number := range filteredNumbers {
    filteredNumbersForOutput = append(filteredNumbersForOutput, number)
  }

  fmt.Print(filteredNumbersForOutput...)
}',
  '[
    { "Input": "1 2 3 4 5 1", "ExpectedOutput": "2 3 4 5" },
    { "Input": "1 2 3 4 5 5", "ExpectedOutput": "1 2 3 4" },
    { "Input": "1 2 3 4 5 3", "ExpectedOutput": "1 2 4 5" },
    { "Input": "1 3 2 3 4 3 5 3", "ExpectedOutput": "1 2 4 5" },
    { "Input": "3 3 3 3 3 3", "ExpectedOutput": "" },
    { "Input": "1 2 3 4 5 6", "ExpectedOutput": "1 2 3 4 5" },
    { "Input": "1", "ExpectedOutput": "" }
  ]'
),
(
  NOW() + INTERVAL '11 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Filtering of Array (In-place)',
  'Write a function that removes the specified element from the array (in-place) and returns the new array length.',
  'package main

// TODO: finish this function
func filterArrayInPlace(numbers []int, unwantedNumber int) int {
}

func main() {
  var totalNumbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    totalNumbers = append(totalNumbers, number)
  }

  numbers, unwantedNumber :=
    totalNumbers[:len(totalNumbers)-1], totalNumbers[len(totalNumbers)-1]
  filteredNumberCount := filterArrayInPlace(numbers, unwantedNumber)
  filteredNumbers := numbers[:filteredNumberCount]

  var filteredNumbersForOutput []interface{}
  for _, number := range filteredNumbers {
    filteredNumbersForOutput = append(filteredNumbersForOutput, number)
  }

  fmt.Print(filteredNumbersForOutput...)
}',
  '[
    { "Input": "1 2 3 4 5 1", "ExpectedOutput": "2 3 4 5" },
    { "Input": "1 2 3 4 5 5", "ExpectedOutput": "1 2 3 4" },
    { "Input": "1 2 3 4 5 3", "ExpectedOutput": "1 2 4 5" },
    { "Input": "1 3 2 3 4 3 5 3", "ExpectedOutput": "1 2 4 5" },
    { "Input": "3 3 3 3 3 3", "ExpectedOutput": "" },
    { "Input": "1 2 3 4 5 6", "ExpectedOutput": "1 2 3 4 5" },
    { "Input": "1", "ExpectedOutput": "" }
  ]'
),
(
  NOW() + INTERVAL '12 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Removing of Duplicates',
  'Write a function that creates a copy of an array with no repeating elements.',
  'package main

// TODO: finish this function
func removeDuplicates(numbers []int) []int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  cleanedNumbers := removeDuplicates(numbers)

  var cleanedNumbersForOutput []interface{}
  for _, number := range cleanedNumbers {
    cleanedNumbersForOutput = append(cleanedNumbersForOutput, number)
  }

  fmt.Print(cleanedNumbersForOutput...)
}',
  '[
    { "Input": "1 2 3 4 5", "ExpectedOutput": "1 2 3 4 5" },
    { "Input": "5 1 3 1 2 1 4", "ExpectedOutput": "5 1 3 2 4" },
    { "Input": "1 1 1 1 1", "ExpectedOutput": "1" },
    { "Input": "1", "ExpectedOutput": "1" },
    { "Input": "", "ExpectedOutput": "" }
  ]'
),
(
  NOW() + INTERVAL '13 minutes',
  NOW(),
  CURRVAL(PG_GET_SERIAL_SEQUENCE('users', 'id')),
  'Removing of Duplicates (In-place)',
  'Write a function that removes repeating elements from the array (in-place) and returns the new array length.',
  'package main

// TODO: finish this function
func removeDuplicatesInPlace(numbers []int) int {
}

func main() {
  var numbers []int
  for {
    var number int
    if _, err := fmt.Scan(&number); err != nil {
      break
    }

    numbers = append(numbers, number)
  }

  cleanedNumberCount := removeDuplicatesInPlace(numbers)
  cleanedNumbers := numbers[:cleanedNumberCount]

  var cleanedNumbersForOutput []interface{}
  for _, number := range cleanedNumbers {
    cleanedNumbersForOutput = append(cleanedNumbersForOutput, number)
  }

  fmt.Print(cleanedNumbersForOutput...)
}',
  '[
    { "Input": "1 2 3 4 5", "ExpectedOutput": "1 2 3 4 5" },
    { "Input": "5 1 3 1 2 1 4", "ExpectedOutput": "5 1 3 2 4" },
    { "Input": "1 1 1 1 1", "ExpectedOutput": "1" },
    { "Input": "1", "ExpectedOutput": "1" },
    { "Input": "", "ExpectedOutput": "" }
  ]'
);
