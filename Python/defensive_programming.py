def fizzBuzz(number):
    assert type(number) == int, "Fizz Buzz only takes integer arguments"
    if number % 3 == 0 and number % 5 ==0:
        return "Fizz Buzz"
    elif number % 3 == 0:
        return "Fizz"
    elif number % 5 ==0:
        return "Buzz"
    else:
        return number

passingValue =  fizzBuzz("10")
print(passingValue)