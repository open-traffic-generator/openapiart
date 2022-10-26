import fnmatch
import os
import re
import sys
import shutil
import subprocess
import platform


def arch():
    return getattr(platform.uname(), "machine", platform.uname()[-1]).lower()


def on_arm():
    return arch() in ["arm64", "aarch64"]


def on_x86():
    return arch() == "x86_64"


def on_linux():
    print("The platform is {}".format(sys.platform))
    return "linux" in sys.platform


def linux_home_dir():
    return os.path.expanduser("~")


def linux_bin_exists(name):
    code, _ = subprocess.getstatusoutput(
        "which {} > /dev/null 2>&1".format(name)
    )

    return code == 0


def dot_profile_path():
    return os.path.join(linux_home_dir(), ".profile")


def dot_local_path():
    return os.path.join(linux_home_dir(), ".local")


def go_install_path():
    return "/usr/local/go/bin"


def go_bin_path():
    return os.path.join(linux_home_dir(), "go/bin")


def protoc_bin_path():
    return os.path.join(dot_local_path(), "bin")


def set_paths():
    if on_linux():
        os.environ["PATH"] = "{}:{}".format(
            os.environ["PATH"],
            ":".join([go_install_path(), go_bin_path(), protoc_bin_path()]),
        )


def py_env_installed():
    return ".env" in py()


def go_installer(version):
    if on_arm():
        return "go{}.linux-arm64.tar.gz".format(version)
    elif on_x86():
        return "go{}.linux-amd64.tar.gz".format(version)
    else:
        print("host architecture not supported")
        return None


def get_go(version="1.19"):
    if linux_bin_exists("go"):
        return

    installer = go_installer(version)
    if installer is None:
        return

    print("Installing Go ...")
    cmd = "curl -kL {} | sudo tar -C {} -xzf -".format(
        "https://dl.google.com/go/{}".format(installer), "/usr/local/"
    )
    run([cmd])

    with open(dot_profile_path(), "a") as f:
        f.write(
            "export PATH=$PATH:{}".format(
                ":".join([go_install_path(), go_bin_path()])
            )
        )


def get_go_ci_lint(version):
    run(
        [
            "curl -kLs {} | sh -s -- -b {} {}".format(
                "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh",
                go_bin_path(),
                version,
            )
        ]
    )


def get_go_deps(
    gen_go="v1.28.1", gen_grpc="v1.2.0", gen_doc="v1.5.1", ci_lint="v1.50.1"
):
    print("Installing Go dependencies for SDK generation ...")
    cmd = "CGO_ENABLED=0 go install -v {}@{}"
    run(
        [
            cmd.format("google.golang.org/protobuf/cmd/protoc-gen-go", gen_go),
            cmd.format(
                "google.golang.org/grpc/cmd/protoc-gen-go-grpc", gen_grpc
            ),
            cmd.format("golang.org/x/tools/cmd/goimports", "latest"),
            cmd.format(
                "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc",
                gen_doc,
            ),
        ]
    )
    get_go_ci_lint(ci_lint)


def protoc_installer(version):
    if on_arm():
        return "protoc-{}-linux-aarch_64.zip".format(version)
    elif on_x86():
        return "protoc-{}-linux-x86_64.zip".format(version)
    else:
        print("host architecture not supported")
        return None


def get_protoc(version="21.8"):
    if linux_bin_exists("protoc"):
        return

    installer = protoc_installer(version)
    if installer is None:
        return

    print("Installing protoc ...")
    cmd = "curl -kLO {0} && unzip {1} -d {2} && rm -rf {1}".format(
        "https://github.com/protocolbuffers/protobuf/releases/download/v{}/{}".format(
            version, installer
        ),
        installer,
        dot_local_path(),
    )
    run([cmd])

    with open(dot_profile_path(), "a") as f:
        f.write("export PATH=$PATH:{}".format(protoc_bin_path()))


def go_deps():
    print("Setting up pre-requisites for Go SDK")
    if on_linux():
        get_go()
        get_go_deps()
        get_protoc()
    else:
        print("Skipping go and protoc installation on non-linux platform ...")


def py_deps():
    print("Setting up python dependencies for SDK generation ...")
    run(
        [
            py() + " -m pip install -r requirements.txt",
            py() + " -m pip install -r test_requirements.txt",
        ]
    )


def setup():
    if not py_env_installed():
        print("Setting up python virtual environment ...")
        shutil.rmtree(".env", ignore_errors=True)
        run(
            [
                py() + " -m pip install --upgrade pip",
                py() + " -m pip install --upgrade virtualenv",
                py() + " -m virtualenv .env",
            ]
        )


def py_lint(modify="False"):
    paths = [pkg()[0], "setup.py", "do.py"]

    run(
        [
            py()
            + " -m black {} ".format(" ".join(paths))
            + "--exclude=openapiart/common.py {} --required-version {}".format(
                "" if modify == "True" else "--check", "22.1.0"
            )
        ]
    )

    run(
        [
            py() + " -m flake8 " + " ".join(paths),
        ]
    )


def generate(lang="all", import_from="source"):
    print(
        "Generating SDK for language {} and import path {}".format(
            lang, import_from
        )
    )
    if import_from == "package":
        import sys

        old_syspath = sys.path
        sys.path = [path for path in sys.path if pkg()[0] not in path]

        import openapiart

        sys.path = old_syspath
    else:
        import openapiart

    open_api = openapiart.OpenApiArt(
        api_files=[
            "openapiart/tests/api/info.yaml",
            "openapiart/tests/common/common.yaml",
            "openapiart/tests/api/api.yaml",
            "openapiart/goserver/api/service_a.api.yaml",
            "openapiart/goserver/api/service_b.api.yaml",
        ],
        artifact_dir="art",
        extension_prefix="sanity",
        proto_service="Openapi",
    )
    if lang == "all" or lang == "python":
        open_api.GeneratePythonSdk(package_name="sanity")

    if lang == "all" or lang == "go":
        open_api.GenerateGoSdk(
            package_dir="github.com/open-traffic-generator/openapiart/pkg",
            package_name="openapiart",
        )
        open_api.GenerateGoServer(
            module_path="github.com/open-traffic-generator/openapiart/pkg",
            models_prefix="openapiart",
            models_path="github.com/open-traffic-generator/openapiart/pkg",
        )
        open_api.GoTidy(
            relative_package_dir="pkg",
        )


def test_py_sdk():
    run(
        [
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


def test_go_sdk():
    try:
        min_cov = 35
        print("Running unit tests against Go SDK")
        os.chdir("pkg")

        cmd = "CGO_ENABLED=0 {}"

        out = run(
            [
                cmd.format("go mod tidy"),
                cmd.format("go test ./... -v -coverprofile coverage.txt"),
            ],
            capture_output=True,
        )

        result = re.findall(r"coverage:.*\s(\d+)", out)[0]
        if int(result) < min_cov:
            raise Exception(
                "Go tests achieved {1}% which is less than Coverage thresold {0}%,".format(
                    min_cov, result
                )
            )
        else:
            print(
                "Go tests achieved {1}% ,Coverage thresold {0}%".format(
                    min_cov, result
                )
            )
        if "FAIL" in out:
            raise Exception("Go Tests Failed")
    finally:
        os.chdir("..")


def go_lint():
    try:
        os.chdir("pkg")
        run(["CGO_ENABLED=0 golangci-lint run -v"])
    finally:
        os.chdir("..")


def test():
    generate(lang="all")
    py_lint(modify="True")
    go_lint()
    test_py_sdk()
    test_go_sdk()


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
        return py.path
    except AttributeError:
        py.path = os.path.join(".env", "bin", "python")
        if not os.path.exists(py.path):
            py.path = sys.executable

        # since some paths may contain spaces
        py.path = '"' + py.path + '"'
        print("Using python executable ", py.path)
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
            print(cmd)
            subprocess.check_call(cmd, shell=True, stdout=fd)
        return flush_output(fd, logfile)
    except Exception:
        flush_output(fd, logfile)
        sys.exit(1)


def main():
    if len(sys.argv) >= 2:
        set_paths()
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        print("usage: python do.py [args]")


if __name__ == "__main__":
    main()
