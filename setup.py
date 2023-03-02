"""
To build distribution: python setup.py sdist --formats=gztar bdist_wheel --universal
"""
import os
import setuptools

pkg_name = "openapiart"
version = "0.2.13"

base_dir = os.path.dirname(os.path.abspath(__file__))
with open(os.path.join(base_dir, "README.md")) as fid:
    long_description = fid.read()

requirements_path = os.path.join(base_dir, "openapiart", "requirements.txt")
test_req_path = os.path.join(base_dir, "test_requirements.txt")
installation_requires = []
test_requires = []
if os.path.exists(requirements_path) is False:
    raise Exception("Could not find requirements path")
with open(requirements_path) as f:
    installation_requires = f.read().splitlines()
    if "--prefer-binary" in installation_requires:
        installation_requires.remove("--prefer-binary")
if os.path.exists(test_req_path):
    with open("test_requirements.txt") as f:
        test_requires = f.read().splitlines()

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
    package_data={"openapiart": ["*.go", "goserver/*.go", "*.txt"]},
    include_package_data=True,
    packages=[pkg_name],
    python_requires=">=2.7, <4",
    install_requires=installation_requires,
    extras_require={"testing": test_requires},
    test_suite="tests",
)
