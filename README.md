# Shred file function

`Shred(path)` function, implemented in Golang, that will overwrite the given 3 times with random data and delete the file afterwards.

## References

* <https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs>
* <https://www.practical-go-lessons.com/post/how-to-generate-random-bytes-with-golang-ccc9755gflds70ubqc2g>
* <https://blog.boot.dev/golang/golang-logging-best-practices/>
* <https://pkg.go.dev/os#File.Write>
* <https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package>
* <https://zetcode.com/golang/writefile/>
* <https://stackoverflow.com/questions/10516662/how-to-measure-test-coverage-in-go>
* <https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permissions>

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

## Run tests

Each time test is executed the coverage is printed.
It is also possible to generate an html coverage report as shown below.

```bash
make test
make cov_report
```
