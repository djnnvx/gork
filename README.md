# gork

CLI to fish for files with google dorks

## install

If you have go in your `GOPATH`:

```bash
go install github.com/bogdzn/gork@latest
```

To build from source:

```bash
git clone git@git.github.com/bogdzn/gork.git

cd gork

make

# or, with Docker
make docker


# if you wish to show the current available recipes
make help
```

With docker:

```
docker pull bogdzn/gork:canon
```

## usage

```bash
gork is a CLI to perform Google dorks in order to retrieve cool files :)~

Usage:
  gork [flags]

Flags:
  -a, --append-results           append dork results to out file
  -e, --extensions stringArray   filetype extensions (default [doc,docx,csv,pdf,txt,log,bak,json,xlsx])
  -h, --help                     help for gork
  -o, --outfile string           directory storing dorks results (default "./gork.txt")
  -p, --proxy string             proxy URL
  -t, --target string            target site for your dorks
  -u, --user-agent string        Which user-agent gork should use (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
  -v, --version                  version for gork
```

> If you're using docker, don't forget to mount your current directory with `/app` in the container to access your outfile

## example usage

```bash
> ./gork -t nmap.org
[+] Running gork on nmap.org
[+] done.

> cat gork.txt
        -== GORK RESULTS ==-

        --==== txt ===-
https://nmap.org/hobbit.ftpbounce.txt The FTP Bounce Attack - Nmap

        --==== xlsx ===-

        --==== csv ===-

        --==== bak ===-

        --==== log ===-
https://nmap.org/5/screenshots/sample-scan.log sample-scan.log - Nmap

        --==== docx ===-
https://nmap.org/oem/docs/Nmap-License-Contract.docx Nmap OEM Technology License Agreement

        --==== doc ===-

        --==== pdf ===-
https://nmap.org/book/toc.pdf Table of Contents - Nmap
https://nmap.org/nmapbook-toc.pdf Table of Contents - Nmap
https://nmap.org/docs/discovery.pdf nmap Host Discovery Techniques
https://nmap.org/misc/split-handshake.pdf The TCP Split Handshake: Practical Effects on Modern ... - Nmap
https://nmap.org/docs/nmap-mindmap.pdf nmap-mindmap.pdf
https://nmap.org/book/xlate/nmap-opensourcepress-de-cover.pdf nmap-opensourcepress-de-cover.pdf
https://nmap.org/book/cover/nns-cover.pdf Official Nmap Project Guide to Network Discovery and ...
https://nmap.org/presentations/BHDC08/bh-webcast-fyodor.pdf Scanning the Internet - Nmap
https://nmap.org/presentations/iSec08/isec08-slides-fyodor.pdf The New Nmap
https://nmap.org/presentations/BHDC08/bhdc08-slides-fyodor.pdf Scanning the Internet - Nmap

        --==== json ===-

```

