# hibpPasswordAnalizer
Analyze the password dump by frequency

# Steps

1. Download the password leak list (order by prevalence) [here](https://downloads.pwnedpasswords.com/passwords/pwned-passwords-ordered-by-count.7z)
2. Unzip using 7z
3. clone this repository
4. go build
5. run the binary with the unzipped txt file as the argument.
6. You will get two dicts printed to stdout
7. First one is "There are <value> passwords with a count equal to or higher than <key>"
8. The second one is "The <key>th most prevalent password has a count of <value>" (once every 10 million)

# Terminal

```bash
$ git clone https://github.com/junderw/hibpPasswordAnalizer.git
$ cd hibpPasswordAnalizer
$ go build
$ wget https://downloads.pwnedpasswords.com/passwords/pwned-passwords-ordered-by-count.7z
$ 7z x pwned-passwords-ordered-by-count.7z
$ ./hibpPasswordAnalizer pwned-passwords-ordered-by-count.txt
# Run at 2018/12/23
{
  "10": 33256195,
  "100": 2005322,
  "1000": 198624,
  "10000": 12175,
  "2": 344791361,
  "3": 191602470,
  "4": 124428050,
  "5": 87246372
}
{
  "10000000": 24,
  "100000000": 4,
  "110000000": 4,
  "120000000": 4,
  "130000000": 3,
  "140000000": 3,
  "150000000": 3,
  "160000000": 3,
  "170000000": 3,
  "180000000": 3,
  "190000000": 3,
  "20000000": 15,
  "200000000": 2,
  "210000000": 2,
  "220000000": 2,
  "230000000": 2,
  "240000000": 2,
  "250000000": 2,
  "260000000": 2,
  "270000000": 2,
  "280000000": 2,
  "290000000": 2,
  "30000000": 10,
  "300000000": 2,
  "310000000": 2,
  "320000000": 2,
  "330000000": 2,
  "340000000": 2,
  "350000000": 1,
  "360000000": 1,
  "370000000": 1,
  "380000000": 1,
  "390000000": 1,
  "40000000": 8,
  "400000000": 1,
  "410000000": 1,
  "420000000": 1,
  "430000000": 1,
  "440000000": 1,
  "450000000": 1,
  "460000000": 1,
  "470000000": 1,
  "480000000": 1,
  "490000000": 1,
  "50000000": 7,
  "500000000": 1,
  "510000000": 1,
  "60000000": 6,
  "70000000": 5,
  "80000000": 5,
  "90000000": 4
}
```
