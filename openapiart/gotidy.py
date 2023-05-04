import os
import subprocess


class GoTidy(object):
    def __init__(
        self,
        output_root_path,
    ):
        self._output_root_path = output_root_path

    def goTidy(self):
        print(
            "GoTidy output directory: {path}".format(
                path=self._output_root_path
            )
        )
        self._format_go()
        self._replace_specific_versions()
        self._tidy_mod()

    def _format_go(self):
        """Format the generated go code"""
        try:
            process_args = [
                "goimports",
                "-w",
                self._output_root_path,
            ]
            cmd = " ".join(process_args)
            print("Formatting generated go files in folder: {}".format(cmd))
            process = subprocess.Popen(
                cmd, cwd=self._output_root_path, shell=True
            )
            process.wait()
        except Exception as e:
            print("Bypassed formatting of generated go ux file: {}".format(e))

    def _replace_specific_versions(self):
        imports = []
        with open(os.path.join(self._output_root_path, "go.mod"), "r") as fh:
            imports = fh.read().splitlines()

        new_imports = []
        for pkg in imports:
            if "go.opentelemetry.io/contrib" in pkg:
                module = pkg.split(" ")[0]
                new_imports.append(module + " v0.37.0")
            elif "go.opentelemetry.io/otel" in pkg:
                module = pkg.split(" ")[0]
                new_imports.append(module + " v1.14.0")
            else:
                new_imports.append(pkg)

        with open(os.path.join(self._output_root_path, "go.mod"), "w") as fh:
            for npkg in new_imports:
                fh.write(npkg + "\n")
            fh.flush()
            fh.close()

    def _tidy_mod(self):
        """Tidy the mod file"""
        try:
            process_args = [
                "go",
                "mod",
                "tidy",
            ]
            os.environ["GO111MODULE"] = "on"
            print(
                "Tidying the generated go mod file: {}".format(
                    " ".join(process_args)
                )
            )
            process = subprocess.Popen(
                process_args,
                cwd=self._output_root_path,
                shell=False,
                env=os.environ,
            )
            process.wait()
        except Exception as e:
            print("Bypassed tidying the generated mod file: {}".format(e))
