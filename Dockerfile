# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.234.0/containers/cpp/.devcontainer/base.Dockerfile

# [Choice] Debian / Ubuntu version (use Debian 11, Ubuntu 18.04/22.04 on local arm64/Apple Silicon): debian-11, debian-10, ubuntu-22.04, ubuntu-20.04, ubuntu-18.04
FROM golang:1.18-bullseye as builder



# [Optional] Install CMake version different from what base image has already installed. 
# CMake reinstall choices: none, 3.21.5, 3.22.2, or versions from https://cmake.org/download/
ARG REINSTALL_CMAKE_VERSION_FROM_SOURCE="3.23.2"

# Optionally install the cmake for vcpkg
COPY tools/reinstall-cmake.sh /tmp/
RUN if [ "${REINSTALL_CMAKE_VERSION_FROM_SOURCE}" != "none" ]; then \
        chmod +x /tmp/reinstall-cmake.sh && /tmp/reinstall-cmake.sh ${REINSTALL_CMAKE_VERSION_FROM_SOURCE}; \
    fi \
    && rm -f /tmp/reinstall-cmake.sh

# [Optional] Uncomment this section to install additional vcpkg ports.
# RUN su vscode -c "${VCPKG_ROOT}/vcpkg install <your-port-name-here>"

# [Optional] Uncomment this section to install additional packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
     && apt-get -y install --no-install-recommends software-properties-common ninja-build

# Install llvm
COPY tools/llvm.sh /tmp/
RUN chmod +x /tmp/llvm.sh
RUN /tmp/llvm.sh

RUN update-alternatives --install /usr/bin/cc cc /usr/bin/clang-14 100
RUN update-alternatives --install /usr/bin/c++ c++ /usr/bin/clang++-14 100
RUN update-alternatives --install /usr/bin/clang cc /usr/bin/clang-14 100
RUN update-alternatives --install /usr/bin/clang++ c++ /usr/bin/clang++-14 100
RUN update-alternatives --install /usr/bin/lldb lldb /usr/bin/lldb-14 100

RUN groupadd -g 1001 appuser
RUN useradd -r -u 1001 -g appuser appuser
RUN mkdir /opt/sciarc
RUN chown appuser /opt/sciarc

COPY . /opt/sciarc
WORKDIR /opt/sciarc
RUN mkdir build \
    && cd build \
    && cmake -DCMAKE_C_COMPILER=/usr/bin/clang -DCMAKE_CXX_COMPILER=/usr/bin/clang++ -GNinja .. \
    && ninja \
    && ninja install

RUN cd /opt/sciarc \
    && go mod tidy \
    && go build


from debian:bullseye

RUN mkdir /opt/app
WORKDIR /opt/app
COPY --from=builder /opt/sciarc/build/local/lib /opt/app/lib
COPY --from=builder /opt/sciarc/sciarc /opt/app
ENV PATH="${PATH}:/opt/app"
ENV LD_LIBRARY_PATH="${LD_LIBRARY_PATH}:/opt/app/lib:/opt/app/lib/linux-x86_64"
EXPOSE 8000
ENTRYPOINT ["sciarc"]