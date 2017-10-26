#!/usr/bin/env python
# sudo pip install feedparser
# pep8 --ignore=W191,E501 calend.py

import feedparser
import time
import locale
import sys


def get_today():
	f = feedparser.parse('http://www.calend.ru/img/export/today-holidays.rss')
	locale.setlocale(locale.LC_TIME, "ru_RU")
	reload(sys)
	sys.setdefaultencoding('utf8')
	now = time.strftime("%d %B")

	hol = now + ": "
	for i in f.entries:
		if now in i['title']:
			hol = hol + i['title'].split('-')[1].strip() + "; "
	return hol[:-2]  # remove trailing ;


def main():
	print get_today()

if __name__ == "__main__":
	main()
