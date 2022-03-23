"""
To build distribution: python setup.py sdist --formats=gztar bdist_wheel --universal
"""
import imp
import os
import sys
import setuptools
from generate_requirements import *

pkg_name = "openapiart"
version = "0.1.42"

base_dir = os.path.dirname(os.path.abspath(__file__))
with open(os.path.join(base_dir, "README.md")) as fid:
    long_description = fid.read()

if not os.path.exists(os.path.join(base_dir, 'requirements.txt')):
    base_path = os.getcwd()
    openapiart_path = os.path.join(base_path, 'openapiart')
    test_path = os.path.join(openapiart_path, "tests")
    generate_requirements(openapiart_path, ignore_path=test_path, save_path="new_requirements.txt")
    generate_requirements(test_path,save_path="test_requirements")

with open("requirements.txt") as f:
    installation_requires = f.read().splitlines()
    if '--prefer-binary' in installation_requires:
        installation_requires.remove('--prefer-binary')

installation_requires.append("black==22.1.0 ; python_version > '2.7'")

test_pkgs = ["flake8", "black"]
with open("test_requirements.txt") as f:
    test_requires = f.read().splitlines()
    test_requires = test_requires.extend(test_pkgs)

setuptools.setup(
    name=pkg_name,
    version=version,
    description="The OpenAPI Artifact Generator Python Package",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/open-traffic-generator/openapiart",
    author="https://github.com/open-traffic-generator/openapiart",
    author_email="andy.balogh@keysight.com",
    license="MIT",
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Testing :: Traffic Generation",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 2.7",
        "Programming Language :: Python :: 3",
    ],
    keywords="testing openapi artifact generator",
    package_data={"openapiart": ["*.go", "goserver/*.go"]},
    include_package_data=True,
    packages=[pkg_name],
    python_requires=">=2.7, <4",
    install_requires=installation_requires,
    extras_require={"testing": test_requires},
    test_suite="tests",
)
