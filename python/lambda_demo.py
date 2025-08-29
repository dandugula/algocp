def main():
    people = [
        {"name": "Alice", "age": 25},
        {"name": "Chaitanya", "age": 39},
        {"name": "Joye", "age": 14},
    ]

    # use lambda to sort list of dicts by age
    people.sort(key=lambda person: person["age"])

    for person in people:
        print(f"{person['name']}: {person['age']}")


if __name__ == "__main__":
    main()
