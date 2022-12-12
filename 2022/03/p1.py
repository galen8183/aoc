def main():
    data = open('input.txt')
    ans = 0

    for line in data:
        items = len(line.strip())
        comp1 = line[(items // 2):]
        comp2 = line[:(items // 2)]
        seen = list()

        for item in comp1:
            if item in comp2 and not item in seen:
                prio = ord(item) - 96 if ord(item) >= 97 else ord(item) - 64 + 26
                ans += prio
                seen.append(item)
                print(item, prio)

    print(ans)

if __name__ == '__main__':
    main()
