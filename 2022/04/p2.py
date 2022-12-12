def overlaps(r1, r2):
    for n in range(r1[0], r1[1] + 1):
        if n >= r2[0] and n <= r2[1]: return True

    return False


def main():
    data = open('input.txt')
    ans = 0

    for line in data:
        pair = line.strip().split(',')
        elf1 = [int(n) for n in pair[0].split('-')]
        elf2 = [int(n) for n in pair[1].split('-')]

        if overlaps(elf1, elf2):
            ans += 1

    print(ans)
    data.close()

if __name__ == "__main__":
    main()
