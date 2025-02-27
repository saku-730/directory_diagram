# directory_diagram

Making directory structure diagram. Output is .txt file or standard output.

![alt text](/img/example.png)

## Install

Dawnload or Git colne. dir-diagram is compiled binary program. It can be used in most OS - Windows, Mac, Linux.

## USAGE

standard output

```
./dir-diagram 
```

txt file output. 

```
./dir-diagram -f Test.txt
```

If you don't specify any argument, root of diagram is current directory.

If you want to set a directory as the root, do the following:

```
./dir-diagram /home/user/Documents
```

## Options

- -h show hidden file. Not displayed by default.
- -t depth of file structure. Default depth is 3
- -f output to file.
