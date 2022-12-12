def main():
    data = open("input.txt")
    score = 0

    for line in data:
        opp = ord(line[0]) - 64
        play = ord(line[2]) - 87

        if opp == play:
            score += 3
        elif not (opp == 1 and play == 3 or opp - 1 == play):
            score += 6

        score += play

    print("Score:", score)

if __name__ == "__main__":
    main()
