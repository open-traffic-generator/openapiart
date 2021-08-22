"""
To build distribution: python setup.py sdist --formats=gztar bdist_wheel --universal
"""
import os
import setuptools

pkg_name = "openapiart"
version = "0.0.64"

base_dir = os.path.dirname(os.path.abspath(__file__))
with open(os.path.join(base_dir, "README.md")) as fid:
    long_description = fid.read()

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
    package_data={"openapiart": ["*.go"]},
    include_package_data=True,
    packages=[pkg_name],
    python_requires=">=2.7, <4",
    install_requires=[
        "grpcio",
        "grpcio-tools",
        "requests",
        "pyyaml",
        "pytest",
        "openapi-spec-validator",
        "jsonpath-ng",
        "typing",
    ],
    extras_require={"testing": ["pytest", "flake8", "black", "flask"]},
    test_suite="tests",
)
