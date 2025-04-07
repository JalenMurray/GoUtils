
# CSV Utilities

This package includes some potentially useful utility functions for CSV files


## Run Locally

Clone the project

```bash
  git clone https://github.com/JalenMurray/GoUtils
```

Go to the project directory

```bash
  cd GoUtils/CSV
```

Build Executable
```bash
  go build
```

Run Executable
```
  ./csv-utils.exe
```



## Usage
```bash
Usage:
    ./csv-utils.exe <command> [args]
Available Commands:
    get-headers <file1> <file2> ... <fileN> Show headers of provided CSV files
    get-rows <file1> <file2> ... <fileN> Show number of rows in provided CSV files
    random-row <file1> <file2> ... <fileN> Show a random row from provided CSV files
```

## Example
```bash
  ./csv-utils get-headers test.csv test2.csv
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

