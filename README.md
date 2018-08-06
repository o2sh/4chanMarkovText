# 4chanMarkovText

Some Outputs for /biz/ /pol/ /b/ /fit/ : https://pastebin.com/sxdapK9p

How to use
----------

```cpp
go run *go -n=3 -words=12 -capital -sentence -input="./data/fit.txt"
```

  -capital  
        start output with a capitalized prefix  
  -in string  
        input file (default "./data/biz.txt")  
  -n int  
        number of words to use as prefix (default 2)  
  -sentence  
        end output at a sentence ending punctuation mark (after n words)  
  -words int  
        number of words per run (default 200)  
