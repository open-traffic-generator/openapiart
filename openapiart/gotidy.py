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
