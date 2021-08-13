import fnmatch
import os
import re
import sys
import shutil
import subprocess
import platform

global GOPATH
global GO_VERSION
global GO_TARGZ
global PROTOC_VERSION
global PROTOC_ZIP


def setup():
    run(
        [
            py() + " -m pip install --upgrade pip",
            py() + " -m pip install --upgrade virtualenv",
            py() + " -m virtualenv .env",
        ]
    )


def init():
    run(
        [
            py() + " -m pip install -r requirements.txt",
        ]
    )


def lint():
    paths = [pkg()[0], "tests", "setup.py", "do.py"]

    run(
        [
            py() + " -m black " + " ".join(paths),
            py() + " -m flake8 " + " ".join(paths),
            py() + " -m pytype " + " ".join(paths),
        ]
    )


def generate():
    artifacts = os.path.normpath(os.path.join(os.path.dirname(__file__), "openapiart", "tests", "artifacts.py"))
    run(
        [
            py() + " " + artifacts,
        ]
    )


def test():
    run(
        [
            py() + " -m pip install pytest-cov",
            py() + " -m pytest -sv --cov=sanity --cov-report term --cov-report html:cov_report",
        ]
    )
    import re

    coverage_threshold = 50
    with open("./cov_report/index.html") as fp:
        out = fp.read()
        result = re.findall(r"data-ratio.*?[>](\d+)\b", out)[0]
        if int(result) < coverage_threshold:
            raise Exception("Coverage thresold[{0}] is NOT achieved[{1}]".format(coverage_threshold, result))
        else:
            print("Coverage thresold[{0}] is achieved[{1}]".format(coverage_threshold, result))


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
            "{} -m pip install --upgrade --force-reinstall {}[testing]".format(py(), os.path.join("dist", wheel)),
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
        return py.path
    except AttributeError:
        py.path = os.path.join(".env", "bin", "python")
        if not os.path.exists(py.path):
            py.path = sys.executable

        # since some paths may contain spaces
        py.path = '"' + py.path + '"'
        return py.path


def setup_globals():
    global GOPATH
    global GO_VERSION
    global GO_TARGZ
    global PROTOC_VERSION
    global PROTOC_ZIP

    GOPATH = "./"
    GO_VERSION = "1.16.6"
    PROTOC_VERSION = "3.17.3"

    arch = platform.uname().machine.lower()
    if arch == 'arm64' or arch == "aarch64":
        print("Host architecture is ARM64")
        GO_TARGZ = f"go{GO_VERSION}.linux-arm64.tar.gz"
        PROTOC_ZIP=f"protoc-{PROTOC_VERSION}-linux-aarch_64.zip"
    elif arch == "x86_64":
        print("Host architecture is x86_64")
        GO_TARGZ = f"go{GO_VERSION}.linux-amd64.tar.gz"
        PROTOC_ZIP = f"protoc-{PROTOC_VERSION}-linux-x86_64.zip"
    else:
        print(f"Host architecture {arch} is not supported")


def get_go():
    global GO_TARGZ
    global GOPATH
    process = subprocess.run([f"curl -kL https://dl.google.com/go/{GO_TARGZ} | tar -C {GOPATH} -xzf -"], shell=True)
    if process.returncode == 0:
        print("Installed go successfully")
    else:
        print("installation unsucessful")
        print(process.stderr)
    return


def get_protoc():
    global PROTOC_VERSION
    global PROTOC_ZIP
    global GOPATH
    process_args = [
        "curl",
        "-kL",
        "-o",
        "./protc.zip",
        "https://github.com/protocolbuffers/protobuf/releases/download/v%s/%s" % (PROTOC_VERSION, PROTOC_ZIP)
    ]
    unzip = f"unzip -o ./protc.zip -d {GOPATH}/go bin/protoc \"include/*\""

    print("Executing", " ".join(process_args))
    process = subprocess.run([" ".join(process_args)], shell=True)
    if process.returncode == 0:
        print("installed go protoc")
        print("Executing", unzip)
        process = subprocess.run([unzip], shell=True)
        print("Executing", "rm -rf ./protc.zip")
        process = subprocess.run(["rm -rf ./protc.zip"], shell=True)
    else:
        print("installation unsucessful")
        print(process.stderr)
    return


def get_go_dependencies():
    deps = " ".join([
        "go get -v",
        "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        "google.golang.org/protobuf/cmd/protoc-gen-go",
        "golang.org/x/tools/cmd/goimports"
    ])
    print(f"installing dependencies {deps}")
    process = subprocess.run([deps], shell=True, capture_output=True)
    if process == 0:
        print("go dependencies installed sucessfully")
    else:
        print("installation unsucessful")
        print(process.stderr)
    return


def run(commands):
    """
    Executes a list of commands in a native shell and raises exception upon
    failure.
    """
    try:
        for cmd in commands:
            print(cmd)
            if sys.platform != "win32":
                cmd = cmd.encode("utf-8", errors="ignore")
            subprocess.check_call(cmd, shell=True)
    except Exception as e:
        print(e)
        sys.exit(1)


def main():
    if len(sys.argv) >= 2:
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        print("usage: python do.py [args]")


if __name__ == "__main__":
    setup_globals()
    main()
