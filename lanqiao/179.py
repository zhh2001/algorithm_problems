import sys


def violent(logs, d, k):
    """暴力解法"""
    hots = set()
    for index, log in enumerate(logs):
        count = 0
        for _index, _log in enumerate(logs):
            if log['ts'] <= _log['ts'] < log['ts'] + d and _log['id'] == log['id']:
                count += 1
        if count >= k:
            hots.add(log['id'])
    hots = list(hots)
    hots.sort()
    return hots


def double_pointer(logs, d, k):
    """双指针解法"""
    hots = set()
    logs = sorted(logs, key=lambda log: log['ts'])

    i = j = 0
    total = {}
    while True:
        while j < len(logs) and logs[j]['ts'] < logs[i]['ts'] + d:
            total[logs[j]['id']] = total.get(logs[j]['id'], 0) + 1
            j += 1

        for _id, _count in total.items():
            if _count >= k:
                hots.add(_id)

        if j == len(logs):
            break

        while logs[j]['ts'] >= logs[i]['ts'] + d:
            total[logs[i]['id']] -= 1
            if total[logs[i]['id']] == 0:
                del total[logs[i]['id']]
            i += 1

    hots = list(hots)
    hots.sort()
    return hots


def main():
    line = sys.stdin.readline().rstrip()
    n, d, k = map(int, line.split())
    logs = []
    for _ in range(n):
        s = sys.stdin.readline().rstrip()
        ts, id_ = map(int, s.split())
        log = {
            'ts': ts,
            'id': id_
        }
        logs.append(log)

    print('====== 暴力解法 ======')
    hots = violent(logs, d, k)
    for hot in hots:
        print(hot)

    print('====== 双指针解法 ======')
    hots = double_pointer(logs, d, k)
    for hot in hots:
        print(hot)


if __name__ == '__main__':
    main()
