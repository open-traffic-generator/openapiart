"""Build distributions

To build `python setup.py sdist --formats=gztar bdist_wheel --universal`
"""
import os
import setuptools

pkg_name = 'openapiart'

base_dir = os.path.dirname(os.path.abspath(__file__))
with open(os.path.join(base_dir, 'README.md')) as fid:
    long_description = fid.read()
with open(os.path.join(base_dir, 'VERSION')) as fid:
    version_number = fid.read()

setuptools.setup(
    name=pkg_name,
    version=version_number,
    description='The OpenAPI Artifact Generator Python Package',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/open-traffic-generator/oapiag',
    author='ajbalogh',
    author_email='andy.balogh@keysight.com',
    license='MIT',
    classifiers=[
        'Development Status :: 3 - Alpha', 'Intended Audience :: Developers',
        'Topic :: Software Development :: Testing :: Traffic Generation',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3'
    ],
    keywords='testing openapi artifact generator',
    include_package_data=True,
    packages=[pkg_name],
    python_requires='>=2.7, <4',
    install_requires=[
        'requests',
        'pyyaml',
        'jsonpath-ng',
        'typing'
    ],
    tests_require=['pytest'],
    test_suite='tests'
)
