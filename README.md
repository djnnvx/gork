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
  -e, --extensions stringArray   filetype extensions (default [doc,docx,csv,pdf,txt,log,bak,json,xlsx,xml,conf,cnf,reg,inf,rdp,cfg,ora,ini,sql,mdb,dbf,bkf,bkp,old,backup,rtf,odt,ppt,sxt,pptx])
  -h, --help                     help for gork
  -o, --outfile string           directory storing dorks results (default "./gork.txt")
  -p, --proxy string             proxy URL
  -t, --target string            target site for your dorks
  -u, --user-agent string        Which user-agent gork should use (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
  -v, --version
```

> If you're using docker, don't forget to mount your current directory with `/app` in the container to access your outfile

## example

```bash
> ./gork -t nmap.org
[+] running gork on nmap.org
        found result(s) for open redirects
        found result(s) for dir listing
        found result(s) for project setup files
        found result(s) for pdf
[+] scan completed

> cat gork.txt
        -== GORK RESULTS FOR nmap.org ==-

                --==== open redirects ===-
https://nmap.org/nsedoc/scripts/http-open-redirect.html http-open-redirect NSE script - Nmap
https://nmap.org/nsedoc/scripts/https-redirect.html https-redirect NSE script - Nmap
https://nmap.org/nsedoc/scripts/rtsp-url-brute.html rtsp-url-brute NSE script - Nmap
https://nmap.org/nsedoc/scripts/url-snarf.html url-snarf NSE script — Nmap Scripting Engine documentation
https://nmap.org/nsedoc/lib/url.html url NSE Library — Nmap Scripting Engine documentation

                --==== dir listing ===-
http://scanme.nmap.org/shared/templates/ Index of /shared/templates
http://scanme.nmap.org/shared/css/ Index of /shared/css
http://scanme.nmap.org/shared/images/Acunetix/ Index of /shared/images/Acunetix - Nmap ScanMe
http://scanme.nmap.org/images/ Index of /images - Nmap ScanMe
http://scanme.nmap.org/shared/error/includes/ Index of /shared/error/includes
http://scanme.nmap.org/shared/error/ Index of /shared/error - Nmap ScanMe
https://svn.nmap.org/nmap-mswin32-aux/Python/Lib/site-packages/gtk-2.0/runtime/share/gtk-doc/html/glib/api-index-2-20.html Index of new symbols in 2.20
https://svn.nmap.org/nmap-mswin32-aux/Python/Lib/site-packages/gtk-2.0/runtime/share/gtk-doc/html/gio/api-index-2-28.html Index of new symbols in 2.28
https://svn.nmap.org/nmap-mswin32-aux/Python/Lib/site-packages/gtk-2.0/runtime/share/gtk-doc/html/glib/api-index-2-18.html Index of new symbols in 2.18

                --==== project setup files ===-
https://nmap.org/book/install.html Chapter 2. Obtaining, Compiling, Installing, and ... - Nmap
https://nmap.org/nsedoc/scripts/http-config-backup.html http-config-backup NSE script - Nmap
https://nmap.org/nsedoc/scripts/snmp-ios-config.html snmp-ios-config NSE script - Nmap
https://nmap.org/nsedoc/scripts/ms-sql-config.html ms-sql-config NSE script - Nmap
https://nmap.org/oem/docs/Nmap-License-Contract.pdf Nmap OEM Technology License Agreement
https://nmap.org/npcap/oem/docs/Npcap-OEM-Internal-Use-License.pdf Npcap OEM Internal Use End User License Agreement
https://nmap.org/npcap/oem/docs/Npcap-OEM-Redistribution-License-Term.pdf npcap oem technology license agreement - Nmap
https://nmap.org/nsedoc/scripts/lexmark-config.html lexmark-config NSE script - Nmap
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

