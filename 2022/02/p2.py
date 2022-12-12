def main():
    data = open("input.txt")
    score = 0

    for line in data:
        opp = ord(line[0]) - 64
        win = line[2]

        if win == 'X':
            score += opp - 1 if opp > 1 else 3
        elif win == 'Z':
            score += 6 + (opp + 1 if opp < 3 else 1)
        else:
            score += opp + 3

    print("Score:", score)

if __name__ == "__main__":
    main()
