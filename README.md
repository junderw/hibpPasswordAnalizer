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
517238891 passwords were leaked      1 times or more
344791361 passwords were leaked      2 times or more
191602470 passwords were leaked      3 times or more
124428050 passwords were leaked      4 times or more
 87246372 passwords were leaked      5 times or more
 65670784 passwords were leaked      6 times or more
 52376695 passwords were leaked      7 times or more
 44035113 passwords were leaked      8 times or more
 37712431 passwords were leaked      9 times or more
 33256195 passwords were leaked     10 times or more
  4476068 passwords were leaked     50 times or more
  2005322 passwords were leaked    100 times or more
   410333 passwords were leaked    500 times or more
   198624 passwords were leaked   1000 times or more
    29216 passwords were leaked   5000 times or more
    12175 passwords were leaked  10000 times or more
     1595 passwords were leaked  50000 times or more
      566 passwords were leaked 100000 times or more

```
