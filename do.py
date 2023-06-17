import fnmatch
import os
import re
import sys
import shutil
import subprocess
import platform


BLACK_VERSION = "23.3.0"
GO_VERSION = "1.20"
PROTOC_VERSION = "3.20.3"

# this is where go and protoc shall be installed (and expected to be present)
LOCAL_PATH = os.path.join(os.path.expanduser("~"), ".local")
# path where protoc bin shall be installed or expected to be present
LOCAL_BIN_PATH = os.path.join(LOCAL_PATH, "bin")
# path where go bin shall be installed or expected to be present
GO_BIN_PATH = os.path.join(LOCAL_PATH, "go", "bin")
# path for go package source and installations
GO_HOME_PATH = os.path.join(os.path.expanduser("~"), "go")
GO_HOME_BIN_PATH = os.path.join(GO_HOME_PATH, "bin")

os.environ["GOPATH"] = GO_HOME_PATH
os.environ["PATH"] = "{}:{}:{}:{}".format(
    os.environ["PATH"], GO_BIN_PATH, GO_HOME_BIN_PATH, LOCAL_BIN_PATH
)


def arch():
    return getattr(platform.uname(), "machine", platform.uname()[-1]).lower()


def on_arm():
    return arch() in ["arm64", "aarch64"]


def on_x86():
    return arch() == "x86_64"


def on_linux():
    print("The platform is {}".format(sys.platform))
    return "linux" in sys.platform


def get_go(version=GO_VERSION, targz=None):
    if targz is None:
        if on_arm():
            targz = "go" + version + ".linux-arm64.tar.gz"
        elif on_x86():
            targz = "go" + version + ".linux-amd64.tar.gz"
        else:
            print("host architecture not supported")
            return

    print("Installing Go ...")

    if not os.path.exists(LOCAL_PATH):
        os.mkdir(LOCAL_PATH)

    cmd = "go version 2> /dev/null"
    cmd += " || (rm -rf $(dirname {})".format(GO_BIN_PATH)
    cmd += " && curl -kL -o go-installer https://dl.google.com/go/{}".format(
        targz
    )
    cmd += " && tar -C {} -xzf go-installer".format(LOCAL_PATH)
    cmd += " && rm -rf go-installer"
    cmd += " && echo 'PATH=$PATH:{}:{}' >> ~/.profile".format(
        GO_BIN_PATH, GO_HOME_BIN_PATH
    )
    cmd += " && echo 'export GOPATH={}' >> ~/.profile)".format(GO_HOME_PATH)
    run([cmd])


def get_go_deps():
    print("Getting Go libraries for grpc / protobuf ...")
    cmd = "GO111MODULE=on CGO_ENABLED=0 go install"
    run(
        [
            cmd + " -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0",
            cmd + " -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1",
            cmd + " -v golang.org/x/tools/cmd/goimports@v0.6.0",
            cmd
            + " -v github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1",
        ]
    )


def get_protoc(version=PROTOC_VERSION, zipfile=None):
    if zipfile is None:
        if on_arm():
            zipfile = "protoc-" + version + "-linux-aarch_64.zip"
        elif on_x86():
            zipfile = "protoc-" + version + "-linux-x86_64.zip"
        else:
            print("host architecture not supported")
            return

    print("Installing protoc ...")

    if not os.path.exists(LOCAL_PATH):
        os.mkdir(LOCAL_PATH)

    cmd = "protoc --version 2> /dev/null || (curl -kL -o ./protoc.zip "
    cmd += "https://github.com/protocolbuffers/protobuf/releases/download/v{}/{}".format(
        version, zipfile
    )
    cmd += " && unzip -o ./protoc.zip -d {}".format(LOCAL_PATH)
    cmd += " && rm -rf ./protoc.zip"
    cmd += " && echo 'PATH=$PATH:{}' >> ~/.profile)".format(LOCAL_BIN_PATH)
    run([cmd])


def setup_ext(go_version=GO_VERSION, protoc_version=PROTOC_VERSION):
    if on_linux():
        get_go(go_version)
        get_protoc(protoc_version)
        get_go_deps()
    else:
        print("Skipping go and protoc installation on non-linux platform ...")


def setup():
    if platform.python_version_tuple()[0] == 3:
        run(
            [
                py() + " -m pip install --upgrade pip",
                py() + " -m {} .env".format(pkg),
            ]
        )
    else:
        run(
            [
                py() + " -m pip install --upgrade pip",
                py() + " -m pip install --upgrade virtualenv",
                py() + " -m virtualenv .env",
            ]
        )


def init(use_sdk=None):
    base_dir = os.path.dirname(os.path.abspath(__file__))
    if use_sdk is None:
        req = os.path.join(base_dir, "openapiart", "requirements.txt")
        test_req = os.path.join(
            base_dir, "openapiart", "test_requirements.txt"
        )
        run(
            [
                py() + " -m pip install -r {}".format(req),
                py() + " -m pip install -r {}".format(test_req),
            ]
        )
    else:
        art_path = os.path.join(base_dir, "art", "requirements.txt")
        art_test = os.path.join(base_dir, "art", "test_requirements.txt")
        run(
            [
                py() + " -m pip install -r {}".format(art_path),
                py() + " -m pip install -r {}".format(art_test),
            ]
        )


def lint(check="false"):
    paths = [
        pkg()[0],
        "openapiart",
        "setup.py",
        "do.py",
    ]
    # --check will check for any files to be formatted with black
    # if linting fails, format the files with black and commit
    cmd = " --exclude=openapiart/common.py"
    if check.lower() == "true":
        cmd += " --check"
    cmd += " --required-version {}".format(BLACK_VERSION)
    ret, out = getstatusoutput(py() + " -m black " + " ".join(paths) + cmd)
    if ret == 1:
        raise Exception(
            "Black formatting failed, with black version {}.\n{}".format(
                BLACK_VERSION, out
            )
        )
    else:
        print(out)
    run(
        [
            py() + " -m flake8 " + " ".join(paths),
        ]
    )


def generate(sdk="", cicd=""):
    artifacts = os.path.normpath(
        os.path.join(os.path.dirname(__file__), "artifacts.py")
    )
    run(
        [
            py() + " " + artifacts + " " + sdk + " " + cicd,
        ]
    )


def testpy():
    run(
        [
            # py() + " -m pip install flask",
            # py() + " -m pip install pytest-cov",
            py()
            + " -m pytest -sv --cov=sanity --cov-report term --cov-report html:cov_report",
        ]
    )
    import re

    coverage_threshold = 45
    with open("./cov_report/index.html") as fp:
        out = fp.read()
        result = re.findall(r"data-ratio.*?[>](\d+)\b", out)[0]
        if int(result) < coverage_threshold:
            raise Exception(
                "Coverage thresold[{0}] is NOT achieved[{1}]".format(
                    coverage_threshold, result
                )
            )
        else:
            print(
                "Coverage thresold[{0}] is achieved[{1}]".format(
                    coverage_threshold, result
                )
            )


def testgo():
    go_coverage_threshold = 35
    # TODO: not able to run the test from main directory
    os.chdir("pkg")
    run(["go mod tidy"], capture_output=True)
    ret = run(
        ["go test ./... -v -coverprofile coverage.txt"], capture_output=True
    )
    os.chdir("..")
    result = re.findall(r"coverage:.*\s(\d+)", ret)[0]
    if int(result) < go_coverage_threshold:
        raise Exception(
            "Go tests achieved {1}% which is less than Coverage thresold {0}%,".format(
                go_coverage_threshold, result
            )
        )
    else:
        print(
            "Go tests achieved {1}% ,Coverage thresold {0}%".format(
                go_coverage_threshold, result
            )
        )
    if "FAIL" in ret:
        raise Exception("Go Tests Failed")


def go_lint():
    try:
        output = run(["go version"], capture_output=True)
        if "go1.17" in output or "go1.18" in output:
            print("Using older linter version for go version older than 1.19")
            version = "1.46.2"
        else:
            version = "1.51.1"

        pkg = "{}go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v{}".format(
            "" if sys.platform == "win32" else "GO111MODULE=on CGO_ENABLED=0 ",
            version,
        )
        run([pkg])
        os.chdir("pkg")
        run(["golangci-lint run -v"])
    finally:
        os.chdir("..")


def dist():
    clean()
    run(
        [
            py() + " setup.py sdist bdist_wheel --universal",
        ]
    )
    print(os.listdir("dist"))


def install():
    wheel = "{}-{}-py2.py3-none-any.whl".format(*pkg())
    run(
        [
            "{} -m pip install --force-reinstall --no-cache-dir {}[testing]".format(
                py(), os.path.join("dist", wheel)
            ),
        ]
    )


def install_package_only():
    wheel = "{}-{}-py2.py3-none-any.whl".format(*pkg())
    run(
        [
            "{} -m pip install --force-reinstall --no-cache-dir {}".format(
                py(), os.path.join("dist", wheel)
            ),
        ]
    )


def release():
    run(
        [
            py() + " -m pip install --upgrade twine",
            "{} -m twine upload -u {} -p {} dist/*".format(
                py(),
                os.environ["PYPI_USERNAME"],
                os.environ["PYPI_PASSWORD"],
            ),
        ]
    )


def clean():
    """
    Removes filenames or dirnames matching provided patterns.
    """
    pwd_patterns = [
        ".pytype",
        "dist",
        "build",
        "*.egg-info",
    ]
    recursive_patterns = [
        ".pytest_cache",
        "__pycache__",
        "*.pyc",
        "*.log",
    ]

    for pattern in pwd_patterns:
        for path in pattern_find(".", pattern, recursive=False):
            rm_path(path)

    for pattern in recursive_patterns:
        for path in pattern_find(".", pattern, recursive=True):
            rm_path(path)


def version():
    print(pkg()[-1])


def pkg():
    """
    Returns name of python package in current directory and its version.
    """
    try:
        return pkg.pkg
    except AttributeError:
        with open("setup.py") as f:
            out = f.read()
            name = re.findall(r"pkg_name = \"(.+)\"", out)[0]
            version = re.findall(r"version = \"(.+)\"", out)[0]

            pkg.pkg = (name, version)
        return pkg.pkg


def rm_path(path):
    """
    Removes a path if it exists.
    """
    if os.path.exists(path):
        if os.path.isdir(path):
            shutil.rmtree(path)
        else:
            os.remove(path)


def pattern_find(src, pattern, recursive=True):
    """
    Recursively searches for a dirname or filename matching given pattern and
    returns all the matches.
    """
    matches = []

    if not recursive:
        for name in os.listdir(src):
            if fnmatch.fnmatch(name, pattern):
                matches.append(os.path.join(src, name))
        return matches

    for dirpath, dirnames, filenames in os.walk(src):
        for names in [dirnames, filenames]:
            for name in names:
                if fnmatch.fnmatch(name, pattern):
                    matches.append(os.path.join(dirpath, name))

    return matches


def py():
    """
    Returns path to python executable to be used.
    """
    try:
        print(py.path)
        return py.path
    except AttributeError:
        py.path = os.path.join(".env", "bin", "python")
        if not os.path.exists(py.path):
            py.path = sys.executable

        # since some paths may contain spaces
        py.path = '"' + py.path + '"'
        print(py.path)
        return py.path


def flush_output(fd, filename):
    """
    Flush the log file and print to console
    """
    if fd is None:
        return
    fd.flush()
    fd.seek(0)
    ret = fd.read()
    print(ret)
    fd.close()
    os.remove(filename)
    return ret


def run(commands, capture_output=False):
    """
    Executes a list of commands in a native shell and raises exception upon
    failure.
    """
    fd = None
    logfile = "log.txt"
    if capture_output:
        fd = open(logfile, "w+")
    try:
        for cmd in commands:
            if sys.platform != "win32":
                cmd = cmd.encode("utf-8", errors="ignore")
            subprocess.check_call(cmd, shell=True, stdout=fd)
        return flush_output(fd, logfile)
    except Exception:
        flush_output(fd, logfile)
        sys.exit(1)


def getstatusoutput(command):
    return (
        subprocess.getstatusoutput(command)[0],
        subprocess.getstatusoutput(command)[1],
    )


def build(sdk="all", env_setup=None):
    print("\nStep 1: Set up virtaul env")

    if env_setup is not None and env_setup.lower() == "clean":
        print("\nCleaning up exsisting env")
        run(["rm -rf .env"])

    if not os.path.exists(".env"):
        setup()
    else:
        print("\nvirtualenv already exists.\n")

    py.path = os.path.join(".env", "bin", "python")
    print(
        "\nWill be using the following python interpreter path "
        + py.path
        + "\n"
    )

    print("\nStep 2: Install current changes of openapiart to venv\n")
    base_dir = os.path.dirname(os.path.abspath(__file__))
    test_req = os.path.join(base_dir, "openapiart", "test_requirements.txt")
    run([py() + " setup.py install", py() + " -m pip install -r " + test_req])
    print("\nStep 3: Generating python and Go SDKs\n")
    generate(sdk=sdk, cicd="True")
    if sdk == "python" or sdk == "all":
        print("\nStep 4: Perform Python lint\n")
        lint()
        print("\nStep 5: Run Python Tests\n")
        testpy()
    else:
        print("\nSkipping Step 4: python lint and Step 5: run python tests\n")
    if sdk == "go" or sdk == "all":
        print("\nStep 6: Run Go Lint\n")
        go_lint()
        print("\nStep 7: Run Go Tests")
        testgo()
    else:
        print("\nStep 6: Perform Go lint and Step 7: run go tests\n")
    print("\nBuild Successfull\n")


def main():
    if len(sys.argv) >= 2:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        print("usage: python do.py [args]")


if __name__ == "__main__":
    main()
