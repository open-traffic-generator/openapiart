import os
import sys
import subprocess
import re

base_dir = os.path.dirname(os.path.abspath(__file__))


def _package_name(requirement):
    match = re.match(r"\s*([A-Za-z0-9_.-]+)", requirement.split(";", 1)[0])
    if match is None:
        return None
    return match.group(1).replace("_", "-").lower()


def _without_packages(packages, ignored_packages):
    ignored = {
        name
        for name in (_package_name(pkg) for pkg in ignored_packages)
        if name is not None
    }
    return [pkg for pkg in packages if _package_name(pkg) not in ignored]


def _resolve_required_packages(
    new_packages, orig_packages, test_packages, ignored_packages=None
):
    if ignored_packages is None:
        ignored_packages = []

    new_packages = _without_packages(new_packages, ignored_packages)
    orig_packages = _without_packages(orig_packages, ignored_packages)
    test_packages = _without_packages(test_packages, ignored_packages)

    required_names = {
        name
        for name in (_package_name(pkg) for pkg in new_packages)
        if name is not None
    }

    final_packages = []
    for pkg in orig_packages + test_packages:
        name = _package_name(pkg)
        if name in required_names and pkg not in final_packages:
            final_packages.append(pkg)
    return final_packages


def generate_requirements(path, file_name=None):
    """
    To generate the requirements.txt of library in its path
    """
    file_name = "requirements.txt"
    save_path = path

    new_save_path = os.path.join(save_path, file_name)

    process_args = [
        "{} -m pipreqs.pipreqs --force {}".format(sys.executable, path),
        "--mode no-pin",
        "--savepath {}".format(new_save_path),
    ]

    subprocess.check_call(" ".join(process_args), shell=True)

    not_required_pkgs = [
        "sanity",
        "typing_extensions",
    ]

    with open(os.path.join(base_dir, "requirements.txt"), "r") as fd:
        orig_packages = fd.read().splitlines()

    with open(os.path.join(save_path, "requirements.txt"), "r") as fh:
        new_pkgs = fh.read().splitlines()

    with open(os.path.join(base_dir, "test_requirements.txt"), "r") as fh:
        test_pkgs = fh.read().splitlines()

    final_pkgs = _resolve_required_packages(
        new_pkgs, orig_packages, test_pkgs, not_required_pkgs
    )

    with open(os.path.join(save_path, "requirements.txt"), "w+") as fh:
        fh.write("--prefer-binary")
        fh.write("\n")
        for pkg in final_pkgs:
            fh.write(pkg + "\n")
        fh.flush()
        fh.close()
