def overlaps(r1, r2):
    if ((r1[0] <= r2[0] and r1[1] >= r2[1]) or
        (r1[0] >= r2[0] and r1[1] <= r2[1])):
        return True

    return False


def main():
    data = open('input.txt')
    ans = 0

    for line in data:
        pair = line.strip().split(',')
        elf1 = [int(i) for i in pair[0].split('-')]
        elf2 = [int(i) for i in pair[1].split('-')]

        if overlaps(elf1, elf2): ans += 1

    print(ans)
    data.close()

if __name__ == "__main__":
    main()
