from typing import Optional

def greet(name: Optional[str] = None) -> None:
    if name:
        print(f"Hello, {name}")
    else:
        print("Hello, World!")

def sum (a, b):
    return a + b

def main():
    greet()
    sum(4, 8)

if __name__ == "__main__":
    main()
