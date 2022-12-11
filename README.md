<div align="center">

# 🍃 mamuro-email

</div>

This is an application for indexing, uploading, and showing mail content on ZincSearch. Developed with Go, Chi and Vue/Tailwind. It also features:
 - Bash scripts for installation;
 - Bash scripts for indexing (without Go);
 - Supports WSL2;

## 🚀 Usage

### Download the mailing file and use the indexer

1. Make ZincSearch use the port 4080

```sh
#After unzipping
./zinc &
```

2. Execute the indexing binary

```sh
# Using the bash indexer
cd bin
./bash-indexer enron_mail_20110402

# Using the go indexer
cd bin
./indexer enron_mail_20110402

# You can use the go non-binary executable too
cd indexer
go run indexer.go enron_mail_20110402
```

3. Open Mamuro-email on desired port

```sh
cd chi-router
./mamuro -port 3000
```

4. Open [localhost:3000](http://localhost:3000/) and start searching!

