import os
import sys
import subprocess

base_dir = os.path.dirname(os.path.abspath(__file__))


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
        new_pkgs = list(set(new_pkgs) - set(not_required_pkgs))

    final_pkgs = []
    for n_pkg in new_pkgs:
        for pkg in orig_packages:
            if n_pkg in pkg and pkg not in final_pkgs:
                final_pkgs.append(pkg)

    with open(os.path.join(save_path, "requirements.txt"), "w+") as fh:
        fh.write("--prefer-binary")
        fh.write("\n")
        for pkg in final_pkgs:
            fh.write(pkg + "\n")
        fh.flush()
        fh.close()
