def solve(data):
    buf = []

    for i, char in enumerate(data):
        if i < 4:
            buf.append(char)
            continue

        marker = False
        for j in buf:
            if buf.count(j) > 1:
                marker = True
                break
        if not marker: return i

        buf = buf[1:]
        buf.append(char)

    return 0

if __name__ == "__main__":
    # data = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
    f = open('input.txt')
    data = f.read()
    f.close

    print(solve(data))
