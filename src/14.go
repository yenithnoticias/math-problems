import sys

def check_even(x):
    if x % 2 == 0:
        print("Even")
    else:
        print("Odd")

if __name__ == "__main__":
    # Example usage: check even or odd number
    num = int(input())
    check_even(num)
