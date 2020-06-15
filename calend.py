#!/usr/bin/env python3
# sudo pip3 install feedparser

import locale
import time

import feedparser


def get_today():
    f = feedparser.parse('http://www.calend.ru/img/export/today-holidays.rss')
    locale.setlocale(locale.LC_TIME, "ru_RU")
    now = "%s %s" % (int(time.strftime("%d")), time.strftime("%B"))

    hol = now + ": "
    for i in f.entries:
        if now in i['title']:
            hol = hol + i['title'].split('-')[1].strip() + "; "
    return hol[:-2]  # remove trailing ;


def main():
    print(get_today())


if __name__ == "__main__":
    main()
