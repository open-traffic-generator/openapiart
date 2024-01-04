import click
import os

from .generate_from_yaml import GenerateFromYaml


@click.command()
@click.option(
    "--config_file",
    help="the config file for openapiart operations",
    required=True,
)
def generate(config_file):

    config_file = os.path.normpath(os.path.abspath(config_file))

    if not os.path.exists(config_file):
        raise Exception("the file %s does not exsist" % config_file)

    GenerateFromYaml(config_file)


if __name__ == "__main__":
    generate()
