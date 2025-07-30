import datetime

Date = datetime.datetime
fmts = "%y/%m/%d", "%m/%d/%y", "%d/%m/%y"


def main():
    date = input()
    dates = []
    for fmt in fmts:
        try:
            date = Date.strptime(date, fmt)
            if date.year > 2059:
                date = date.replace(year=1900 + date.year % 100)
            if date not in dates:
                dates.append(date)
        except:
            pass
    for date in sorted(dates):
        print(Date.strftime(date, "%Y-%m-%d"))


if __name__ == '__main__':
    main()
