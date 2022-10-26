FROM ubuntu:22.04
ENV SRC_ROOT=/home/openapiart
ENV PATH="${PATH}:/usr/local/go/bin:$HOME/go/bin:$HOME/.local/bin"
RUN mkdir -p ${SRC_ROOT}
RUN apt-get update \
    && apt-get -y install --no-install-recommends sudo curl git vim unzip net-tools python-is-python3 python3-pip
# Get project source, install dependencies and build it
COPY . ${SRC_ROOT}/
RUN cd ${SRC_ROOT} && python ./do.py setup
RUN cd ${SRC_ROOT} && python ./do.py py_deps
RUN cd ${SRC_ROOT} && python ./do.py go_deps
WORKDIR ${SRC_ROOT}
CMD ["/bin/bash"]
