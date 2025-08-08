from typing import Final

b: Final[int] = 1
Kb: Final[int] = 1024 * b
Mb: Final[int] = 1024 * Kb
Gb: Final[int] = 1024 * Mb

B: Final[int] = 8 * b
KB: Final[int] = 1024 * B
MB: Final[int] = 1024 * KB
GB: Final[int] = 1024 * MB


def main():
    print(round(200 / 8))


if __name__ == '__main__':
    main()
