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

## available dorks

- database files
- documents (docx, pdf, pptx, etc...)
- project files (xml, ini, license, install, etc...)
- directory listing
- backup files
- config files

## usage

```bash
gork is a CLI to perform Google dorks on a target domain :)~ (Example: ./gork -t nmap.org)

Usage:
  gork [flags]

Flags:
  -a, --append-results           append dork results to out file
  -x, --exclude stringArray      exclude specific filetype (default [html])
  -e, --extensions stringArray   filetype extensions (default [asp,aspx,backup,bak,bkf,bkp,cfg,cgi,cnf,conf,csv,dbf,doc,docx,fla,inf,ini,json,jsp,jspx,log,mdb,odt,old,ora,pdf,php,ppt,pptx,rdp,reg,rtf,sql,sxt,txt,xlsx,xml])
  -h, --help                     help for gork
  -o, --outfile string           directory storing dorks results (default "./gork.txt")
  -p, --proxy string             proxy URL
  -t, --target string            target site for your dorks
  -u, --user-agent string        Which user-agent gork should use (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
  -v, --version                  version for gork
```

> If you're using docker, don't forget to mount your current directory with `/app` in the container to access your outfile

## example

```bash
> ./gork -t nmap.org
[+] running gork on nmap.org
        found result(s) for dir listing
        found result(s) for project setup files
        found result(s) for pdf
[+] scan completed

> cat gork.txt
        -== GORK RESULTS FOR nmap.org ==-

                --==== dir listing ===-
http://scanme.nmap.org/shared/templates/ Index of /shared/templates
http://scanme.nmap.org/shared/css/ Index of /shared/css
http://scanme.nmap.org/shared/images/Acunetix/ Index of /shared/images/Acunetix - Nmap ScanMe
http://scanme.nmap.org/images/ Index of /images - Nmap ScanMe
http://scanme.nmap.org/shared/error/includes/ Index of /shared/error/includes
http://scanme.nmap.org/shared/error/ Index of /shared/error - Nmap ScanMe

                --==== project setup files ===-
https://nmap.org/oem/docs/Nmap-License-Contract.pdf Nmap OEM Technology License Agreement
https://nmap.org/npcap/oem/docs/Npcap-OEM-Internal-Use-License.pdf Npcap OEM Internal Use End User License Agreement
https://nmap.org/npcap/oem/docs/Npcap-OEM-Redistribution-License-Term.pdf npcap oem technology license agreement - Nmap
https://svn.nmap.org/nmap/mswin32/license-format/ - Revision 38450: /nmap/mswin32/license-format

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
```
