# bmake-shim: a workaround for `bmake` command missing from various (b)make packages

# EXAMPLE

```console
$ cd test
$ bmake
Works!
```

# DOWNLOAD

https://github.com/mcandre/bmake-shim/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/bmake-shim

# ABOUT

It may surprise developers that make is unreliable, with many projects implicitly relying on bmake syntax and semantics, whereas many environments default to GNU make instead. This tends to break the build. Developers interested in cross-platform support may wish to rewrite these files as `BSDmakefile`'s as a way to disambiguate, though unfortunately BSD make supports a different syntax and semantic set than bmake, requiring some effort to accomplish. One could then run (BSD) `bsdmake` to build a project.

While this would improve the build in GNU environments, ironically this breaks the build in many BSD environments! Evidently, package maintainers are also unreliable, often omitting the bsdmake command from their bsdmake packages. You're safe with macOS/Homebrew, but unsafe in other systems. Not to mention that neither bmake nor BSD make are (easily) available for Windows users compared to GNU make. Even if it were, it seems that there isn't a standard way to trigger BSD make-based builds, without sacrificing support for many platforms along the way. make just isn't that portable, unfortunately. How do we deal with this situation?

* Vendor-lock to BSD, bmake, and running `make`. No thanks.
* Wrap `make` calls in a dispatcher that invoke `make`, `bmake`, `bsdmake`, `gmake`, etc. according to the particular environment, overcomplicating everything. Non!
* Rewrite the build in terms of pure POSIX make syntax and semantics. This means no `.ifdef`, so this isn't a realistic option but for the very simplest of builds. Also, there aren't many ways to statically verify that a makefile avoids GNU or BSD features, so this is a difficult posture to maintain.
* Rewrite the build explicitly in terms of GNU make. Unfortunately, GNU make may be unavailable as the `gmake` command on GNU/Linux systems, so you would arrive at the same problem as bmake and BSD make!
* Rewrite the build in a tool with more standard behavior across different platforms, such as Mage, Tinyrick, Shake, Rake, Invoke, Gulp, Grunt, Gradle, Maven, CMake, Autotools, Makefile.PL, build.sh scripts, dedicated applications, etc. Long term, this is by far the most reliable solution. However, rewriting a project's build system can involve a lot of time and effort.

For projects that depend on bmake, the most immediately practical option for improving cross-platform support is to run build commands as `bmake` and shim environments that omit this command.

In a pinch, soft links, wrapper scripts, and shell aliases can serve, though these are limited to UNIX environments and the exact details will vary across environments, so we leave these to the particular system in question and encourage package maintainers to resolve this. For Windows, the situation is trickier, because there isn't even a standard shell, but an explosion of Command Prompt, PowerShell, and cygwin-like environments, each with different syntaxes and semantics, requiring a multitude of shims to guarantee uniform execution for different configurations.

Fortunately, we can cut this Gordian knot by simply providing binary shims, that reliably run regardless of shell particulars. Hence, bmake-shim.

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/gmake-shim/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/bmake-shim.git $GOPATH/src/github.com/mcandre/bmake-shim
$ cd $GOPATH/src/github.com/mcandre/bmake-shim
$ git submodule update --init --recursive
$ go install ./...
```

# UNIT TEST

```console
$ go test
```

# INTEGRATION TEST

```console
$ mage integrationTest
```

# UNIT + INTEGRATION TEST

```console
$ mage test
```

# LINT

```console
$ mage lint
```

# PORT

```console
$ mage port
```

# CLEAN ALL ARTIFACTS

```console
$ mage clean; mage uninstall; mage -clean
```

# SEE ALSO

https://github.com/mcandre/gmake-shim
