def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

class PrimeChecker:
    def __init__(self):
        self.primes = []

    def add_prime(self, number):
        while number > 1:
            if is_prime(number):
                self.primes.append(number)
            number -= 1
        return self.primes

if __name__ == "__main__":
    checker = PrimeChecker()
    checker.add_prime(20)  # Add a prime number to the list of primes
    print(checker.primes)   # Print the list of primes
