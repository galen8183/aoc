def main():
    data = open('input.txt')
    maxCals = 0
    cals = 0


    for line in data:
        if line == "\n":
            if cals > maxCals: maxCals = cals
            cals = 0
        else:
            cals += int(line[:len(line) - 1])

    print(f'Highest calorie count: {maxCals}')

    data.close()

if __name__ == "__main__":
    main()
