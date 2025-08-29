import asyncio

def func1():
    print("func1")

def func2():
    print("func2")

def func3():
    print("func3")

async def main():
    results = await asyncio.gather(
        func1(),
        func2(),
        func3(),
    )
    await results

# if __name__ == "__main__":
asyncio.run(main())

