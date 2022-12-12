from itertools import islice

def main():
    data = open('input.txt')
    ans = 0

    while True:
        group = list(islice(data, 3))
        if not group: break

        seen = list()

        for item in group[0].strip():
            if item in group[1] and item in group[2] and not item in seen:
                ans += ord(item) - 96 if ord(item) >= 97 else ord(item) - 64 + 26
                seen.append(item)
                print(item, ord(item) - 96 if ord(item) >= 97 else ord(item) - 64 + 26)

    print(ans)

if __name__ == '__main__':
    main()
