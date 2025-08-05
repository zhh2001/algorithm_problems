import hashlib
import random
import string


def sha256(text: str) -> str:
    return hashlib.sha256(text.encode()).hexdigest()


enc_len = 0
enc_key = ''


def H(s: str) -> str:
    combined = enc_key + s
    return sha256(combined)[:enc_len]


def solve() -> tuple[str, str]:
    mapping = {}
    while True:
        s = random.choices(string.ascii_lowercase, k=enc_len)
        s = ''.join(s)
        h = H(s)
        if h in mapping and mapping[h] != s:
            return mapping[h], s
        mapping[h] = s


def main():
    global enc_len, enc_key
    enc_len = int(input())
    enc_key = input()

    s1, s2 = solve()
    print(f'{s1} {s2}')


if __name__ == '__main__':
    main()
