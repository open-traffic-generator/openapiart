import os
import sys
import subprocess

def generate_requirements(path, save_path=None, ignore_path=None, file_name=None):
    '''
        To generate the requirements.txt of library in its path
    '''
    if save_path is None:
        file_name = "new_requirements.txt"
        save_path = path

    new_save_path = os.path.join(save_path, file_name)

    if sys.version_info[0] < 3:
        run([py() + " -m pip install pipreqs== 0.4.8"])
    else:
        run([py() + " -m pip install pipreqs"])
    
    run(
        [
            py() + " -m pipreqs.pipreqs --force " + path + " --mode no-pin --ignore " + str(ignore_path) + " --savepath " + new_save_path
        ]
    )

    if file_name == "test_requirements.txt":
        generate_test_requirements(save_path)
        return

    not_required_pkgs = ['grpc', 'grpcio', 'grpcio-tools', 'protobuf']

    version_restrict = ["grpcio==1.38.0 ; python_version > '2.7'", 
                        "grpcio-tools==1.38.0 ; python_version > '2.7'",
                        "grpcio==1.35.0 ; python_version == '2.7'",
                        "grpcio-tools==1.35.0 ; python_version == '2.7'",
                        "protobuf==3.15.0"
                       ]

    with open(os.path.join(save_path, file_name)) as f:
        new_packages = f.read().splitlines()
        if 'grpc' in new_packages:
            new_packages = list(set(new_packages) - set(not_required_pkgs))
        new_packages.extend(version_restrict)
    
    if file_name == "new_requirements.txt":
        os.remove(os.path.join(save_path, "new_requirements.txt"))

    if os.path.exists(os.path.join(save_path, 'requirements.txt')):
        with open(os.path.join(save_path, 'requirements.txt'), 'r')as fp:
            packages = fp.read().splitlines()
        new_packages = list(set(new_packages) - set(packages))

    if new_packages:
        with open(os.path.join(save_path, 'requirements.txt'), 'a+') as fh:
            fh.write("\n")
            for pkg in new_packages:
                fh.write(pkg + "\n")
            fh.flush()
            fh.close()


def generate_test_requirements(save_path):
    '''
        To generate the requirememts.txt for the test lib present in the path.
    '''
    not_required_pkgs = ['grpc', 'grpcio', 'grpcio-tools', 'protobuf']

    if os.path.exists(os.path.join(save_path, 'requirements.txt')):
        with open(os.path.join(save_path, 'requirements.txt'), 'r')as fp:
            packages = fp.read().splitlines()

    with open(os.path.join(save_path, "test_requirements.txt")) as fp:
        test_packages = fp.read().splitlines()
        if 'grpc' in test_packages:
            test_packages = list(set(test_packages) - set(not_required_pkgs))
    
    diff_packages = list(set(test_packages) - set(packages))
    with open(os.path.join(save_path, "test_requirements.txt"), 'w+') as fp:
        for pkg in diff_packages:
            fp.write(pkg + "\n")


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