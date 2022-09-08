import fnmatch
import os
import re
import sys
import shutil
import subprocess
import platform

# from openapiart.generate_requirements import generate_requirements


BLACK_VERSION = "22.1.0"

os.environ["GOPATH"] = os.path.join(
    os.path.dirname(os.path.abspath(__file__)), ".local"
)
os.environ["PATH"] = os.environ["PATH"] + ":{0}/go/bin:{0}/bin".format(
    os.environ["GOPATH"]
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


def get_go(go_version="1.17"):
    version = go_version
    targz = None

    if on_arm():
        targz = "go" + version + ".linux-arm64.tar.gz"
    elif on_x86():
        targz = "go" + version + ".linux-amd64.tar.gz"
    else:
        print("host architecture not supported")
        return

    if not os.path.exists(os.environ["GOPATH"]):
        os.mkdir(os.environ["GOPATH"])

    print("Installing Go ...")
    cmd = (
        "go version 2> /dev/null || curl -kL https://dl.google.com/go/" + targz
    )
    cmd += " | tar -C " + os.environ["GOPATH"] + " -xzf -"
    run([cmd])


def get_go_deps():
    print("Getting Go libraries for grpc / protobuf ...")
    cmd = "GO111MODULE=on CGO_ENABLED=0 go install"
    run(
        [
            cmd + " -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0",
            cmd + " -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0",
            cmd + " -v golang.org/x/tools/cmd/goimports@latest",
            cmd
            + " -v github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest",
        ]
    )


def get_protoc():
    version = "3.17.3"
    zipfile = None

    if on_arm():
        zipfile = "protoc-" + version + "-linux-aarch_64.zip"
    elif on_x86():
        zipfile = "protoc-" + version + "-linux-x86_64.zip"
    else:
        print("host architecture not supported")
        return

    print("Installing protoc ...")
    cmd = "protoc --version 2> /dev/null || ( curl -kL -o ./protoc.zip "
    cmd += "https://github.com/protocolbuffers/protobuf/releases/download/v"
    cmd += version + "/" + zipfile
    cmd += " && unzip ./protoc.zip -d " + os.environ["GOPATH"]
    cmd += " && rm -rf ./protoc.zip )"
    run([cmd])


def setup_ext(go_version="1.17"):
    if on_linux():
        get_go(go_version)
        get_protoc()
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
        run(
            [
                py() + " -m pip install -r {}".format(req),
                py() + " -m pip install -r test_requirements.txt",
            ]
        )
    else:
        art_path = os.path.join(base_dir, "art", "requirements.txt")
        run(
            [
                py() + " -m pip install -r {}".format(art_path),
                py() + " -m pip install -r test_requirements.txt",
            ]
        )


def lint():
    paths = [
        pkg()[0],
        "openapiart",
        "setup.py",
        "do.py",
    ]
    # --check will check for any files to be formatted with black
    # if linting fails, format the files with black and commit
    ret, out = getstatusoutput(
        py()
        + " -m black "
        + " ".join(paths)
        + " --exclude=openapiart/common.py --check --required-version {}".format(
            BLACK_VERSION
        )
    )
    if ret == 1:
        out = out.split("\n")
        out = [
            e.replace("would reformat", "Black formatting failed for")
            for e in out
            if "would reformat" in e
        ]
        print("\n".join(out))
        raise Exception(
            "Black formatting failed, use black {} version to format the files".format(
                BLACK_VERSION
            )
        )
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
    pkg = "{env_set}go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2".format(
        env_set=""
        if sys.platform == "win32"
        else "GO111MODULE=on CGO_ENABLED=0 "
    )
    run([pkg])
    os.chdir("pkg")
    run(["golangci-lint run -v"])


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


def main():
    if len(sys.argv) >= 2:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        print("usage: python do.py [args]")


if __name__ == "__main__":
    main()
