# Shred file function

`Shred(path)` function, implemented in Golang, that will overwrite the given 3 times with random data and delete the file afterwards.

## References

* <https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs>
* <https://www.practical-go-lessons.com/post/how-to-generate-random-bytes-with-golang-ccc9755gflds70ubqc2g>
* <https://blog.boot.dev/golang/golang-logging-best-practices/>
* <https://pkg.go.dev/os#File.Write>

## Run go script

Clone the repository and run the following command in project folder, we expect go properly installed.

```bash
echo "Sample content" > demo.txt
go run main.go shred.go demo.txt
```

## Build shred program

Once cloned the repository, in the project folder execute

```bash
make build
```

a `bin/shred` executable file will be generated. You clean via `make clean`.
