def main():
    data = open('input.txt')
    counts = list()
    cals = 0


    for line in data:
        if line == "\n":
            counts.append(cals)
            cals = 0
        else:
            cals += int(line[:len(line) - 1])

    counts.sort(reverse=True)
    print(f'Top 3 calorie counts are {counts[0:3]}, total {sum(counts[0:3])}')

    data.close()

if __name__ == "__main__":
    main()
